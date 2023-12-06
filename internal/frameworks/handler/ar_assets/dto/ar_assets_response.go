package dto

type ArAssetsResponse struct {
	ID                   string `json:"id" example:"id"`
	SpeakingDescription  string `json:"speaking_description" example:"description"`
	SpeakingAudioPath    string `json:"speaking_audio_path" example:"url"`
	ThreeDimentionalPath string `json:"three_dimentional_path" example:"url"`
	QrcodeIconImagePath  string `json:"qrcode_image_path" example:"url"`
}
