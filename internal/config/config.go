package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerAddress string `json:"serverAddress"`
	ServerCert    string `json:"serverCert"`
	ServerKey     string `json:"serverKey"`
}

func LoadConfig(path string) (*Config, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
