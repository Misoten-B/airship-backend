package dto

type CreateArAssetsRequest struct {
	SpeakingDescription string `form:"speaking_description" example:"description"`
	ThreeDimentionalID  string `form:"three_dimentional_ID" example:"url"`
}
