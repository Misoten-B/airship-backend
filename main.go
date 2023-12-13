//go:generate swag init
//go:generate wire ./internal/container

package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Misoten-B/airship-backend/config"
	_ "github.com/Misoten-B/airship-backend/docs"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
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
		ctx.Next()
	})

	router.HealthCheckRegister(r)
	router.OpenapiRegister(r)
	v1.Register(r)

	log.Fatal(r.Run())
}

type AppConfig struct {
	Config *config.Config
}

func setup() (AppConfig, error) {
	config, err := config.GetConfig()
	if err != nil {
		return AppConfig{}, err
	}

	return AppConfig{
		Config: config,
	}, nil
}
