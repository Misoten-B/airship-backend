package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/config"
	usecase "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/container"
	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/drivers/database"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/ar_assets/dto"
	"github.com/Misoten-B/airship-backend/internal/infrastructure/voice"
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
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	uid, err := frameworks.GetUID(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// リクエスト取得
	request := dto.CreateArAssetsRequest{}
	if err = c.ShouldBind(&request); err != nil {
		frameworks.ErrorHandling(c, err, http.StatusBadRequest)
		return
	}

	file, fileHeader, err := c.Request.FormFile("qrcodeIcon")
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusBadRequest)
		return
	}

	var usecaseImpl usecase.ARAssetsUsecase

	// ARAssetsUsecaseの生成
	if config.DevMode {
		usecaseImpl = container.InitializeCreateARAssetsUsecaseForDev()
	} else {
		db, dbErr := database.ConnectDB()
		if dbErr != nil {
			frameworks.ErrorHandling(c, dbErr, http.StatusInternalServerError)
			return
		}

		usecaseImpl = container.InitializeCreateARAssetsUsecaseForProd(db, config)
	}

	input := usecase.ARAssetsCreateInput{
		UID:                 uid,
		SpeakingDescription: request.SpeakingDescription,
		ThreeDimentionalID:  request.ThreeDimentionalID,
		File:                file,
		FileHeader:          fileHeader,
	}

	output, err := usecaseImpl.Create(input)
	if err != nil {
		var appErr *customerror.ApplicationError

		if errors.As(err, &appErr) {
			frameworks.ErrorHandling(c, err, appErr.StatusCode())
			return
		}
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// レスポンス
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
	// コンテキストから取得
	config, err := frameworks.GetConfig(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	uid, err := frameworks.GetUID(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// リクエスト取得
	id := c.Param("ar_assets_id")
	if id == "" {
		reqErr := errors.New("ar_assets_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
	}

	// ユースケース実行
	usecaseImpl, err := newUsecase(config)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	input := usecase.ARAssetsFetchByIDInput{
		ID:  id,
		UID: uid,
	}

	output, err := usecaseImpl.FetchByID(input)
	if err != nil {
		var appErr *customerror.ApplicationError

		if errors.As(err, &appErr) {
			frameworks.ErrorHandling(c, err, appErr.StatusCode())
			return
		}
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusOK, dto.ArAssetsResponse{
		ID:                   output.ID,
		SpeakingDescription:  output.SpeakingDescription,
		SpeakingAudioPath:    output.SpeakingAudioPath,
		ThreeDimentionalPath: output.ThreeDimentionalPath,
		QrcodeIconImagePath:  output.QrcodeIconImagePath,
	})
}

// @Tags ArAssets
// @Router /v1/ar_assets/{ar_assets_id} [GET]
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 200 {object} dto.ArAssetsResponse
func ReadArAssetsByIDPublic(c *gin.Context) {
	// コンテキストから取得
	config, err := frameworks.GetConfig(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	log.Printf("config: %v", config)

	// リクエスト取得
	id := c.Param("ar_assets_id")
	if id == "" {
		reqErr := errors.New("ar_assets_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
	}

	// ユースケース実行

	// レスポンス
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
	// コンテキストから取得
	config, err := frameworks.GetConfig(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	uid, err := frameworks.GetUID(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	log.Printf("config: %v, uid: %v", config, uid)

	// リクエスト取得

	// ユースケース実行

	// レスポンス
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
	/// コンテキストから取得
	config, err := frameworks.GetConfig(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	uid, err := frameworks.GetUID(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	log.Printf("config: %v, uid: %v", config, uid)

	// リクエスト取得
	id := c.Param("ar_assets_id")
	if id == "" {
		reqErr := errors.New("ar_assets_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
	}

	request := dto.CreateArAssetsRequest{}
	if err = c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("formData: %v", request)

	// ユースケース実行

	// レスポンス
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
	// コンテキストから取得
	config, err := frameworks.GetConfig(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	uid, err := frameworks.GetUID(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	log.Printf("config: %v, uid: %v", config, uid)

	// リクエスト取得
	id := c.Param("ar_assets_id")
	if id == "" {
		reqErr := errors.New("ar_assets_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
	}

	// ユースケース実行

	// レスポンス
	c.JSON(http.StatusNoContent, nil)
}

// TODO: 後々DIコンテナなどから
func newUsecase(config *config.Config) (*usecase.ARAssetsUsecaseImpl, error) {
	if config.DevMode {
		usecaseImpl, err := newUsecaseDev()
		if err != nil {
			return nil, err
		}
		return usecaseImpl, nil
	}

	usecaseImpl, err := newUsecaseProd(config)
	if err != nil {
		return nil, err
	}
	return usecaseImpl, nil
}

func newUsecaseProd(config *config.Config) (*usecase.ARAssetsUsecaseImpl, error) {
	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	voiceRepo := voice.NewGormVoiceRepository(db)
	tmodelRepo := threedimentionalmodel.NewGormThreeDimentionalModelRepository(db)

	usecase, err := usecase.NewARAssetsUsecase(
		usecase.WithGormARAssetsRepository(db),
		usecase.WithAzureQRCodeImageStorage(config),
		usecase.WithExternalAPIVoiceModelAdapter(),
		usecase.WithVoiceServiceImpl(voiceRepo),
		usecase.WithThreeDimentionalModelServiceImpl(tmodelRepo),
	)
	if err != nil {
		return nil, err
	}

	return usecase, nil
}

func newUsecaseDev() (*usecase.ARAssetsUsecaseImpl, error) {
	voiceRepo := voiceservice.NewMockVoiceRepository()
	tmodelRepo := threeservice.NewMockThreeDimentionalModelRepository()

	db, err := database.ConnectDB()
	if err != nil {
		return nil, err
	}

	usecaseImpl, err := usecase.NewARAssetsUsecase(
		// usecase.WithMockARAssetsRepository(),
		usecase.WithGormARAssetsRepository(db),
		usecase.WithMockQRCodeImageStorage(),
		usecase.WithMockVoiceModelAdapter(),
		usecase.WithVoiceServiceImpl(voiceRepo),
		usecase.WithThreeDimentionalModelServiceImpl(tmodelRepo),
	)
	if err != nil {
		return nil, err
	}

	return usecaseImpl, nil
}
