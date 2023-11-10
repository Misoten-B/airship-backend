package controller

import (
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/user/dto"
	"github.com/gin-gonic/gin"
)

// @Tags User
// @Router /user [POST]
// @Success 201 {object} dto.UserResponse
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param CreateUserRequest body dto.CreateUserRequest true "User ID"
func CreateUser(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	request := dto.CreateUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("body: %v", request)

	// c.Header("Location", c.Request.Host+c.Request.URL.Path+"1")
	c.JSON(http.StatusCreated, dto.UserResponse{
		ID:                "1",
		RecordedVoicePath: "https://example.com/recorded_voice.mp3",
		RecordedModelPath: "https://example.com/recorded_model.tflite",
		IsToured:          request.IsToured,
	})
}

// @Tags User
// @Router /user/{user_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param user_id path string true "User ID"
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
// @Router /user/{user_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param CreateUserRequest body dto.CreateUserRequest true "update user"
// @Param user_id path string true "User ID"
// @Success 200 {object} nil
func UpdateUser(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	request := dto.CreateUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("body: %v", request)

	c.JSON(http.StatusOK, nil)
}

// @Tags User
// @Router /user/{user_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param user_id path string true "User ID"
// @Success 200 {object} nil
func DeleteUser(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
	log.Printf("user_id: %s", c.Param("user_id"))

	c.JSON(http.StatusOK, nil)
}
