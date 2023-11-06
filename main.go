package main

import (
	"log"

	"github.com/Misoten-B/airship-backend/internal/database"
	"github.com/Misoten-B/airship-backend/internal/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/Misoten-B/airship-backend/docs"
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
	r := gin.Default()
	routes.Register(r)

	_, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	log.Fatal(r.Run())
}
