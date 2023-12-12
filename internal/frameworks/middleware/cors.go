package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		originList := []string{"http://localhost:3000", "https://airship.azurewebsites.net/"}
		origin := c.Request.Header.Get("Origin")
		isOriginAllowed := slices.Contains(originList, origin)

		if isOriginAllowed {

			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}

			c.Next()
		}
	}
}
