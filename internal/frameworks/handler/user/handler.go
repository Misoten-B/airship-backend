package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/drivers/config"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/user/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

const (
	containerName = "train-sounds"
)

// @Tags User
// @Router /v1/users [POST]
// @Success 201 {object} dto.UserResponse
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param CreateUserRequest body dto.CreateUserRequest true "create user"
func CreateUser(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: リクエストのバリデーション
	request := dto.CreateUserRequest{}
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		ID:                uid,
		RecordedModelPath: "",
		IsToured:          false,
		Status:            model.GormStatusNone,
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Location", fmt.Sprintf("/%s", uid))
	c.JSON(http.StatusCreated, dto.UserResponse{
		ID:                user.ID,
		RecordedModelPath: user.RecordedModelPath,
		IsToured:          user.IsToured,
		Status:            user.Status,
	})
}

// @Tags User
// @Router /v1/users/myprofile [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} dto.UserResponse
func ReadUserByID(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}
	err = db.First(&user, "id = ?", uid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Location", fmt.Sprintf("/%s", uid))
	c.JSON(http.StatusOK, dto.UserResponse{
		ID:                user.ID,
		RecordedModelPath: user.RecordedModelPath,
		IsToured:          user.IsToured,
		Status:            user.Status,
	})
}

// @Tags User
// @Router /v1/users [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param recorded_voice formData file true "Audio file to be uploaded"
// @Param CreateUserRequest formData dto.CreateUserRequest false "update user"
// @Success 200 {object} dto.UserResponse
func UpdateUser(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// TODO: リクエストのバリデーション
	request := dto.CreateUserRequest{}
	if err = c.ShouldBind(&request); err != nil {
		log.Print("aaaa")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: AI側に送信
	formFile, fileHeader, err := c.Request.FormFile("recorded_voice")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appConfig, err := config.GetConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ab := drivers.NewAzureBlobDriver(appConfig)
	castedFile := file.NewMyFile(formFile, fileHeader)
	castedFile.FileHeader().Filename = fmt.Sprintf("%s%s", uid, filepath.Ext(castedFile.FileHeader().Filename))
	if err = ab.SaveBlob(containerName, castedFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	url := fmt.Sprintf("https://airship-ml.japaneast.cloudapp.azure.com/voice-model/%s", uid)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"train_sound_file_name": castedFile.FileHeader().Filename,
			"language":              "ja",
		}).
		Post(url)
	if resp.StatusCode() != http.StatusOK {
		c.JSON(resp.StatusCode(), gin.H{"error": "AI server error"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		ID:                uid,
		RecordedModelPath: "",
		IsToured:          request.IsToured,
		Status:            model.GormStatusCompleted,
	}

	err = db.Model(&user).Updates(user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:                user.ID,
		RecordedModelPath: user.RecordedModelPath,
		IsToured:          user.IsToured,
		Status:            user.Status,
	})
}

// @Tags User
// @Router /v1/users [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 204 {object} nil
func DeleteUser(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
	log.Printf("user_id: %s", c.Param("user_id"))

	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := frameworks.GetDB(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}
	err = db.Model(&user).Where("id = ?", uid).Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Tags User
// @Router /v1/users/{user_id}/voice_model/status/done [POST]
// @Param user_id path string true "User ID"
// @Success 200 {object} nil
func PostVoiceModelStatusDone(c *gin.Context) {
	// コンテキストから取得

	// リクエスト取得
	userID := c.Param("user_id")
	if userID == "" {
		frameworks.ErrorHandling(
			c,
			errors.New("user_id is empty"),
			http.StatusBadRequest,
		)
		return
	}

	// ユースケース実行

	// データベース接続
	db, err := frameworks.GetDB(c)
	if err != nil {
		frameworks.ErrorHandling(
			c,
			err,
			http.StatusInternalServerError,
		)
		return
	}

	// ユーザーの取得
	user := model.User{}
	err = db.First(&user, "id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			frameworks.ErrorHandling(
				c,
				err,
				http.StatusNotFound,
			)
			return
		}

		frameworks.ErrorHandling(
			c,
			err,
			http.StatusInternalServerError,
		)
	}

	// ユーザーの状態を更新
	err = db.Model(&user).Update("status", model.GormStatusCompleted).Error
	if err != nil {
		frameworks.ErrorHandling(
			c,
			err,
			http.StatusInternalServerError,
		)
		return
	}

	// レスポンス
	c.JSON(http.StatusOK, nil)
}

// @Tags User
// @Router /v1/users/{user_id}/voice_model/status/failed [POST]
// @Param user_id path string true "User ID"
// @Success 200 {object} nil
func PostVoiceModelStatusFailed(c *gin.Context) {
	// コンテキストから取得

	// リクエスト取得

	// ユースケース実行

	// レスポンス
	c.JSON(http.StatusOK, nil)
}
