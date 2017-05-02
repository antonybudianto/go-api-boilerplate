package main

import (
	"encoding/json"
	"os"
)

// Config is the struct of our app's config
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// LoadConfig load the config.json file and return its data
func LoadConfig(fileName string) (Config, error) {
	config := Config{}

	configFile, err := os.Open(fileName)
	defer configFile.Close()
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(configFile).Decode(&config)

	return config, err
}
