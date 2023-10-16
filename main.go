package main

import (
	"github.com/Misoten-B/airship-backend/internal/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/Misoten-B/airship-backend/docs"
)

// @title           Swagger Example API
// @version         1.0
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.Default()
	routes.Register(r)
	r.Run()
}
