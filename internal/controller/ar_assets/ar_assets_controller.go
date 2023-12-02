package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/ar_assets/dto"
	"github.com/Misoten-B/airship-backend/internal/database"
	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/gin-gonic/gin"
)

// @Tags ArAssets
// @Router /v1/users/ar_assets [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param qrcodeIcon formData file false "Image file to be uploaded"
// @Param dto.CreateArAssetsRequest formData dto.CreateArAssetsRequest true "ArAssets"
// @Success 201 {object} dto.ArAssetsResponse
func CreateArAssets(c *gin.Context) {
	log.Printf("Authorization: %s", c.Request.Header.Get("Authorization"))

	request := dto.CreateArAssetsRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("formData: %v", request)

	multipartFile, fileHeader, err := c.Request.FormFile("qrcodeIcon")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("file: %v", multipartFile)
	log.Printf("fileHeader: %v", fileHeader)

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	file := file.NewMyFile(multipartFile, fileHeader)
	uid := c.GetString("uid")
	arasset, err := arassets.NewARAssets(
		uid,
		file,
		request.SpeakingDescription,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result := db.Create(&arasset)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}
	log.Printf("arasset: %v", arasset)

	c.Header("Location", fmt.Sprintf("/%s", "1"))
	c.JSON(http.StatusCreated, dto.ArAssetsResponse{
		ID:                   "1",
		SpeakingDescription:  "こんにちは",
		SpeakingAudioPath:    "https://example.com",
		ThreeDimentionalPath: "https://example.com",
		QrcodeIconImagePath:  "https://example.com",
	})
}

// @Tags ArAssets
// @Router /v1/users/ar_assets/{ar_assets_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 200 {object} dto.ArAssetsResponse
func ReadArAssetsByID(c *gin.Context) {
	log.Print("Authorization: ", c.Request.Header.Get("Authorization"))
	log.Print("ar_assets_id: ", c.Param("ar_assets_id"))

	c.JSON(http.StatusOK, dto.ArAssetsResponse{
		ID:                   "1",
		SpeakingDescription:  "こんにちは",
		SpeakingAudioPath:    "https://example.com",
		ThreeDimentionalPath: "https://example.com",
		QrcodeIconImagePath:  "https://example.com",
	})
}

// @Tags ArAssets
// @Router /v1/ar_assets/{ar_assets_id} [GET]
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 200 {object} dto.ArAssetsResponse
func ReadArAssetsByIDPublic(c *gin.Context) {
	log.Print("ar_assets_id: ", c.Param("ar_assets_id"))

	c.JSON(http.StatusOK, dto.ArAssetsResponse{
		ID:                   "1",
		SpeakingDescription:  "こんにちは",
		SpeakingAudioPath:    "https://example.com",
		ThreeDimentionalPath: "https://example.com",
		QrcodeIconImagePath:  "https://example.com",
	})
}

// @Tags ArAssets
// @Router /v1/users/ar_assets [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.ArAssetsResponse
func ReadAllArAssets(c *gin.Context) {
	log.Printf("Authorization: %s", c.Request.Header.Get("Authorization"))

	c.JSON(http.StatusOK, []dto.ArAssetsResponse{
		{
			ID:                   "1",
			SpeakingDescription:  "こんにちは",
			SpeakingAudioPath:    "https://example.com",
			ThreeDimentionalPath: "https://example.com",
			QrcodeIconImagePath:  "https://example.com",
		},
	})
}

// @Tags ArAssets
// @Router /v1/users/ar_assets/{ar_assets_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param dto.CreateArAssetsRequest formData dto.CreateArAssetsRequest true "ArAssets"
// @Param ar_assets_id path string true "ArAssets ID"
// @Accept multipart/form-data
// @Param qrcodeIcon formData file false "Image file to be uploaded"
// @Success 200 {object} dto.ArAssetsResponse
func UpdateArAssets(c *gin.Context) {
	log.Printf("Authorization: %s", c.Request.Header.Get("Authorization"))

	request := dto.CreateArAssetsRequest{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("formData: %v", request)

	c.JSON(http.StatusCreated, dto.ArAssetsResponse{
		ID:                   "1",
		SpeakingDescription:  "こんにちは",
		SpeakingAudioPath:    "https://example.com",
		ThreeDimentionalPath: "https://example.com",
		QrcodeIconImagePath:  "https://example.com",
	})
}

// @Tags ArAssets
// @Router /v1/users/ar_assets/{ar_assets_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 204 {object} nil
func DeleteArAssets(c *gin.Context) {
	log.Printf("Authorization: %s", c.Request.Header.Get("Authorization"))
	log.Printf("ar_assets_id: %s", c.Param("ar_assets_id"))

	c.JSON(http.StatusNoContent, nil)
}
