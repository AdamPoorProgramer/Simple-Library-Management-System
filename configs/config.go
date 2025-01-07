package configs

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Database struct {
		Name string `json:"name"`
	} `json:"database"`
	Logging struct {
		Level string `json:"level"`
		Path  string `json:"path"`
	} `json:"logging"`
}

// LoadConfig reads the config file and unmarshals it into the Config struct
func LoadConfig() (*Config, error) {
	file, err := os.Open("./config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(byteValue, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
