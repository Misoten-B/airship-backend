package routes

import (
	controller "github.com/Misoten-B/airship-backend/internal/controller/user"
	"github.com/gin-gonic/gin"
)

func user(r *gin.Engine) {
	ur := r.Group("/user")
	{
		ur.POST("/", controller.CreateUser)
		ur.GET("/:id", controller.ReadUserById)
		ur.PUT("/:id", controller.UpdateUser)
		ur.DELETE("/:id", controller.DeleteUser)
	}
}
