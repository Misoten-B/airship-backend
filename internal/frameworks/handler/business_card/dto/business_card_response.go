package dto

import "github.com/Misoten-B/airship-backend/internal/drivers/database/model"

type BusinessCardResponse struct {
	ID string `json:"id"`
	// background
	BusinessCardBackgroundColor string `json:"color_code" example:"#ffffff"`
	BusinessCardBackgroundImage string `json:"image_path" example:"https://example.com/image.png"`
	// ar assets
	ThreeDimentionalModel string `json:"model_path" example:"https://example.com/model.glb"`
	SpeakingDescription   string `json:"description" example:"This is a description"`
	SpeakingAudioPath     string `json:"audio_path" example:"https://example.com/audio.mp3"`
	// business card
	BusinessCardPartsCoordinate BusinessCardPartsCoordinate `json:"business_card_parts_coordinate"`
	BusinessCardName            string                      `json:"business_card_name" example:"Business Card Name"`
	DisplayName                 string                      `json:"display_name" example:"Display Name"`
	CompanyName                 string                      `json:"company_name" example:"Company Name"`
	Department                  string                      `json:"department" example:"Department"`
	OfficialPosition            string                      `json:"official_position" example:"Official Position"`
	PhoneNumber                 string                      `json:"phone_number" example:"080-1234-5678"`
	Email                       string                      `json:"email" example:"sample@example.com"`
	PostalCode                  string                      `json:"postal_code" example:"123-4567"`
	Address                     string                      `json:"address" example:"Tokyo"`
	AccessCount                 int                         `json:"access_count" example:"0"`
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
