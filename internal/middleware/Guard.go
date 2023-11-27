package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Guard(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}

		fmt.Printf("PRD_MODE: %s\n", os.Getenv("PRD_MODE"))

		if os.Getenv("PRD_MODE") != "true" {
			log.Println("Development mode - bypassing authentication")
			c.Next()
			return
		}

		idToken := c.GetHeader("Authorization")
		token, err := client.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "error": "Unauthorized"})
			return
		}

		log.Printf("Verified ID token: %v\n", token)
		c.Next()
	}
}
