package businesscardpartscoordinates

import (
	handler "github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card_parts_coordinate"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/business_card_parts_coordinates", handler.ReadAllBusinessCardPartsCoordinate)
}
