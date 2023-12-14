package handler

import (
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card_parts_coordinate/dto"
	"github.com/gin-gonic/gin"
)

// @Tags BusinessCardPartsCoordinate
// @Router /v1/business_card_parts_coordinates [GET]
// @Success 200 {object} []dto.BusinessCardPartsCoordinateResponse
func ReadAllBusinessCardPartsCoordinate(c *gin.Context) {
	// データベースに接続
	db, err := frameworks.GetDB(c)
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
}
