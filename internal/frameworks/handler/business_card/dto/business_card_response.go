package dto

import "github.com/Misoten-B/airship-backend/internal/drivers/database/model"

type BusinessCardResponse struct {
	ID string `json:"id" binding:"required"`
	// background
	BusinessCardBackgroundColor string `json:"businessCardBackgroundColor" example:"#ffffff" binding:"required"`
	BusinessCardBackgroundImage string `json:"businessCardBackgroundImage" example:"https://example.com/image.png"`
	// ar assets
	ThreeDimentionalModel string `json:"threeDimentionalModel" example:"https://example.com/model.glb" binding:"required"`
	SpeakingDescription   string `json:"speakingDescription" example:"This is a description" binding:"required"`
	SpeakingAudioPath     string `json:"speakingAudioPath" example:"https://example.com/audio.mp3" binding:"required"`
	// business card
	BusinessCardPartsCoordinate BusinessCardPartsCoordinate `json:"businessCardPartsCoordinate" binding:"required"`
	BusinessCardName            string                      `json:"businessCardName" example:"Business Card Name"`
	DisplayName                 string                      `json:"displayName" example:"Display Name" binding:"required"`
	CompanyName                 string                      `json:"companyName" example:"Company Name"`
	Department                  string                      `json:"department" example:"Department"`
	OfficialPosition            string                      `json:"officialPosition" example:"Official Position"`
	PhoneNumber                 string                      `json:"phoneNumber" example:"080-1234-5678"`
	Email                       string                      `json:"email" example:"sample@example.com"`
	PostalCode                  string                      `json:"postalCode" example:"123-4567"`
	Address                     string                      `json:"address" example:"Tokyo"`
	AccessCount                 int                         `json:"accessCount" example:"0"`
}

type BusinessCardPartsCoordinate struct {
	ID                string `json:"id"`
	DisplayNameX      int    `json:"displayNameX"`
	DisplayNameY      int    `json:"displayNameY"`
	CompanyNameX      int    `json:"companyNameX"`
	CompanyNameY      int    `json:"companyNameY"`
	DepartmentX       int    `json:"departmentX"`
	DepartmentY       int    `json:"departmentY"`
	OfficialPositionX int    `json:"officialPositionX"`
	OfficialPositionY int    `json:"officialPositionY"`
	PhoneNumberX      int    `json:"phoneNumberX"`
	PhoneNumberY      int    `json:"phoneNumberY"`
	EmailX            int    `json:"emailX"`
	EmailY            int    `json:"emailY"`
	PostalCodeX       int    `json:"postalCodeX"`
	PostalCodeY       int    `json:"postalCodeY"`
	AddressX          int    `json:"addressX"`
	AddressY          int    `json:"addressY"`
	QrcodeX           int    `json:"qrcodeX"`
	QrcodeY           int    `json:"qrcodeY"`
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
