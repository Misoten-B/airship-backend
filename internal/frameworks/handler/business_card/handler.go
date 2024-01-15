// FIXME: このファイルの関数の長さが長いので、分割する
//
//nolint:funlen
package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/drivers/config"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/business_card/dto"
	"github.com/gin-gonic/gin"
)

const (
	backgroundContainer            = "background-images"
	qrcodeContainer                = "qrcode-images"
	threeDimentionalModelContainer = "three-dimentiional-models"
	audioContainer                 = "audios"
)

// @Tags BusinessCard
// @Router /v1/users/business_cards [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param CreateBusinessCardRequest formData dto.CreateBusinessCardRequest true "BusinessCard"
// @Success 201 {object} dto.BusinessCardResponse
func CreateBusinessCard(c *gin.Context) {
	request := dto.CreateBusinessCardRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := shared.NewID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	businesscard := model.BusinessCard{
		BusinessCardBackgroundID:      request.BusinessCardBackgroundID,
		ARAssetID:                     request.ArAssetsID,
		BusinessCardPartsCoordinateID: request.BusinessCardPartsCoordinateID,
		BusinessCardName:              request.BusinessCardName,
		ID:                            id.String(),
		UserID:                        uid,
		DisplayName:                   request.DisplayName,
		CompanyName:                   request.CompanyName,
		Department:                    request.Department,
		OfficialPosition:              request.OfficialPosition,
		PhoneNumber:                   request.PhoneNumber,
		Email:                         request.Email,
		PostalCode:                    request.PostalCode,
		Address:                       request.Address,
		CreatedAt:                     time.Now(),
	}
	if err = db.Create(&businesscard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcc := model.BusinessCardPartsCoordinate{ID: request.BusinessCardPartsCoordinateID}
	if err = db.First(&bcc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcb := model.BusinessCardBackground{ID: request.BusinessCardBackgroundID}
	if err = db.First(&bcb).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	arassets := model.ARAsset{ID: request.ArAssetsID}
	if err = db.First(&arassets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ConvertBC(businesscard, bcc, bcb, arassets)

	c.Header("Location", fmt.Sprintf("/%s", "1"))
	c.JSON(http.StatusCreated, response)
}

// @Tags BusinessCard
// @Router /v1/users/business_cards [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.BusinessCardResponse
func ReadAllBusinessCard(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	businesscards := []model.BusinessCard{}
	if err = db.Where("user_id = ?", uid).Find(&businesscards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appConfig, err := config.GetConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ab := drivers.NewAzureBlobDriver(appConfig)

	bcURL, err := ab.GetContainerURL(backgroundContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	qcURL, err := ab.GetContainerURL(qrcodeContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tdmcURL, err := ab.GetContainerURL(threeDimentionalModelContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	acURL, err := ab.GetContainerURL(audioContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses := []dto.BusinessCardResponse{}
	for _, businesscard := range businesscards {
		bcc := model.BusinessCardPartsCoordinate{ID: businesscard.BusinessCardPartsCoordinateID}
		if err = db.First(&bcc).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		bcb := model.BusinessCardBackground{ID: businesscard.BusinessCardBackgroundID}
		if err = db.First(&bcb).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		bcb.ImagePath = bcURL.Path(bcb.ImagePath)

		arassets := model.ARAsset{ID: businesscard.ARAssetID}
		if err = db.First(&arassets).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		arassets.ThreeDimentionalModel.ModelPath = tdmcURL.Path(arassets.ThreeDimentionalModel.ModelPath)
		arassets.SpeakingAsset.AudioPath = acURL.Path(arassets.SpeakingAsset.AudioPath)

		var qrCodeImagePath string
		if arassets.QRCodeImagePath != "" {
			qrCodeImagePath = qcURL.Path(arassets.QRCodeImagePath)
		}
		arassets.QRCodeImagePath = qrCodeImagePath

		response := dto.ConvertBC(businesscard, bcc, bcb, arassets)

		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, responses)
}

// @Tags BusinessCard
// @Router /v1/users/business_cards/{business_card_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Success 200 {object} dto.BusinessCardResponse
func ReadBusinessCardByID(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appConfig, err := config.GetConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ab := drivers.NewAzureBlobDriver(appConfig)

	bcURL, err := ab.GetContainerURL(backgroundContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	qcURL, err := ab.GetContainerURL(qrcodeContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tdmcURL, err := ab.GetContainerURL(threeDimentionalModelContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	acURL, err := ab.GetContainerURL(audioContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	businesscard := model.BusinessCard{}
	if err = db.Where("id = ? AND user_id = ?", c.Param("business_card_id"), uid).First(&businesscard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcc := model.BusinessCardPartsCoordinate{ID: businesscard.BusinessCardPartsCoordinateID}
	if err = db.First(&bcc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcb := model.BusinessCardBackground{ID: businesscard.BusinessCardBackgroundID}
	if err = db.First(&bcb).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bcb.ImagePath = bcURL.Path(bcb.ImagePath)

	arassets := model.ARAsset{ID: businesscard.ARAssetID}
	if err = db.First(&arassets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	arassets.ThreeDimentionalModel.ModelPath = tdmcURL.Path(arassets.ThreeDimentionalModel.ModelPath)
	arassets.SpeakingAsset.AudioPath = acURL.Path(arassets.SpeakingAsset.AudioPath)

	var qrCodeImagePath string
	if arassets.QRCodeImagePath != "" {
		qrCodeImagePath = qcURL.Path(arassets.QRCodeImagePath)
	}
	arassets.QRCodeImagePath = qrCodeImagePath

	response := dto.ConvertBC(businesscard, bcc, bcb, arassets)

	c.JSON(http.StatusOK, response)
}

// @Tags BusinessCard
// @Router /v1/users/business_cards/{business_card_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Accept multipart/form-data
// @Param CreateBusinessCardRequest formData dto.CreateBusinessCardRequest true "BusinessCard"
// @Success 200 {object} dto.BusinessCardResponse
func UpdateBusinessCard(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
	log.Printf("business_card_id: %s", c.Param("business_card_id"))

	request := dto.CreateBusinessCardRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("formData: %v", request)

	businessCardPartsCoordinate := dto.BusinessCardPartsCoordinate{
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
	}

	c.JSON(http.StatusCreated, dto.BusinessCardResponse{
		ID:                          "1",
		BusinessCardBackgroundColor: "#ffffff",
		BusinessCardBackgroundImage: "https://example.com/image.png",
		BusinessCardName:            "会社",
		ThreeDimentionalModel:       "https://example.com/model.gltf",
		SpeakingDescription:         "こんにちは",
		SpeakingAudioPath:           "https://example.com/audio.mp3",
		BusinessCardPartsCoordinate: businessCardPartsCoordinate,
		DisplayName:                 "山田太郎",
		CompanyName:                 "株式会社山田",
		Department:                  "開発部",
		OfficialPosition:            "部長",
		PhoneNumber:                 "090-1234-5678",
		Email:                       "sample@example.com",
		PostalCode:                  "123-4567",
		Address:                     "東京都渋谷区1-1-1",
	})
}

// @Tags BusinessCard
// @Router /v1/users/business_cards/{business_card_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Success 204 {object} nil
func DeleteBusinessCard(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	c.JSON(http.StatusNoContent, nil)
}

// @Tags BusinessCard
// @Router /v1/business_cards/{business_card_id} [GET]
// @Param business_card_id path string true "BusinessCard ID"
// @Success 200 {object} dto.BusinessCardResponse
func ReadBusinessCardByIDPublic(c *gin.Context) {
	// FIXME: まるごとコピペ
	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appConfig, err := config.GetConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ab := drivers.NewAzureBlobDriver(appConfig)

	bcURL, err := ab.GetContainerURL(backgroundContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	qcURL, err := ab.GetContainerURL(qrcodeContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tdmcURL, err := ab.GetContainerURL(threeDimentionalModelContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	acURL, err := ab.GetContainerURL(audioContainer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	businesscard := model.BusinessCard{}
	if err = db.Where("id = ?", c.Param("business_card_id")).First(&businesscard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcc := model.BusinessCardPartsCoordinate{ID: businesscard.BusinessCardPartsCoordinateID}
	if err = db.First(&bcc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bcb := model.BusinessCardBackground{ID: businesscard.BusinessCardBackgroundID}
	if err = db.First(&bcb).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bcb.ImagePath = bcURL.Path(bcb.ImagePath)

	arassets := model.ARAsset{ID: businesscard.ARAssetID}
	if err = db.First(&arassets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	arassets.ThreeDimentionalModel.ModelPath = tdmcURL.Path(arassets.ThreeDimentionalModel.ModelPath)
	arassets.SpeakingAsset.AudioPath = acURL.Path(arassets.SpeakingAsset.AudioPath)

	var qrCodeImagePath string
	if arassets.QRCodeImagePath != "" {
		qrCodeImagePath = qcURL.Path(arassets.QRCodeImagePath)
	}
	arassets.QRCodeImagePath = qrCodeImagePath

	response := dto.ConvertBC(businesscard, bcc, bcb, arassets)

	c.JSON(http.StatusOK, response)
}
