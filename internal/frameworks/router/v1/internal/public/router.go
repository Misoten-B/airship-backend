package public

import (
	arassets "github.com/Misoten-B/airship-backend/internal/frameworks/handler/ar_assets"
	businesscard "github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/ar_assets/:ar_assets_id", arassets.ReadArAssetsByIDPublic)
	r.GET("/business_cards/:business_card_id", businesscard.ReadBusinessCardByIDPublic)
}
