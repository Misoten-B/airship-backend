package router

import (
	controller "github.com/Misoten-B/airship-backend/internal/controller/business_card_parts_coordinate"
	"github.com/gin-gonic/gin"
)

func businessCardPartsCoordinate(r *gin.Engine) {
	r.GET("/business_card_parts_coordinates", controller.ReadAllBusinessCardPartsCoordinate)
}
