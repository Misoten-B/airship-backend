package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	create "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/create"
	fetchbyid "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_id"
	fetchbyidpublic "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_id_public"
	fetchbyuserid "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_userid"
	statusdone "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/status_done"
	"github.com/Misoten-B/airship-backend/internal/container"
	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	"github.com/Misoten-B/airship-backend/internal/drivers/database"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/ar_assets/dto"
	arassets "github.com/Misoten-B/airship-backend/internal/infrastructure/ar_assets"
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

	var usecaseImpl create.ARAssetsUsecase

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

	input := create.ARAssetsCreateInput{
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
	var usecaseImpl fetchbyid.Usecase
	if config.DevMode {
		usecaseImpl = container.InitializeFetchByIDARAssetsUsecaseForDev()
	} else {
		db, dbErr := database.ConnectDB()
		if dbErr != nil {
			frameworks.ErrorHandling(c, dbErr, http.StatusInternalServerError)
			return
		}

		usecaseImpl = container.InitializeFetchByIDARAssetsUsecaseForProd(db, config)
	}

	input := fetchbyid.Input{
		ID:     id,
		UserID: uid,
	}

	output, err := usecaseImpl.Execute(input)
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

	// リクエスト取得
	id := c.Param("ar_assets_id")
	if id == "" {
		reqErr := errors.New("ar_assets_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
	}

	// ユースケース実行
	var usecaseImpl fetchbyidpublic.Usecase
	if config.DevMode {
		usecaseImpl = container.InitializeFetchByIDPublicARAssetsUsecaseForDev()
	} else {
		db, dbErr := database.ConnectDB()
		if dbErr != nil {
			frameworks.ErrorHandling(c, dbErr, http.StatusInternalServerError)
			return
		}

		usecaseImpl = container.InitializeFetchByIDPublicARAssetsUsecaseForProd(db, config)
	}

	input := fetchbyidpublic.Input{
		ID: id,
	}

	output, err := usecaseImpl.Execute(input)
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

	// ユースケース実行
	var usecaseImpl fetchbyuserid.Usecase

	if config.DevMode {
		usecaseImpl = container.InitializeFetchByUserIDARAssetsUsecaseForDev()
	} else {
		db, dbErr := database.ConnectDB()
		if dbErr != nil {
			frameworks.ErrorHandling(c, dbErr, http.StatusInternalServerError)
			return
		}

		usecaseImpl = container.InitializeFetchByUserIDARAssetsUsecaseForProd(db, config)
	}

	input := fetchbyuserid.Input{
		UserID: uid,
	}

	output, err := usecaseImpl.Execute(input)
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
	responses := []dto.ArAssetsResponse{}
	for _, item := range output.Items {
		responses = append(responses, dto.ArAssetsResponse{
			ID:                   item.ID,
			SpeakingDescription:  item.SpeakingDescription,
			SpeakingAudioPath:    item.SpeakingAudioPath,
			ThreeDimentionalPath: item.ThreeDimentionalPath,
			QrcodeIconImagePath:  item.QrcodeIconImagePath,
		})
	}

	c.JSON(http.StatusOK, responses)
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

// @Tags ArAssets
// @Router /v1/users/ar_assets/{ar_assets_id}/status/done [POST]
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 200 {object} nil
func PostStatusDone(c *gin.Context) {
	// コンテキストから取得
	config, err := frameworks.GetConfig(c)
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
	var usecaseImpl statusdone.Usecase

	if config.DevMode {
		usecaseImpl = statusdone.NewInteractor(
			service.NewMockARAssetsRepository(),
		)
	} else {
		db, dbErr := database.ConnectDB()
		if dbErr != nil {
			frameworks.ErrorHandling(c, dbErr, http.StatusInternalServerError)
			return
		}

		usecaseImpl = statusdone.NewInteractor(
			arassets.NewGormARAssetsRepository(db),
		)
	}

	input := statusdone.Input{
		ID: id,
	}

	err = usecaseImpl.Execute(input)
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
	c.JSON(http.StatusOK, nil)
}

// @Tags ArAssets
// @Router /v1/users/ar_assets/{ar_assets_id}/status/failed [POST]
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 200 {object} nil
func PostStatusFailed(c *gin.Context) {
	// コンテキストから取得

	// リクエスト取得

	// ユースケース実行

	// レスポンス
	c.JSON(http.StatusOK, nil)
}
