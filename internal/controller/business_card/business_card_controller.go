package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/business_card/dto"
	"github.com/Misoten-B/airship-backend/internal/database"
	"github.com/Misoten-B/airship-backend/internal/database/model"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/gin-gonic/gin"
)

// @Tags BusinessCard
// @Router /v1/users/business_cards [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param CreateBusinessCardRequest formData dto.CreateBusinessCardRequest true "BusinessCard"
// @Success 201 {object} dto.BusinessCardResponse
func CreateBusinessCard(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	request := dto.CreateBusinessCardRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("formData: %v", request)

	file, fileHeader, err := c.Request.FormFile("BusinessCardBackgroundImage")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("file: %v", file)
	log.Printf("fileHeader: %v", fileHeader)

	c.Header("Location", fmt.Sprintf("/%s", "1"))

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
// @Router /v1/users/business_cards [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.BusinessCardResponse
func ReadAllBusinessCard(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	businessCardResponse := []dto.BusinessCardResponse{}
	businessCards := model.BusinessCard{}
	db.Find(&businessCards).Select(`
        business_cards.id, 
        business_card_backgrounds.color_code as BusinessCardBackgroundColor, 
        business_card_backgrounds.image_path as BusinessCardBackgroundImage, 
        business_cards.business_card_name, 
        three_dimentional_models.model_path as ThreeDimentionalModel, 
        speaking_assets.description as SpeakingDescription, 
        speaking_assets.audio_path as SpeakingAudioPath, 
        business_card_parts_coordinates.*, 
        business_cards.display_name, 
        business_cards.company_name, 
        business_cards.department, 
        business_cards.official_position, 
        business_cards.phone_number, 
        business_cards.email, 
        business_cards.postal_code, 
        business_cards.address`).
		Joins("left join business_card_backgrounds on business_card_backgrounds.id = business_cards.business_card_background_id").
		Joins("left join ar_assets on ar_assets.id = business_cards.ar_asset_id").
		Joins("left join three_dimentional_models on three_dimentional_models.id = ar_assets.three_dimentional_model_id").
		Joins("left join speaking_assets on speaking_assets.id = ar_assets.speaking_asset_id").
		Joins("left join business_card_parts_coordinates on business_card_parts_coordinates.id = business_cards.business_card_parts_coordinate_id").
		Joins("inner join users on users.id = business_cards.user_id").
		Where("users.id = ?", uid).
		Scan(&businessCardResponse)

	c.JSON(http.StatusOK, businessCardResponse)
}

// @Tags BusinessCard
// @Router /v1/users/business_cards/{business_card_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Success 200 {object} dto.BusinessCardResponse
func ReadBusinessCardByID(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

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

	c.JSON(http.StatusOK, dto.BusinessCardResponse{
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
// @Router /v1/users/business_cards/{business_card_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Accept multipart/form-data
// @Param BusinessCardBackgroundImage formData file true "Image file to be uploaded"
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

	file, fileHeader, err := c.Request.FormFile("BusinessCardBackgroundImage")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("file: %v", file)
	log.Printf("fileHeader: %v", fileHeader)

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
