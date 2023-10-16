package routes

import "github.com/gin-gonic/gin"

func healthCheck(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "alive")
	})
}
