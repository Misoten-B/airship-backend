package router

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	healthCheck(r)
	openapi(r)

	// bussiness
	user(r)
	businessCardPartsCoordinate(r)
}
