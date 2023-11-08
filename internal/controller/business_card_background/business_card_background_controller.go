package controller

import "github.com/gin-gonic/gin"

// @Tags BusinessCardBackground
// @Router /user/business_card_background [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param dto.CreateBackgroundRequest body dto.CreateBackgroundRequest true "BusinessCardBackground"
// @Success 201 {object} nil
func CreateBusinessCardBackground(_ *gin.Context) {}

// @Tags BusinessCardBackground
// @Router /user/business_card_background [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 201 {object} []dto.BusinessCardBackgroundResponse
func ReadAllBusinessCardBackground(_ *gin.Context) {}
