package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/create"
	"github.com/Misoten-B/airship-backend/internal/container"
	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/drivers/database"
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
	if config.DevMode {
		usecaseImpl = container.InitializeCreateThreeDimentionalModelUsecaseForDev()
	} else {
		db, dbErr := database.ConnectDB()
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

	log.Printf("config: %v", config)
	log.Printf("uid: %s", uid)

	// ユースケース実行
	//   バリデーション&オブジェクト生成
	//   userIDをもとに3Dモデルのテンプレートとユーザー定義モデルを取得
	//   コンテナのURL生成

	// レスポンス
	c.JSON(http.StatusOK, []dto.ThreeDimentionalResponse{
		{
			ID:   "1",
			Path: "https://example.com/3dmodel.tflite",
		},
	})
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

	log.Printf("config: %v", config)
	log.Printf("uid: %s", uid)

	// リクエスト取得
	log.Printf("three_dimentional_id: %s", c.Param("three_dimentional_id"))
	id := c.Param("three_dimentional_id")
	if id == "" {
		reqErr := errors.New("three_dimentional_id is empty")
		frameworks.ErrorHandling(c, reqErr, http.StatusBadRequest)
		return
	}

	// ユースケース実行
	//   バリデーション&オブジェクト生成
	//   IDをもとに3Dモデルを取得
	//   権限確認（ユーザー定義の場合はそれが自分のものか）
	//   URL生成

	// レスポンス
	c.JSON(http.StatusOK, dto.ThreeDimentionalResponse{
		ID:   "1",
		Path: "https://example.com/3dmodel.tflite",
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
