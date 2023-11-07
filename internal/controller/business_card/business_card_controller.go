package controller

import "github.com/gin-gonic/gin"

// @Tags BusinessCard
// @Router /user/business_card [POST]
// @Param CreateBusinessCardRequest body dto.CreateBusinessCardRequest true "BusinessCard"
// @Success 201 {object} nil
func CreateBusinessCard(_ *gin.Context) {}

// @Tags BusinessCard
// @Router /user/business_card [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.BusinessCardResponse
func ReadAllBusinessCard(_ *gin.Context) {}

// @Tags BusinessCard
// @Router /user/business_card/{business_card_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Success 200 {object} dto.BusinessCardResponse
func ReadBusinessCardByID(_ *gin.Context) {}

// @Tags BusinessCard
// @Router /user/business_card/{business_card_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Param CreateBusinessCardRequest body dto.CreateBusinessCardRequest true "BusinessCard"
// @Success 201 {object} nil
func UpdateBusinessCard(_ *gin.Context) {}

// @Tags BusinessCard
// @Router /user/business_card/{business_card_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param business_card_id path string true "BusinessCard ID"
// @Success 200 {object} nil
func DeleteBusinessCard(_gin *gin.Context) {}
