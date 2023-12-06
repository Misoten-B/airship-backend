package arassets

import (
	handler "github.com/Misoten-B/airship-backend/internal/frameworks/handler/ar_assets"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/ar_assets/:ar_assets_id", handler.ReadArAssetsByIDPublic)
}
