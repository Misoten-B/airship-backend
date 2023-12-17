package dto

type CreateArAssetsRequest struct {
	SpeakingDescription string `form:"speakingDescription" example:"description"`
	ThreeDimentionalID  string `form:"threeDimentionalID" example:"url"`
}
