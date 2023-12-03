package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/business_card_background/dto"
	"github.com/gin-gonic/gin"
)

// @Tags BusinessCardBackground
// @Router /v1/users/business_card_backgrounds [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param BusinessCardBackgroundImage formData file true "Image file to be uploaded"
// @Param dto.CreateBackgroundRequest formData dto.CreateBackgroundRequest true "BusinessCardBackground"
// @Success 201 {object} dto.BackgroundResponse
func CreateBusinessCardBackground(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	request := dto.CreateBackgroundRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("body: %v", request)

	c.Header("Location", fmt.Sprintf("/%s", "1"))
	c.JSON(http.StatusCreated, dto.BackgroundResponse{
		ID:                          "1",
		BusinessCardBackgroundColor: "#000000",
		BusinessCardBackgroundImage: "https://example.com/business_card_background_image.png",
	})
}

// @Tags BusinessCardBackground
// @Router /v1/users/business_card_backgrounds [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.BackgroundResponse
func ReadAllBusinessCardBackground(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	c.JSON(http.StatusOK, []dto.BackgroundResponse{
		{
			ID:                          "1",
			BusinessCardBackgroundColor: "#000000",
			BusinessCardBackgroundImage: "https://example.com/business_card_background_image.png",
		},
	})
}
