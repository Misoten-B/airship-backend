package dto

type CreateArAssetsRequest struct {
	SpeakingDescription string `json:"speaking_description" example:"description"`
	ThreeDimentionalID  string `json:"three_dimentional_ID" example:"url"`
	// QrcodeIconImage     string `json:"qrcode_image" extensions:"x-nullable"`
}
