package users

import (
	"log"

	arController "github.com/Misoten-B/airship-backend/internal/controller/ar_assets"
	bcController "github.com/Misoten-B/airship-backend/internal/controller/business_card"
	bcbController "github.com/Misoten-B/airship-backend/internal/controller/business_card_background"
	tdController "github.com/Misoten-B/airship-backend/internal/controller/three_dimentional"
	userController "github.com/Misoten-B/airship-backend/internal/controller/user"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	client, err := drivers.GetFirebaseAuthClient()
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	users := r.Group("/users")
	users.Use(middleware.Guard(client))
	{
		users.POST("", userController.CreateUser)
		users.GET("/:user_id", userController.ReadUserByID)
		users.PUT("/:user_id", userController.UpdateUser)
		users.DELETE("/:user_id", userController.DeleteUser)
		businessCardBackgroundRegister(users)
		businessCardRegister(users)
		threeDimentionalRegister(users)
		arAssetsRegister(users)
	}
}

func arAssetsRegister(r *gin.RouterGroup) {
	ar := r.Group("/ar_assets")
	{
		ar.GET("", arController.ReadAllArAssets)
		ar.POST("", arController.CreateArAssets)
		ar.GET("/:ar_assets_id", arController.ReadArAssetsByID)
		ar.PUT("/:ar_assets_id", arController.UpdateArAssets)
		ar.DELETE("/:ar_assets_id", arController.DeleteArAssets)
	}
}

func threeDimentionalRegister(r *gin.RouterGroup) {
	td := r.Group("/three_dimentionals")
	{
		td.GET("", tdController.ReadAllThreeDimentional)
		td.POST("", tdController.CreateThreeDimentional)
		td.GET("/:three_dimentional_id", tdController.ReadThreeDimentionalByID)
		td.PUT("/:three_dimentional_id", tdController.UpdateThreeDimentional)
		td.DELETE("/:three_dimentional_id", tdController.DeleteThreeDimentional)
	}
}

func businessCardBackgroundRegister(r *gin.RouterGroup) {
	bcb := r.Group("/business_card_backgrounds")
	{
		bcb.GET("", bcbController.ReadAllBusinessCardBackground)
		bcb.POST("", bcbController.CreateBusinessCardBackground)
	}
}

func businessCardRegister(r *gin.RouterGroup) {
	bc := r.Group("/business_cards")
	{
		bc.GET("", bcController.ReadAllBusinessCard)
		bc.GET("/:business_card_id", bcController.ReadBusinessCardByID)
		bc.POST("", bcController.CreateBusinessCard)
		bc.PUT("/:business_card_id", bcController.UpdateBusinessCard)
		bc.DELETE("/:business_card_id", bcController.DeleteBusinessCard)
	}
}
