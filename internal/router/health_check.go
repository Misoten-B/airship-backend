package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckRegister(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "alive")
	})
}
