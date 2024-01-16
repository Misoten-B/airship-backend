package dto

type UpdateArAssetsRequest struct {
	SpeakingDescription string `form:"speakingDescription" example:"description" binding:"required"`
	ThreeDimentionalID  string `form:"threeDimentionalID" example:"url" binding:"required"`
}
