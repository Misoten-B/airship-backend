package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/create"
	fetchbyid "github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/fetch_by_id"
	fetchbyuserid "github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/fetch_by_userid"
	"github.com/Misoten-B/airship-backend/internal/container"
	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/three_dimentional/dto"
	"github.com/gin-gonic/gin"
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
	// if config.DevMode {
	if false {
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
	// if config.DevMode {
	if false {
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
	// if config.DevMode {
	if false {
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
// @Success 200 {object} dto.ThreeDimentionalResponse
func UpdateThreeDimentional(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	file, fileHeader, err := c.Request.FormFile("ThreeDimentionalModel")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("file: %v", file)
	log.Printf("fileHeader: %v", fileHeader)

	c.JSON(http.StatusCreated, dto.ThreeDimentionalResponse{
		ID:   "1",
		Path: "https://example.com/3dmodel.tflite",
	})
}

// @Tags ThreeDimentionalModel
// @Router /v1/users/three_dimentionals/{three_dimentional_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Success 204 {object} nil
func DeleteThreeDimentional(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
	log.Printf("three_dimentional_id: %s", c.Param("three_dimentional_id"))

	c.JSON(http.StatusNoContent, nil)
}
