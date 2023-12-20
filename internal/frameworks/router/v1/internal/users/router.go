package users

import (
	"log"

	"github.com/Misoten-B/airship-backend/internal/drivers"
	arHandler "github.com/Misoten-B/airship-backend/internal/frameworks/handler/ar_assets"
	bcHandler "github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card"
	bcbHandler "github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card_background"
	tdHandler "github.com/Misoten-B/airship-backend/internal/frameworks/handler/three_dimentional"
	userHandler "github.com/Misoten-B/airship-backend/internal/frameworks/handler/user"
	"github.com/Misoten-B/airship-backend/internal/frameworks/middleware"
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
		users.POST("", userHandler.CreateUser)
		users.GET("/myprofile", userHandler.ReadUserByID)
		users.PUT("", userHandler.UpdateUser)
		users.DELETE("", userHandler.DeleteUser)
		users.POST("/:user_id/voice_model/status/done", userHandler.PostVoiceModelStatusDone)
		users.POST("/:user_id/voice_model/status/failed", userHandler.PostVoiceModelStatusFailed)
		businessCardBackgroundRegister(users)
		businessCardRegister(users)
		threeDimentionalRegister(users)
		arAssetsRegister(users)
	}
}

func arAssetsRegister(r *gin.RouterGroup) {
	ar := r.Group("/ar_assets")
	{
		ar.GET("", arHandler.ReadAllArAssets)
		ar.POST("", arHandler.CreateArAssets)
		ar.GET("/:ar_assets_id", arHandler.ReadArAssetsByID)
		ar.PUT("/:ar_assets_id", arHandler.UpdateArAssets)
		ar.DELETE("/:ar_assets_id", arHandler.DeleteArAssets)
		ar.POST("/:ar_assets_id/status/done", arHandler.PostStatusDone)
		ar.POST("/:ar_assets_id/status/failed", arHandler.PostStatusFailed)
	}
}

func threeDimentionalRegister(r *gin.RouterGroup) {
	td := r.Group("/three_dimentionals")
	{
		td.GET("", tdHandler.ReadAllThreeDimentional)
		td.POST("", tdHandler.CreateThreeDimentional)
		td.GET("/:three_dimentional_id", tdHandler.ReadThreeDimentionalByID)
		td.PUT("/:three_dimentional_id", tdHandler.UpdateThreeDimentional)
		td.DELETE("/:three_dimentional_id", tdHandler.DeleteThreeDimentional)
	}
}

func businessCardBackgroundRegister(r *gin.RouterGroup) {
	bcb := r.Group("/business_card_backgrounds")
	{
		bcb.GET("", bcbHandler.ReadAllBusinessCardBackground)
		bcb.POST("", bcbHandler.CreateBusinessCardBackground)
	}
}

func businessCardRegister(r *gin.RouterGroup) {
	bc := r.Group("/business_cards")
	{
		bc.GET("", bcHandler.ReadAllBusinessCard)
		bc.GET("/:business_card_id", bcHandler.ReadBusinessCardByID)
		bc.POST("", bcHandler.CreateBusinessCard)
		bc.PUT("/:business_card_id", bcHandler.UpdateBusinessCard)
		bc.DELETE("/:business_card_id", bcHandler.DeleteBusinessCard)
	}
}
