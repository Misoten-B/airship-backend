package config

import (
	"errors"
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

	dbHost, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return nil, errors.New("POSTGRES_HOST is not found")
	}
	dbPort, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return nil, errors.New("POSTGRES_PORT is not found")
	}
	dbDbname, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return nil, errors.New("POSTGRES_DB is not found")
	}
	dbUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return nil, errors.New("POSTGRES_USER is not found")
	}
	dbPassword, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return nil, errors.New("POSTGRES_PASSWORD is not found")
	}

	devMode, ok := os.LookupEnv("DEV_MODE")
	if !ok {
		devMode = "false"
	}

	azBlobStorageConnectionString, ok := os.LookupEnv("AZURE_BLOB_STORAGE_CONNECTION_STRING")
	if !ok {
		return nil, errors.New("AZURE_BLOB_STORAGE_CONNECTION_STRING is not found")
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
