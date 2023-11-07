package dto

type CreateArAssetsRequest struct {
	SpeakingDescription  string `json:"speaking_description" example:"description"`
	SpeakingAudioPath    string `json:"speaking_audio_path" example:"url"`
	ThreeDimentionalPath string `json:"three_dimentional_path" example:"url"`
}
