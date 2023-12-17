package dto

type ArAssetsResponse struct {
	ID                   string `json:"id" example:"id"`
	SpeakingDescription  string `json:"speakingDescription" example:"description"`
	SpeakingAudioPath    string `json:"speakingAudioPath" example:"url"`
	ThreeDimentionalPath string `json:"threeDimentionalPath" example:"url"`
	QrcodeIconImagePath  string `json:"qrcodeImagePath" example:"url"`
}
