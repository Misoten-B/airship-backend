package dto

type CreateBusinessCardRequest struct {
	// background
	BusinessCardBackgroundID string `form:"businessCardBackgroundId" example:"id" extensions:"x-nullable"`
	// ar assets
	ArAssetsID string `form:"arAssetsId,omitempty" example:"ar_assets_id" extensions:"x-nullable"`
	// business card
	BusinessCardPartsCoordinateID string `form:"businessCardPartsCoordinateId" example:"id" extensions:"x-nullable"`
	BusinessCardName              string `form:"businessCardName,omitempty" example:"会社" extensions:"x-nullable"`
	DisplayName                   string `form:"displayName,omitempty" example:"名前" extensions:"x-nullable"`
	CompanyName                   string `form:"companyName,omitempty" example:"会社名" extensions:"x-nullable"`
	Department                    string `form:"department,omitempty" example:"部署" extensions:"x-nullable"`
	OfficialPosition              string `form:"officialPosition,omitempty" example:"役職" extensions:"x-nullable"`
	PhoneNumber                   string `form:"phoneNumber,omitempty" example:"090-1234-5678" extensions:"x-nullable"`
	Email                         string `form:"email,omitempty" example:"sample@example.com" extensions:"x-nullable"`
	PostalCode                    string `form:"postalCode,omitempty" example:"123-4567" extensions:"x-nullable"`
	Address                       string `form:"address,omitempty" example:"東京都渋谷区神南1-1-1" extensions:"x-nullable"`
}
