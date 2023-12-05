package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/ar_assets/dto"
	"github.com/Misoten-B/airship-backend/internal/database"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	threeservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	arassets "github.com/Misoten-B/airship-backend/internal/infrastructure/ar_assets"
	threedimentionalmodel "github.com/Misoten-B/airship-backend/internal/infrastructure/three_dimentional_model"
	"github.com/Misoten-B/airship-backend/internal/infrastructure/voice"
	"github.com/Misoten-B/airship-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

// @Tags ArAssets
// @Router /v1/users/ar_assets [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param qrcodeIcon formData file false "Image file to be uploaded"
// @Param dto.CreateArAssetsRequest formData dto.CreateArAssetsRequest true "ArAssets"
// @Success 201 {object} nil
// @Header 201 {string} Location "/{ar_assets_id}"
func CreateArAssets(c *gin.Context) {
	// コンテキストから取得
	config, err := frameworks.GetConfig(c)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	uid, err := frameworks.GetUID(c)
	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// リクエスト取得
	request := dto.CreateArAssetsRequest{}
	if err = c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, fileHeader, err := c.Request.FormFile("qrcodeIcon")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 後々DIコンテナから
	var arassetsRepository service.ARAssetsRepository
	var qrCodeImageStorage service.QRCodeImageStorage
	var voiceModelAdapter voiceservice.VoiceModelAdapter
	var voiceService voiceservice.VoiceService
	var threeDimentionalModelService threeservice.ThreeDimentionalModelService

	if config.DevMode {
		arassetsRepository = service.NewMockARAssetsRepository()
		qrCodeImageStorage = service.NewMockQRCodeImageStorage()

		voiceRepo := voiceservice.MockVoiceRepository{}
		voiceService = *voiceservice.NewVoiceService(&voiceRepo)
		voiceModelAdapter = &voiceservice.MockVoiceModelAdapter{}

		threeDimentionalModelRepository := threeservice.NewMockThreeDimentionalModelRepository()
		threeDimentionalModelService = *threeservice.NewThreeDimentionalModelService(threeDimentionalModelRepository)
	} else {
		db, dbErr := database.ConnectDB()
		if dbErr != nil {
			log.Printf("%s", dbErr)
			c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
			return
		}

		arassetsRepository = arassets.NewGormARAssetsRepository(db)
		qrCodeImageStorage = arassets.NewAzureQRCodeImageStorage(config)

		voiceRepo := voice.NewGormVoiceRepository(db)
		voiceService = *voiceservice.NewVoiceService(&voiceRepo)
		voiceModelAdapter = voice.NewExternalAPIVoiceModelAdapter()

		threeDimentionalModelRepository := threedimentionalmodel.NewGormThreeDimentionalModelRepository(db)
		threeDimentionalModelService = *threeservice.NewThreeDimentionalModelService(threeDimentionalModelRepository)
	}

	usecaseImpl := usecase.NewARAssetsUsecaseImpl(
		arassetsRepository,
		qrCodeImageStorage,
		voiceModelAdapter,
		voiceService,
		threeDimentionalModelService,
	)

	// ユースケース実行
	input := usecase.ARAssetsCreateInput{
		UID:                 uid,
		SpeakingDescription: request.SpeakingDescription,
		ThreeDimentionalID:  request.ThreeDimentionalID,
		File:                file,
		FileHeader:          fileHeader,
	}

	output, err := usecaseImpl.Create(input)

	if err != nil {
		log.Printf("%s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Location", fmt.Sprintf("/%s", output.ID))
	c.JSON(http.StatusCreated, nil)
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
