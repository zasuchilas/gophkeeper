package secret

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
)

func encodeToBase64[T any](v T) ([]byte, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	defer encoder.Close()
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeFromBase64[T any](v *T, enc []byte) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding,
		bytes.NewReader(enc))).Decode(v)
}
