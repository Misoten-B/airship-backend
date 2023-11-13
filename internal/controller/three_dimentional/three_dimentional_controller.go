package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/three_dimentional/dto"
	"github.com/gin-gonic/gin"
)

// @Tags ThreeDimentionalModel
// @Router /users/three_dimentionals [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param ThreeDimentionalModel formData file true "3dmodel file to be uploaded"
// @Success 201 {object} dto.ThreeDimentionalResponse
func CreateThreeDimentional(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	file, fileHeader, err := c.Request.FormFile("ThreeDimentionalModel")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("file: %v", file)
	log.Printf("fileHeader: %v", fileHeader)

	c.Header("Location", fmt.Sprintf("/%s", "1"))
	c.JSON(http.StatusCreated, dto.ThreeDimentionalResponse{
		ID:   "1",
		Path: "https://example.com/3dmodel.tflite",
	})
}

// @Tags ThreeDimentionalModel
// @Router /users/three_dimentionals [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.ThreeDimentionalResponse
func ReadAllThreeDimentional(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	c.JSON(http.StatusOK, []dto.ThreeDimentionalResponse{
		{
			ID:   "1",
			Path: "https://example.com/3dmodel.tflite",
		},
	})
}

// @Tags ThreeDimentionalModel
// @Router /users/three_dimentionals/{three_dimentional_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Success 200 {object} dto.ThreeDimentionalResponse
func ReadThreeDimentionalByID(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
	log.Printf("three_dimentional_id: %s", c.Param("three_dimentional_id"))

	c.JSON(http.StatusOK, dto.ThreeDimentionalResponse{
		ID:   "1",
		Path: "https://example.com/3dmodel.tflite",
	})
}

// @Tags ThreeDimentionalModel
// @Router /users/three_dimentionals/{three_dimentional_id} [PUT]
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
// @Router /users/three_dimentionals/{three_dimentional_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Success 204 {object} nil
func DeleteThreeDimentional(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))
	log.Printf("three_dimentional_id: %s", c.Param("three_dimentional_id"))

	c.JSON(http.StatusNoContent, nil)
}
