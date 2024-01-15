package v1

import (
	businesscardpartscoordinates "github.com/Misoten-B/airship-backend/internal/frameworks/router/v1/internal/business_card_parts_coordinates"
	"github.com/Misoten-B/airship-backend/internal/frameworks/router/v1/internal/public"
	"github.com/Misoten-B/airship-backend/internal/frameworks/router/v1/internal/users"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// users
		users.Register(v1)

		// public endpoints
		public.Register(v1)

		// business card parts coordinates
		businesscardpartscoordinates.Register(v1)
	}
}
