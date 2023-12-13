package config

import (
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

func GetConfig() *Config {
	// .envファイルを読み込む
	// `export`された環境変数が優先される
	err := godotenv.Load()
	if err != nil {
		log.Println("the .env file is not found, so use the default value")
	}

	var config Config

	config.Database.Host = os.Getenv("POSTGRES_HOST")
	config.Database.Port = os.Getenv("POSTGRES_PORT")
	config.Database.Dbname = os.Getenv("POSTGRES_DB")
	config.Database.User = os.Getenv("POSTGRES_USER")
	config.Database.Password = os.Getenv("POSTGRES_PASSWORD")

	config.DevMode = os.Getenv("DEV_MODE") == "true"

	config.AzureBlobStorageConnectionString = os.Getenv("AZURE_BLOB_STORAGE_CONNECTION_STRING")

	return &config
}
