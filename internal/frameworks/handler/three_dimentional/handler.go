//nolint:funlen,cyclop // 時間がないため一旦
package handler

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/create"
	fetchbyid "github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/fetch_by_id"
	fetchbyuserid "github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/fetch_by_userid"
	"github.com/Misoten-B/airship-backend/internal/container"
	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	myfilemod "github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/three_dimentional/dto"
	threeinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/three_dimentional_model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Tags ThreeDimentionalModel
// @Router /v1/users/three_dimentionals [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param ThreeDimentionalModel formData file true "3dmodel file to be uploaded"
// @Success 201 {object} nil
// @Header 201 {string} Location "/{three_dimentional_id}"
func CreateThreeDimentional(c *gin.Context) {
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
	file, fileHeader, err := c.Request.FormFile("ThreeDimentionalModel")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユースケース実行
	var usecaseImpl create.Usecase
	if config.DevMode {
		usecaseImpl = container.InitializeCreateThreeDimentionalModelUsecaseForDev()
	} else {
		db, dbErr := frameworks.GetDB(c)
		if dbErr != nil {
			frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
			return
		}

		usecaseImpl = container.InitializeCreateThreeDimentionalModelUsecaseForProd(db, config)
	}

	input := create.Input{
		UserID:     uid,
		File:       file,
		FileHeader: fileHeader,
	}

	output, err := usecaseImpl.Execute(input)
	if err != nil {
		var appErr *customerror.ApplicationError

		if errors.As(err, &appErr) {
			frameworks.ErrorHandling(c, appErr, appErr.StatusCode())
			return
		}
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.Header("Location", fmt.Sprintf("/%s", output.ID))
	c.JSON(http.StatusCreated, nil)
}

// @Tags ThreeDimentionalModel
// @Router /v1/users/three_dimentionals [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.ThreeDimentionalResponse
func ReadAllThreeDimentional(c *gin.Context) {
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
		usecaseImpl = container.InitializeFetchByUserIDThreeDimentionalModelUsecaseForDev()
	} else {
		db, dbErr := frameworks.GetDB(c)
		if dbErr != nil {
			frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
			return
		}

		usecaseImpl = container.InitializeFetchByUserIDThreeDimentionalModelUsecaseForProd(db, config)
	}

	input := fetchbyuserid.Input{
		UserID: uid,
	}

	output, err := usecaseImpl.Execute(input)
	if err != nil {
		var appErr *customerror.ApplicationError

		if errors.As(err, &appErr) {
			frameworks.ErrorHandling(c, appErr, appErr.StatusCode())
			return
		}

		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// レスポンス
	responses := []dto.ThreeDimentionalResponse{}
	for _, item := range output.Items {
		responses = append(responses, dto.ThreeDimentionalResponse{
			ID:   item.ID,
			Path: item.Path,
		})
	}

	c.JSON(http.StatusOK, responses)
}

// @Tags ThreeDimentionalModel
// @Router /v1/users/three_dimentionals/{three_dimentional_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Success 200 {object} dto.ThreeDimentionalResponse
func ReadThreeDimentionalByID(c *gin.Context) {
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
	id := c.Param("three_dimentional_id")
	if id == "" {
		reqErr := errors.New("three_dimentional_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
		return
	}

	// ユースケース実行
	var usecaseImpl fetchbyid.Usecase
	if config.DevMode {
		usecaseImpl = container.InitializeFetchByIDThreeDimentionalModelUsecaseForDev()
	} else {
		db, dbErr := frameworks.GetDB(c)
		if dbErr != nil {
			frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
			return
		}

		usecaseImpl = container.InitializeFetchByIDThreeDimentionalModelUsecaseForProd(db, config)
	}

	input := fetchbyid.Input{
		ID:     id,
		UserID: uid,
	}

	output, err := usecaseImpl.Execute(input)
	if err != nil {
		var appErr *customerror.ApplicationError

		if errors.As(err, &appErr) {
			frameworks.ErrorHandling(c, appErr, appErr.StatusCode())
			return
		}

		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusOK, dto.ThreeDimentionalResponse{
		ID:   output.ID,
		Path: output.Path,
	})
}

// @Tags ThreeDimentionalModel
// @Router /v1/users/three_dimentionals/{three_dimentional_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Accept multipart/form-data
// @Param ThreeDimentionalModel formData file true "3dmodel file to be uploaded"
// @Success 204 {object} nil
func UpdateThreeDimentional(c *gin.Context) {
	// FIXME: ベタ書き

	// コンテキストから取得
	config, err := frameworks.GetConfig(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	userID, err := frameworks.GetUID(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// リクエスト受け取り
	id := c.Param("three_dimentional_id")
	if id == "" {
		reqErr := errors.New("three_dimentional_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
		return
	}

	file, fileHeader, err := c.Request.FormFile("ThreeDimentionalModel")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// AzureDriverの仕様上、一旦ここでファイル名を変更しています。
	fileName := fmt.Sprintf("%s%s", id, filepath.Ext(fileHeader.Filename))
	fileHeader.Filename = fileName

	// ユースケース実行

	// 3Dモデル取得
	var personalTDM model.PersonalThreeDimentionalModel

	if err = db.Preload("ThreeDimentionalModel").
		First(&personalTDM, "three_dimentional_model_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			frameworks.ErrorHandling(c, err, http.StatusNotFound)
			return
		}
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// 権限確認
	if personalTDM.UserID != userID {
		frameworks.ErrorHandling(
			c,
			errors.New("user does not have permission to update this three dimentional model"),
			http.StatusForbidden,
		)
		return
	}

	// ストレージに保存
	azureDriver := drivers.NewAzureBlobDriver(config)

	myFile := myfilemod.NewMyFile(file, fileHeader)
	if err = azureDriver.SaveBlob("three-dimentional-models", myFile); err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// 旧版削除
	if err = azureDriver.DeleteBlob("three-dimentional-models", personalTDM.ThreeDimentionalModel.ModelPath); err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// データベース更新
	if err = db.Model(&personalTDM.ThreeDimentionalModel).Update("model_path", fileName).Error; err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusNoContent, nil)
}

// @Tags ThreeDimentionalModel
// @Router /v1/users/three_dimentionals/{three_dimentional_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Success 204 {object} nil
func DeleteThreeDimentional(c *gin.Context) {
	// FIXME: ベタ書き

	// コンテキストから取得
	config, err := frameworks.GetConfig(c)
	if err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	userID, err := frameworks.GetUID(c)
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
	id := c.Param("three_dimentional_id")
	if id == "" {
		reqErr := errors.New("three_dimentional_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
		return
	}

	// ユースケース実行

	// 3Dモデル取得
	tdmRepo := threeinfra.NewGormThreeDimentionalModelRepository(db)
	tdm, err := tdmRepo.FindByID(shared.ReconstructID(id))
	if err != nil {
		if errors.Is(err, tdmservice.ErrThreeDimentionalModelNotFound) {
			frameworks.ErrorHandling(c, err, http.StatusNotFound)
			return
		}
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// 権限確認
	if tdm.IsTemplate() {
		frameworks.ErrorHandling(c, errors.New("cannot delete template"), http.StatusForbidden)
		return
	}

	if tdm.UserID() != userID {
		frameworks.ErrorHandling(
			c,
			errors.New("user does not have permission to delete this three dimentional model"),
			http.StatusForbidden,
		)
		return
	}

	// データベースから削除
	// if err = db.Delete(&tdm).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrForeignKeyViolated) {
	// 		frameworks.ErrorHandling(
	// 			c,
	// 			errors.New("failed to delete three dimentional model: foreign key violated"),
	// 			http.StatusBadRequest,
	// 		)
	// 		return
	// 	}
	// 	frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
	// }

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Delete(
		&model.PersonalThreeDimentionalModel{},
		"three_dimentional_model_id = ?",
		tdm.ID()).Error; err != nil {
		tx.Rollback()
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	if err = tx.Delete(&model.ThreeDimentionalModel{}, "id = ?", tdm.ID()).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			frameworks.ErrorHandling(
				c,
				errors.New("failed to delete three dimentional model: foreign key violated"),
				http.StatusBadRequest,
			)
			return
		}
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	if err = tx.Commit().Error; err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// ストレージから削除
	azureDriver := drivers.NewAzureBlobDriver(config)
	if err = azureDriver.DeleteBlob("three-dimentional-models", tdm.Path()); err != nil {
		frameworks.ErrorHandling(c, err, http.StatusInternalServerError)
		return
	}

	// レスポンス
	c.JSON(http.StatusNoContent, nil)
}
