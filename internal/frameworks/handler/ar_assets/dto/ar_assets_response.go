package dto

type ArAssetsResponse struct {
	ID                   string `json:"id" example:"id" binding:"required"`
	SpeakingDescription  string `json:"speakingDescription" example:"description" binding:"required"`
	SpeakingAudioPath    string `json:"speakingAudioPath" example:"url" binding:"required"`
	ThreeDimentionalPath string `json:"threeDimentionalPath" example:"url" binding:"required"`
	QrcodeIconImagePath  string `json:"qrcodeImagePath" example:"url"`
	IsCompleted          bool   `json:"isCompleted" example:"true" binding:"required"`
}
