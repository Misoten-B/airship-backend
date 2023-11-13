package router

import (
	arController "github.com/Misoten-B/airship-backend/internal/controller/ar_assets"
	bcController "github.com/Misoten-B/airship-backend/internal/controller/business_card"
	bcbController "github.com/Misoten-B/airship-backend/internal/controller/business_card_background"
	tdController "github.com/Misoten-B/airship-backend/internal/controller/three_dimentional"
	userController "github.com/Misoten-B/airship-backend/internal/controller/user"
	"github.com/gin-gonic/gin"
)

func user(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("", userController.CreateUser)
		userGroup.GET("/:user_id", userController.ReadUserByID)
		userGroup.PUT("/:user_id", userController.UpdateUser)
		userGroup.DELETE("/:user_id", userController.DeleteUser)

		bcbGroup := userGroup.Group("/business_card_background")
		{
			bcbGroup.GET("", bcbController.ReadAllBusinessCardBackground)
			bcbGroup.POST("", bcbController.CreateBusinessCardBackground)
		}

		bcGroup := userGroup.Group("/business_card")
		{
			bcGroup.GET("", bcController.ReadAllBusinessCard)
			bcGroup.GET("/:business_card_id", bcController.ReadBusinessCardByID)
			bcGroup.POST("", bcController.CreateBusinessCard)
			bcGroup.PUT("/:business_card_id", bcController.UpdateBusinessCard)
			bcGroup.DELETE("/:business_card_id", bcController.DeleteBusinessCard)
		}

		tdGroup := userGroup.Group("/three_dimentional")
		{
			tdGroup.GET("", tdController.ReadAllThreeDimentional)
			tdGroup.POST("", tdController.CreateThreeDimentional)
			tdGroup.GET("/:three_dimentional_id", tdController.ReadThreeDimentionalByID)
			tdGroup.PUT("/:three_dimentional_id", tdController.UpdateThreeDimentional)
			tdGroup.DELETE("/:three_dimentional_id", tdController.DeleteThreeDimentional)
		}

		arGroup := userGroup.Group("/ar_assets")
		{
			arGroup.GET("", arController.ReadAllArAssets)
			arGroup.POST("", arController.CreateArAssets)
			arGroup.GET("/:ar_assets_id", arController.ReadArAssetsByID)
			arGroup.PUT("/:ar_assets_id", arController.UpdateArAssets)
			arGroup.DELETE("/:ar_assets_id", arController.DeleteArAssets)
		}
	}
}
