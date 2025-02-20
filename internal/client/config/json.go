package config

import (
	"encoding/json"
	"os"
)

type jsonConfig struct {
	ServerAddress string `json:"server_address"`
}

func getJSONConfig(filename string) (*jsonConfig, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jc jsonConfig
	err = json.Unmarshal(file, &jc)
	if err != nil {
		return nil, err
	}

	return &jc, nil
}
