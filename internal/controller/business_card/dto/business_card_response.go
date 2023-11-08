package dto

type BusinessCardResponse struct {
	ID string
	// background
	BusinessCardBackgroundColor string `example:"#ffffff"`
	BusinessCardBackgroundImage string `example:"url"`
	// ar assets
	ThreeDimentionalModel string
	SpeakingDescription   string
	SpeakingAudioPath     string
	// business card
	BusinessCardPartsCoordinate string
	BusinessCardName            string
	DisplayName                 string
	CompanyName                 string
	Department                  string
	OfficialPosition            string
	PhoneNumber                 string
	Email                       string
	PostalCode                  string
	Address                     string
	AccessCount                 int
}
