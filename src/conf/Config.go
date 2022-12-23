package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ServerConfig map[string]interface{} `json:"server"`
}

func (cfg *Config) loadConfig(file string) (string, int) {
	data, err := os.ReadFile(file)
	if data == nil {
		log.Println("There is error while reading file:", err.Error())
		return err.Error(), 1
	}

	json.Unmarshal(data, &cfg.ServerConfig)
	return "", 0
}

func (cfg *Config) Init(file_path string) int {
	_, err_code := cfg.loadConfig(file_path)
	return err_code
}
