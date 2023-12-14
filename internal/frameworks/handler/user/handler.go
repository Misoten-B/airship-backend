package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/Misoten-B/airship-backend/internal/frameworks/handler/user/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Tags User
// @Router /v1/users [POST]
// @Success 201 {object} dto.UserResponse
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param CreateUserRequest body dto.CreateUserRequest true "create user"
func CreateUser(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
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
	log.Printf("body: %v", request)

	user := model.User{
		ID:                uid,
		RecordedModelPath: "",
		IsToured:          false,
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Location", fmt.Sprintf("/%s", uid))
	c.JSON(http.StatusOK, dto.UserResponse{
		ID:                user.ID,
		RecordedModelPath: user.RecordedModelPath,
		IsToured:          user.IsToured,
	})
}

// @Tags User
// @Router /v1/users/{user_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param recorded_voice formData file false "Audio file to be uploaded"
// @Param CreateUserRequest formData dto.CreateUserRequest false "update user"
// @Success 200 {object} dto.UserResponse
func UpdateUser(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
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
	file, fileHeader, err := c.Request.FormFile("recorded_voice")
	log.Printf("file: %v", file)
	log.Printf("fileHeader: %v", fileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	})
}

// @Tags User
// @Router /v1/users/{user_id} [DELETE]
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
