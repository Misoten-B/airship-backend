package router

import (
	controller "github.com/Misoten-B/airship-backend/internal/controller/user"
	"github.com/gin-gonic/gin"
)

func user(r *gin.Engine) {
	ur := r.Group("/user")
	{
		ur.POST("/", controller.CreateUser)
		ur.GET("/:user_id", controller.ReadUserByID)
		ur.PUT("/:user_id", controller.UpdateUser)
		ur.DELETE("/:user_id", controller.DeleteUser)
	}
}
