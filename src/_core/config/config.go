package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetConfig() Config {
	var config Config

	// Загрузка конфигурации из файла JSON
	if err := loadConfigFile("config.json", &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: No .env file found - %v\n", err)
		os.Exit(1)
	}

	// Обновление конфигурации переменными окружения
	if err := updateConfigFromEnv(&config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	return config
}

func loadConfigFile(filePath string, config *Config) error {
	configFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening config file: %w", err)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(config); err != nil {
		return fmt.Errorf("error parsing config json: %w", err)
	}

	return nil
}

func updateConfigFromEnv(config *Config) error {
	redisHost, exists := os.LookupEnv("APP_REDIS_HOST")
	if !exists {
		return fmt.Errorf("APP_REDIS_HOST not set in the environment")
	}
	config.Redis.Host = redisHost

	wsPortStr, exists := os.LookupEnv("APP_WS_PORT")
	if !exists {
		return fmt.Errorf("APP_WS_PORT not set in the environment")
	}

	wsPort, err := strconv.Atoi(wsPortStr)
	if err != nil {
		return fmt.Errorf("error converting APP_WS_PORT to integer: %w", err)
	}
	config.Ws.Port = wsPort

	pgHost, exists := os.LookupEnv("APP_PG_HOST")
	if !exists {
		return fmt.Errorf("APP_PG_HOST not set in the environment")
	}
	config.Pg.Host = pgHost

	brokerCoreEnv, exists := os.LookupEnv("APP_BROKER_CORE_ENV")
	if !exists {
		return fmt.Errorf("APP_BROKER_CORE_ENV not set in the environment")
	}
	config.Broker.Core.Env = brokerCoreEnv

	apiPort, exists := os.LookupEnv("APP_API_PORT")
	if !exists {
		return fmt.Errorf("APP_API_PORT not set in the environment")
	}

	config.Api.Port = apiPort

	return nil
}
