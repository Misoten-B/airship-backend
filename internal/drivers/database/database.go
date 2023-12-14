package database

import (
	"fmt"
	"log"

	"github.com/Misoten-B/airship-backend/internal/drivers/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Println("Failed to get config")
		return nil, err
	}

	dialector := GetMySQLDialector(
		MySQLDSNParams{
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			User:     cfg.Database.User,
			Password: cfg.Database.Password,
			Dbname:   cfg.Database.DBName,
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		return nil, err
	}

	return db, nil
}

type MySQLDSNParams struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func GetMySQLDialector(params MySQLDSNParams) gorm.Dialector {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		params.User,
		params.Password,
		params.Host,
		params.Port,
		params.Dbname,
	)

	return mysql.Open(dsn)
}
