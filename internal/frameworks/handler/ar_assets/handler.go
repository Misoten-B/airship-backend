//nolint:dupl // 時間がないため一旦
package handler

import (
	"errors"
	"fmt"
	"net/http"

	create "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/create"
	deletearassets "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/delete_ar_assets"
	deleteqrcodeicon "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/delete_qr_code_icon"
	fetchbyid "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_id"
	fetchbyidpublic "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_id_public"
	fetchbyuserid "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_userid"
	statusdone "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/status_done"
	updatearassets "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/update_ar_assets"
	"github.com/Misoten-B/airship-backend/internal/container"
	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	threeservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/ar_assets/dto"
	arassetsinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/ar_assets"
	threeinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/three_dimentional_model"
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
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		frameworks.ErrorHandling(c, err, http.StatusBadRequest)
		return
	}

	var usecaseImpl create.ARAssetsUsecase

	// ARAssetsUsecaseの生成
	if config.DevMode {
		usecaseImpl = container.InitializeCreateARAssetsUsecaseForDev()
	} else {
		db, dbErr := frameworks.GetDB(c)
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
		return
	}

	// ユースケース実行
	var usecaseImpl fetchbyid.Usecase
	if config.DevMode {
		usecaseImpl = container.InitializeFetchByIDARAssetsUsecaseForDev()
	} else {
		db, dbErr := frameworks.GetDB(c)
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
		ThreeDimentionalID:   output.ThreeDimentionalID,
		ThreeDimentionalPath: output.ThreeDimentionalPath,
		QrcodeIconImagePath:  output.QrcodeIconImagePath,
		IsCompleted:          output.IsCompleted,
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
		return
	}

	// ユースケース実行
	var usecaseImpl fetchbyidpublic.Usecase
	if config.DevMode {
		usecaseImpl = container.InitializeFetchByIDPublicARAssetsUsecaseForDev()
	} else {
		db, dbErr := frameworks.GetDB(c)
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
		IsCompleted:          output.IsCompleted,
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
		db, dbErr := frameworks.GetDB(c)
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
			ThreeDimentionalID:   item.ThreeDimentionalID,
			ThreeDimentionalPath: item.ThreeDimentionalPath,
			QrcodeIconImagePath:  item.QrcodeIconImagePath,
			IsCompleted:          item.IsCompleted,
		})
	}

	c.JSON(http.StatusOK, responses)
}

// @Tags ArAssets
// @Router /v1/users/ar_assets/{ar_assets_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param dto.UpdateArAssetsRequest formData dto.UpdateArAssetsRequest true "ArAssets"
// @Param ar_assets_id path string true "ArAssets ID"
// @Accept multipart/form-data
// @Param qrcodeIcon formData file false "Image file to be uploaded"
// @Success 204 {object} nil
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

	db, err := frameworks.GetDB(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// リクエスト取得
	id := c.Param("ar_assets_id")
	if id == "" {
		reqErr := errors.New("ar_assets_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
		return
	}

	request := dto.UpdateArAssetsRequest{}
	if err = c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, fileHeader, err := c.Request.FormFile("qrcodeIcon")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		frameworks.ErrorHandling(c, err, http.StatusBadRequest)
		return
	}

	// ユースケース実行
	azureDriver := drivers.NewAzureBlobDriver(config)
	voiceRepo := voice.NewGormVoiceRepository(db)
	threeRepo := threeinfra.NewGormThreeDimentionalModelRepository(db)
	usecaseImpl := updatearassets.NewInteractor(
		db,
		arassetsinfra.NewGormARAssetsRepository(db),
		arassetsinfra.NewAzureQRCodeImageStorage(azureDriver),
		voice.NewExternalAPIVoiceModelAdapter(),
		voiceservice.NewVoiceServiceImpl(voiceRepo),
		threeservice.NewThreeDimentionalModelServiceImpl(threeRepo),
	)

	var qrCodeImage *updatearassets.QRCodeImageInput
	if file != nil {
		qrCodeImage = &updatearassets.QRCodeImageInput{
			File:       file,
			FileHeader: fileHeader,
		}
	}

	input := updatearassets.Input{
		ID:                  id,
		UserID:              uid,
		ThreeDimentionalID:  request.ThreeDimentionalID,
		SpeakingDescription: request.SpeakingDescription,
		QRCodeImage:         qrCodeImage,
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
	c.JSON(http.StatusNoContent, nil)
}

// @Tags ArAssets
// @Router /v1/users/ar_assets/{ar_assets_id}/qr_code_icon [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 204 {object} nil
func DeleteArAssetsQRCodeIcon(c *gin.Context) {
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

	db, err := frameworks.GetDB(c)
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
	usecaseImpl := deleteqrcodeicon.NewInteractor(
		db,
		drivers.NewAzureBlobDriver(config),
	)

	input := deleteqrcodeicon.Input{
		ID:     id,
		UserID: uid,
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
	c.JSON(http.StatusNoContent, nil)
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

	db, err := frameworks.GetDB(c)
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
	usecaseImpl := deletearassets.NewInteractor(
		db,
		drivers.NewAzureBlobDriver(config),
	)

	input := deletearassets.Input{
		ID:     id,
		UserID: uid,
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
		return
	}

	// ユースケース実行
	var usecaseImpl statusdone.Usecase

	if config.DevMode {
		usecaseImpl = statusdone.NewInteractor(
			service.NewMockARAssetsRepository(),
		)
	} else {
		db, dbErr := frameworks.GetDB(c)
		if dbErr != nil {
			frameworks.ErrorHandling(c, dbErr, http.StatusInternalServerError)
			return
		}

		usecaseImpl = statusdone.NewInteractor(
			arassetsinfra.NewGormARAssetsRepository(db),
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
