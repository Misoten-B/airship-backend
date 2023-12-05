package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/user/dto"
	"github.com/Misoten-B/airship-backend/internal/database"
	"github.com/Misoten-B/airship-backend/internal/database/model"
	"github.com/Misoten-B/airship-backend/internal/frameworks"
	"github.com/gin-gonic/gin"
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
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("body: %v", request)

	user := model.User{
		ID:                uid,
		RecordedVoicePath: "",
		RecordedModelPath: "",
		IsToured:          false,
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// TODO レスポンスをDTOに変換
	c.Header("Location", fmt.Sprintf("/%s", uid))
	c.JSON(http.StatusCreated, dto.UserResponse{
		ID:                user.ID,
		RecordedVoicePath: user.RecordedVoicePath,
		RecordedModelPath: user.RecordedModelPath,
		IsToured:          user.IsToured,
	})
}

// @Tags User
// @Router /v1/users/{user_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} dto.UserResponse
func ReadUserByID(c *gin.Context) {
	uid, err := frameworks.GetUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}
	result := db.First(&user, "id = ?", uid)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// TODO: レスポンスをDTOに変換
	c.Header("Location", fmt.Sprintf("/%s", uid))
	c.JSON(http.StatusOK, dto.UserResponse{
		ID:                user.ID,
		RecordedVoicePath: user.RecordedVoicePath,
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
	if err := c.ShouldBind(&request); err != nil {
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

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		ID:                uid,
		RecordedVoicePath: "",
		RecordedModelPath: "",
		IsToured:          request.IsToured,
	}

	result := db.Model(&user).Updates(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:                user.ID,
		RecordedVoicePath: user.RecordedVoicePath,
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

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}
	result := db.Model(&user).Where("id = ?", uid).Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
