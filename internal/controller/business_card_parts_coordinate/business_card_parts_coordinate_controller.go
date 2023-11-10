package controller

import (
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/business_card_parts_coordinate/dto"
	"github.com/gin-gonic/gin"
)

// @Tags BusinessCardPartsCoordinate
// @Router /business_card_parts_coordinate [GET]
// @Success 201 {object} []dto.BusinessCardPartsCoordinateResponse
func ReadAllBusinessCardPartsCoordinate(c *gin.Context) {
	c.JSON(http.StatusCreated, []dto.BusinessCardPartsCoordinateResponse{
		{
			ID:                "1",
			DisplayNameX:      0,
			DisplayNameY:      0,
			CompanyNameX:      0,
			CompanyNameY:      0,
			DepartmentX:       0,
			DepartmentY:       0,
			OfficialPositionX: 0,
			OfficialPositionY: 0,
			PhoneNumberX:      0,
			PhoneNumberY:      0,
			EmailX:            0,
			EmailY:            0,
			PostalCodeX:       0,
			PostalCodeY:       0,
			AddressX:          0,
			AddressY:          0,
			QrcodeX:           0,
			QrcodeY:           0,
		},
	})
}
