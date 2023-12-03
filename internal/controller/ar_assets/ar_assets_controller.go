package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Misoten-B/airship-backend/internal/controller/ar_assets/dto"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/id"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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
	log.Printf("config: %v", config)
	log.Printf("uid: %s", uid)

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

	// バリデーション
	ext := filepath.Ext(fileHeader.Filename)
	blobID, err := id.NewID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	blobName := fmt.Sprintf("%s%s", blobID.String(), ext)

	// AI側へリクエスト
	//  - データベースから取得
	//  - 音声モデル取得
	//  - AI側へリクエスト
	if !config.DevMode { // 開発用。AI側が完了次第、実装
		client := resty.New()
		_, err = client.R().
			SetHeader("Content-Type", "application/json").
			Post("http://localhost:8080/example/ai")
		if err != nil {
			log.Printf("failed to request ai: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// QRコードアイコン画像保存
	if !config.DevMode { // 開発用。今後抽出する
		ctx := context.Background()

		serviceClient, azureError := azblob.NewClientFromConnectionString(config.AzureBlobStorageConnectionString, nil)
		if azureError != nil {
			log.Printf("failed to create service client: %s", azureError)
			c.JSON(http.StatusInternalServerError, gin.H{"error": azureError.Error()})
			return
		}

		_, azureError = serviceClient.UploadStream(ctx, "images", blobName, file, &azblob.UploadStreamOptions{})
		if azureError != nil {
			log.Printf("failed to upload stream: %s", azureError)
			c.JSON(http.StatusInternalServerError, gin.H{"error": azureError.Error()})
			return
		}
	}

	// データベース保存

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
