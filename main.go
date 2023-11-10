package main

import (
	"log"

	"github.com/Misoten-B/airship-backend/internal/router"
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

	router.Register(r)

	log.Fatal(r.Run())
}
