package router

import (
	controller "github.com/Misoten-B/airship-backend/internal/controller/ar_assets"
	"github.com/gin-gonic/gin"
)

func arAssets(r *gin.Engine) {
	r.GET("/ar_assets/:ar_assets_id", controller.ReadArAssetsByIDPublic)
}
