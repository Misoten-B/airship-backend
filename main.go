//go:generate swag init
//go:generate wire ./internal/container

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "github.com/Misoten-B/airship-backend/docs"
	"github.com/Misoten-B/airship-backend/internal/drivers/config"
	"github.com/Misoten-B/airship-backend/internal/drivers/database"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/middleware"
	"github.com/Misoten-B/airship-backend/internal/frameworks/router"
	v1 "github.com/Misoten-B/airship-backend/internal/frameworks/router/v1"
)

// @title           AIRship API
// @version         1.0
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host 		  localhost:8080
// @@securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization
func main() {
	appConfig, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set(frameworks.ContextKeyConfig, appConfig.Config)
		ctx.Set(frameworks.ContextKeyDB, appConfig.Database)
		ctx.Next()
	})

	r.Use(middleware.CORSMiddleware())

	router.HealthCheckRegister(r)
	router.OpenapiRegister(r)
	v1.Register(r)

	log.Fatal(r.Run())
}

type AppConfig struct {
	Config   *config.Config
	Database *gorm.DB
}

func setup() (AppConfig, error) {
	config, err := config.GetConfig()
	if err != nil {
		return AppConfig{}, err
	}

	db, err := database.ConnectDB()
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Config:   config,
		Database: db,
	}, nil
}
