package controller

import "github.com/gin-gonic/gin"

// @Tags ThreeDimentionalModel
// @Router /user/three_dimentional [POST]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Accept multipart/form-data
// @Param ThreeDimentionalModel formData file true "3dmodel file to be uploaded"
// @Success 201 {object} nil
func CreateThreeDimentional(_ *gin.Context) {}

// @Tags ThreeDimentionalModel
// @Router /user/three_dimentional [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Success 200 {object} []dto.ThreeDimentionalResponse
func ReadAllThreeDimentional(_ *gin.Context) {}

// @Tags ThreeDimentionalModel
// @Router /user/three_dimentional/{three_dimentional_id} [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Success 200 {object} dto.ThreeDimentionalResponse
func ReadThreeDimentionalByID(_ *gin.Context) {}

// @Tags ThreeDimentionalModel
// @Router /user/three_dimentional/{three_dimentional_id} [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Accept multipart/form-data
// @Param ThreeDimentionalModel formData file true "3dmodel file to be uploaded"
// @Success 201 {object} nil
func UpdateThreeDimentional(_ *gin.Context) {}

// @Tags ThreeDimentionalModel
// @Router /user/three_dimentional/{three_dimentional_id} [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param three_dimentional_id path string true "ThreeDimentional ID"
// @Success 200 {object} nil
func DeleteThreeDimentional(_ *gin.Context) {}
