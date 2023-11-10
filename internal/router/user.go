package router

import (
	bcbController "github.com/Misoten-B/airship-backend/internal/controller/business_card_background"
	userController "github.com/Misoten-B/airship-backend/internal/controller/user"
	"github.com/gin-gonic/gin"
)

func user(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/:user_id", userController.ReadUserByID)
		userGroup.PUT("/:user_id", userController.UpdateUser)
		userGroup.DELETE("/:user_id", userController.DeleteUser)

		bcbGroup := userGroup.Group("/business_card_background")
		{
			bcbGroup.GET("/", bcbController.ReadAllBusinessCardBackground)
			bcbGroup.POST("/", bcbController.CreateBusinessCardBackground)
		}
	}
}
