package controller

import (
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/drivers/database"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card_parts_coordinate/dto"
	"github.com/gin-gonic/gin"
)

// @Tags BusinessCardPartsCoordinate
// @Router /v1/business_card_parts_coordinates [GET]
// @Success 200 {object} []dto.BusinessCardPartsCoordinateResponse
func ReadAllBusinessCardPartsCoordinate(c *gin.Context) {
	// データベースに接続
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// データベースから全ての名刺パーツ座標を取得
	bcpcs := []model.BusinessCardPartsCoordinate{}
	result := db.Find(&bcpcs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}

	// レスポンスに変換
	bcpcr := []dto.BusinessCardPartsCoordinateResponse{}
	for _, bcpc := range bcpcs {
		bcpcr = append(bcpcr, dto.ConvertToBCPCResponse(bcpc))
	}

	c.JSON(http.StatusOK, bcpcr)

	// c.JSON(http.StatusOK, []dto.BusinessCardPartsCoordinateResponse{
	// 	{
	// 		ID:                "1",
	// 		DisplayNameX:      0,
	// 		DisplayNameY:      0,
	// 		CompanyNameX:      0,
	// 		CompanyNameY:      0,
	// 		DepartmentX:       0,
	// 		DepartmentY:       0,
	// 		OfficialPositionX: 0,
	// 		OfficialPositionY: 0,
	// 		PhoneNumberX:      0,
	// 		PhoneNumberY:      0,
	// 		EmailX:            0,
	// 		EmailY:            0,
	// 		PostalCodeX:       0,
	// 		PostalCodeY:       0,
	// 		AddressX:          0,
	// 		AddressY:          0,
	// 		QrcodeX:           0,
	// 		QrcodeY:           0,
	// 	},
	// })
}
