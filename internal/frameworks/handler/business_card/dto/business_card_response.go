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
	BusinessCardPartsCoordinate BusinessCardPartsCoordinate
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

type BusinessCardPartsCoordinate struct {
	ID                string `json:"id"`
	DisplayNameX      int    `json:"display_name_x"`
	DisplayNameY      int    `json:"display_name_y"`
	CompanyNameX      int    `json:"company_name_x"`
	CompanyNameY      int    `json:"company_name_y"`
	DepartmentX       int    `json:"department_x"`
	DepartmentY       int    `json:"department_y"`
	OfficialPositionX int    `json:"official_position_x"`
	OfficialPositionY int    `json:"official_position_y"`
	PhoneNumberX      int    `json:"phone_number_x"`
	PhoneNumberY      int    `json:"phone_number_y"`
	EmailX            int    `json:"email_x"`
	EmailY            int    `json:"email_y"`
	PostalCodeX       int    `json:"postal_code_x"`
	PostalCodeY       int    `json:"postal_code_y"`
	AddressX          int    `json:"address_x"`
	AddressY          int    `json:"address_y"`
	QrcodeX           int    `json:"qrcode_x"`
	QrcodeY           int    `json:"qrcode_y"`
}
