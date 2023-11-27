package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/business_card/dto"
	"github.com/gin-gonic/gin"
)

// @Tags BusinessCard
// @Router /v1/users/business_cards [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param BusinessCardBackgroundImage formData file true "Image file to be uploaded"
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
	c.JSON(http.StatusCreated, dto.BusinessCardResponse{
		ID:                          "1",
		BusinessCardBackgroundColor: "#ffffff",
		BusinessCardBackgroundImage: "https://example.com/image.png",
		BusinessCardName:            "会社",
		ThreeDimentionalModel:       "https://example.com/model.gltf",
		SpeakingDescription:         "こんにちは",
		SpeakingAudioPath:           "https://example.com/audio.mp3",
		BusinessCardPartsCoordinate: "1",
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

	c.JSON(http.StatusOK, []dto.BusinessCardResponse{
		{
			ID:                          "1",
			BusinessCardBackgroundColor: "#ffffff",
			BusinessCardBackgroundImage: "https://example.com/image.png",
			BusinessCardName:            "会社",
			ThreeDimentionalModel:       "https://example.com/model.gltf",
			SpeakingDescription:         "こんにちは",
			SpeakingAudioPath:           "https://example.com/audio.mp3",
			BusinessCardPartsCoordinate: "1",
			DisplayName:                 "山田太郎",
			CompanyName:                 "株式会社山田",
			Department:                  "開発部",
			OfficialPosition:            "部長",
			PhoneNumber:                 "090-1234-5678",
			Email:                       "sample@example.com",
			PostalCode:                  "123-4567",
			Address:                     "東京都渋谷区1-1-1",
		},
	})
}

// @Tags BusinessCard
// @Router /v1/users/business_cards/{business_card_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Success 200 {object} dto.BusinessCardResponse
func ReadBusinessCardByID(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	c.JSON(http.StatusOK, dto.BusinessCardResponse{
		ID:                          "1",
		BusinessCardBackgroundColor: "#ffffff",
		BusinessCardBackgroundImage: "https://example.com/image.png",
		BusinessCardName:            "会社",
		ThreeDimentionalModel:       "https://example.com/model.gltf",
		SpeakingDescription:         "こんにちは",
		SpeakingAudioPath:           "https://example.com/audio.mp3",
		BusinessCardPartsCoordinate: "1",
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

	c.JSON(http.StatusCreated, dto.BusinessCardResponse{
		ID:                          "1",
		BusinessCardBackgroundColor: "#ffffff",
		BusinessCardBackgroundImage: "https://example.com/image.png",
		BusinessCardName:            "会社",
		ThreeDimentionalModel:       "https://example.com/model.gltf",
		SpeakingDescription:         "こんにちは",
		SpeakingAudioPath:           "https://example.com/audio.mp3",
		BusinessCardPartsCoordinate: "1",
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
