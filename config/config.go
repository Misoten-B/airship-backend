package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		Dbname   string
		User     string
		Password string
	}
	DevMode                          bool
	AzureBlobStorageConnectionString string
}

func GetConfig() (*Config, error) {
	// .envファイルを読み込む
	// `export`された環境変数が優先される
	err := godotenv.Load()
	if err != nil {
		log.Println("the .env file is not found, so use the default value")
	}

	dbHost, err := getEnv("POSTGRES_HOST")
	if err != nil {
		return nil, err
	}
	dbPort, err := getEnv("POSTGRES_PORT")
	if err != nil {
		return nil, err
	}
	dbDbname, err := getEnv("POSTGRES_DB")
	if err != nil {
		return nil, err
	}
	dbUser, err := getEnv("POSTGRES_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := getEnv("POSTGRES_PASSWORD")
	if err != nil {
		return nil, err
	}

	devMode := getEnvWithDefaultValue("DEV_MODE", "false")

	azBlobStorageConnectionString, err := getEnv("AZURE_BLOB_STORAGE_CONNECTION_STRING")
	if err != nil {
		return nil, err
	}

	return &Config{
		DevMode:                          devMode == "true",
		AzureBlobStorageConnectionString: azBlobStorageConnectionString,
		Database: struct {
			Host     string
			Port     string
			Dbname   string
			User     string
			Password string
		}{
			Host:     dbHost,
			Port:     dbPort,
			Dbname:   dbDbname,
			User:     dbUser,
			Password: dbPassword,
		},
	}, nil
}

func getEnvWithDefaultValue(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

func getEnv(key string) (string, error) {
	value := getEnvWithDefaultValue(key, "")
	if value == "" {
		return "", fmt.Errorf("required environment variable %s is not set", key)
	}
	return value, nil
}
