package controller

import (
	"log"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/controller/business_card_background/dto"
	"github.com/gin-gonic/gin"
)

// @Tags BusinessCardBackground
// @Router /user/business_card_background [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param dto.CreateBackgroundRequest body dto.CreateBackgroundRequest true "BusinessCardBackground"
// @Success 201 {object} nil
func CreateBusinessCardBackground(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	request := dto.CreateBackgroundRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("body: %v", request)

	c.JSON(http.StatusCreated, nil)
}

// @Tags BusinessCardBackground
// @Router /user/business_card_background [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 201 {object} []dto.BusinessCardBackgroundResponse
func ReadAllBusinessCardBackground(c *gin.Context) {
	log.Printf("Authorization: %s", c.GetHeader("Authorization"))

	c.JSON(http.StatusOK, []dto.BusinessCardBackgroundResponse{
		{
			ID:                          "1",
			BusinessCardBackgroundColor: "#000000",
			BusinessCardBackgroundImage: "https://example.com/business_card_background_image.png",
		},
	})
}
