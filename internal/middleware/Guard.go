package middleware

import (
	"log"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func Guard(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")
		token, err := client.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error: ": "Unauthorized"})
		}

		log.Printf("Verified ID token: %v\n", token)
		c.Next()
	}
}
