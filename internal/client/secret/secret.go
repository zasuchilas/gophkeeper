package secret

type Secret interface {
	EncryptToBase64() ([]byte, error)
	DecryptFromBase64([]byte) error
}
