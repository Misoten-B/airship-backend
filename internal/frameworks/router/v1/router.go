package v1

import (
	arassets "github.com/Misoten-B/airship-backend/internal/frameworks/router/v1/internal/ar_assets"
	businesscardpartscoordinates "github.com/Misoten-B/airship-backend/internal/frameworks/router/v1/internal/business_card_parts_coordinates"
	"github.com/Misoten-B/airship-backend/internal/frameworks/router/v1/internal/users"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// users
		users.Register(v1)
		// ar-asserts
		arassets.Register(v1)
		businesscardpartscoordinates.Register(v1)
	}
}
