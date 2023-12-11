package dto

import "github.com/Misoten-B/airship-backend/internal/drivers/database/model"

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

func ConvertBC(businesscard model.BusinessCard,
	bcc model.BusinessCardPartsCoordinate,
	bcb model.BusinessCardBackground,
	arassets model.ARAsset) BusinessCardResponse {
	return BusinessCardResponse{
		ID:                          businesscard.ID,
		BusinessCardBackgroundColor: bcb.ColorCode,
		BusinessCardBackgroundImage: bcb.ImagePath,
		BusinessCardName:            businesscard.BusinessCardName,
		ThreeDimentionalModel:       arassets.ThreeDimentionalModel.ModelPath,
		SpeakingDescription:         arassets.SpeakingAsset.Description,
		SpeakingAudioPath:           arassets.SpeakingAsset.AudioPath,
		BusinessCardPartsCoordinate: BusinessCardPartsCoordinate{
			ID:                bcc.ID,
			DisplayNameX:      bcc.DisplayNameX,
			DisplayNameY:      bcc.DisplayNameY,
			CompanyNameX:      bcc.CompanyNameX,
			CompanyNameY:      bcc.CompanyNameY,
			DepartmentX:       bcc.DepartmentX,
			DepartmentY:       bcc.DepartmentY,
			OfficialPositionX: bcc.OfficialPositionX,
			OfficialPositionY: bcc.OfficialPositionY,
			PhoneNumberX:      bcc.PhoneNumberX,
			PhoneNumberY:      bcc.PhoneNumberY,
			EmailX:            bcc.EmailX,
			EmailY:            bcc.EmailY,
			PostalCodeX:       bcc.PostalCodeX,
			PostalCodeY:       bcc.PostalCodeY,
			AddressX:          bcc.AddressX,
			AddressY:          bcc.AddressY,
			QrcodeX:           bcc.QRCodeX,
			QrcodeY:           bcc.QRCodeY,
		},
		DisplayName:      businesscard.DisplayName,
		CompanyName:      businesscard.CompanyName,
		Department:       businesscard.Department,
		OfficialPosition: businesscard.OfficialPosition,
		PhoneNumber:      businesscard.PhoneNumber,
		Email:            businesscard.Email,
		PostalCode:       businesscard.PostalCode,
		Address:          businesscard.Address,
	}
}
