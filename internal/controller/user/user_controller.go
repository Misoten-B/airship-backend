package controller

import "github.com/gin-gonic/gin"

// @Tags User
// @Router /user [POST]
// @Success 201 {object} dto.UserResponse
// @Param CreateUserRequest body dto.CreateUserRequest true "User ID"
func CreateUser(_ *gin.Context) {}

// @Tags User
// @Router /user/:user_id [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param user_id path string true "User ID"
// @Success 200 {object} dto.UserResponse
func ReadUserByID(_ *gin.Context) {}

// @Tags User
// @Router /user/:user_id [PUT]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param CreateUserRequest body dto.CreateUserRequest true "update user"
// @Param user_id path string true "User ID"
// @Success 200 {object} nil
func UpdateUser(_ *gin.Context) {}

// @Tags User
// @Router /user/:user_id [DELETE]
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer [Firebase JWT Token]"
// @Param user_id path string true "User ID"
// @Success 200 {object} nil
func DeleteUser(_ *gin.Context) {}
