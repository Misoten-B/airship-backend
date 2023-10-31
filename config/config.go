package config

import (
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
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var config Config
	config.Database.Host = os.Getenv("POSTGRES_HOST")
	config.Database.Port = os.Getenv("POSTGRES_PORT")
	config.Database.Dbname = os.Getenv("POSTGRES_DB")
	config.Database.User = os.Getenv("POSTGRES_USER")
	config.Database.Password = os.Getenv("POSTGRES_PASSWORD")

	return config
}
