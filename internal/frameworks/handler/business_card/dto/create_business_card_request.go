package dto

type CreateBusinessCardRequest struct {
	// background
	BusinessCardBackgroundID string `form:"businessCardBackgroundId" example:"id" binding:"required"`
	// ar assets
	ArAssetsID string `form:"arAssetsId,omitempty" example:"ar_assets_id" binding:"required"`
	// business card
	BusinessCardPartsCoordinateID string `form:"businessCardPartsCoordinateId" example:"id" binding:"required"`
	BusinessCardName              string `form:"businessCardName,omitempty" example:"会社"`
	DisplayName                   string `form:"displayName,omitempty" example:"名前"`
	CompanyName                   string `form:"companyName,omitempty" example:"会社名"`
	Department                    string `form:"department,omitempty" example:"部署"`
	OfficialPosition              string `form:"officialPosition,omitempty" example:"役職"`
	PhoneNumber                   string `form:"phoneNumber,omitempty" example:"090-1234-5678"`
	Email                         string `form:"email,omitempty" example:"sample@example.com"`
	PostalCode                    string `form:"postalCode,omitempty" example:"123-4567"`
	Address                       string `form:"address,omitempty" example:"東京都渋谷区神南1-1-1"`
}
