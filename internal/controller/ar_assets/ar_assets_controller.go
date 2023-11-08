package qr

import "github.com/gin-gonic/gin"

// @Tags ArAssets
// @Router /user/ar_assets [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param qrcodeImage formData file true "Image file to be uploaded"
// @Param dto.CreateArAssetsRequest body dto.CreateArAssetsRequest true "ArAssets"
// @Success 201 {object} nil
func CreateArAssets(_ *gin.Context) {}

// @Tags ArAssets
// @Router /user/ar_assets/{ar_assets_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 201 {object} dto.ArAssetsResponse
func ReadArAssetsByID(_ *gin.Context) {}

// @Tags ArAssets
// @Router /ar_assets/{ar_assets_id} [GET]
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 201 {object} dto.ArAssetsResponse
func ReadArAssetsByIDPublic(_ *gin.Context) {}

// @Tags ArAssets
// @Router /user/ar_assets [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.ArAssetsResponse
func ReadAllArAssets(_ *gin.Context) {}

// @Tags ArAssets
// @Router /user/ar_assets/{ar_assets_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param dto.CreateArAssetsRequest body dto.CreateArAssetsRequest true "ArAssets"
// @Param ar_assets_id path string true "ArAssets ID"
// @Accept multipart/form-data
// @Param qrcodeIcon formData file false "Image file to be uploaded"
// @Success 201 {object} nil
func UpdateArAssets(_ *gin.Context) {}

// @Tags ArAssets
// @Router /user/ar_assets/{ar_assets_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param ar_assets_id path string true "ArAssets ID"
// @Success 200 {object} nil
func DeleteArAssets(_ *gin.Context) {}
