package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/user/dto"
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

	request := dto.CreateUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("body: %v", request)

	c.Header("Location", fmt.Sprintf("/%s", "1"))
	c.JSON(http.StatusCreated, dto.UserResponse{
		ID:                "1",
		RecordedVoicePath: "https://example.com/recorded_voice.mp3",
		RecordedModelPath: "https://example.com/recorded_model.tflite",
		IsToured:          request.IsToured,
	})
}

// @Tags User
// @Router /v1/users/{user_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} dto.UserResponse
func ReadUserByID(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
	log.Printf("user_id: %s", c.Param("user_id"))

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:                "1",
		RecordedVoicePath: "https://example.com/recorded_voice.mp3",
		RecordedModelPath: "https://example.com/recorded_model.tflite",
		IsToured:          false,
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

	request := dto.CreateUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("body: %v", request)

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:                "1",
		RecordedVoicePath: "https://example.com/recorded_voice.mp3",
		RecordedModelPath: "https://example.com/recorded_model.tflite",
		IsToured:          false,
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

	c.JSON(http.StatusNoContent, nil)
}
