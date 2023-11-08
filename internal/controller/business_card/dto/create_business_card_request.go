package dto

type CreateBusinessCardRequest struct {
	// background
	BusinessCardBackgroundID *string `json:"business_card_background_id" example:"id" extensions:"x-nullable"`
	// ar assets
	ArAssetsID *string `json:"ar_assets_id,omitempty" example:"ar_assets_id" extensions:"x-nullable"`
	// business card
	BusinessCardPartsCoordinateID *string `json:"business_card_parts_coordinate" example:"id" extensions:"x-nullable"`
	BusinessCardName              *string `json:"business_card_name,omitempty" example:"会社" extensions:"x-nullable"`
	DisplayName                   *string `json:"display_name,omitempty" example:"名前" extensions:"x-nullable"`
	CompanyName                   *string `json:"company_name,omitempty" example:"会社名" extensions:"x-nullable"`
	Department                    *string `json:"department,omitempty" example:"部署" extensions:"x-nullable"`
	OfficialPosition              *string `json:"official_position,omitempty" example:"役職" extensions:"x-nullable"`
	PhoneNumber                   *string `json:"phone_number,omitempty" example:"090-1234-5678" extensions:"x-nullable"`
	Email                         *string `json:"email,omitempty" example:"sample@example.com" extensions:"x-nullable"`
	PostalCode                    *string `json:"postal_code,omitempty" example:"123-4567" extensions:"x-nullable"`
	Address                       *string `json:"address,omitempty" example:"東京都渋谷区神南1-1-1" extensions:"x-nullable"`
	QrcodeImageIconPath           *string `json:"qrcode_image_path,omitempty" example:"url" extensions:"x-nullable"`
}
