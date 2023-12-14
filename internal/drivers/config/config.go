package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database                         databaseConfig
	DevMode                          bool
	CORSOrigin                       string
	AzureBlobStorageConnectionString string
}

type databaseConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

func GetConfig() (*Config, error) {
	// .envファイルを読み込む
	// `export`された環境変数が優先される
	err := godotenv.Load()
	if err != nil {
		log.Println("using default values for configuration")
	}

	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		return nil, err
	}
	dbPort, err := getEnv("DB_PORT")
	if err != nil {
		return nil, err
	}
	dbName, err := getEnv("DB_NAME")
	if err != nil {
		return nil, err
	}
	dbUser, err := getEnv("DB_USER")
	if err != nil {
		return nil, err
	}
	dbPassword, err := getEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	devModeStr := getEnvWithDefaultValue("DEV_MODE", "false")
	devMode := devModeStr == "true"

	corsOriginStr, err := getEnv("CORS_ORIGIN")
	if err != nil {
		return nil, err
	}

	azBlobStorageConnectionString, err := getEnv("AZURE_BLOB_STORAGE_CONNECTION_STRING")
	if err != nil {
		return nil, err
	}

	dbConfig := databaseConfig{
		Host:     dbHost,
		Port:     dbPort,
		DBName:   dbName,
		User:     dbUser,
		Password: dbPassword,
	}
	return &Config{
		DevMode:                          devMode,
		CORSOrigin:                       corsOriginStr,
		AzureBlobStorageConnectionString: azBlobStorageConnectionString,
		Database:                         dbConfig,
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
