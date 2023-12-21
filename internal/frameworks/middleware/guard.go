package middleware

import (
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/testdata"
	"github.com/gin-gonic/gin"
)

func Guard(client *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, err := frameworks.GetConfig(c)
		if err != nil {
			log.Printf("%s", err)
		}
		devMode := value.DevMode

		if devMode {
			log.Println("Development mode - bypassing authentication")
			c.Set(frameworks.ContextKeyUID, testdata.DEV_UID)
			c.Next()
			return
		}

		idToken := c.GetHeader("Authorization")

		if !strings.HasPrefix(idToken, "Bearer ") {
			c.AbortWithStatusJSON(
				http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "token is invalid format"},
			)
			return
		}

		idToken = strings.TrimPrefix(idToken, "Bearer ")

		token, err := client.VerifyIDToken(c.Request.Context(), idToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "error": err.Error()})
			return
		}

		c.Set(frameworks.ContextKeyUID, token.UID)
		c.Next()
	}
}
