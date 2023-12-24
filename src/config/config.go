package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfig() Config {
	configFile, err := os.Open("config.json")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening config file: %v\n", err)
		os.Exit(1)
	}

	defer configFile.Close()

	var config Config

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing config json: %v\n", err)
		os.Exit(1)
	}

	return config
}
