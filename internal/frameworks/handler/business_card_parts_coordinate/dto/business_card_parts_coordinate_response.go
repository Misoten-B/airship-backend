package dto

import "github.com/Misoten-B/airship-backend/internal/drivers/database/model"

type BusinessCardPartsCoordinateResponse struct {
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

func ConvertToBCPCResponse(bcpc model.BusinessCardPartsCoordinate) BusinessCardPartsCoordinateResponse {
	return BusinessCardPartsCoordinateResponse{
		ID:                bcpc.ID,
		DisplayNameX:      bcpc.DisplayNameX,
		DisplayNameY:      bcpc.DisplayNameY,
		CompanyNameX:      bcpc.CompanyNameX,
		CompanyNameY:      bcpc.CompanyNameY,
		DepartmentX:       bcpc.DepartmentX,
		DepartmentY:       bcpc.DepartmentY,
		OfficialPositionX: bcpc.OfficialPositionX,
		OfficialPositionY: bcpc.OfficialPositionY,
		PhoneNumberX:      bcpc.PhoneNumberX,
		PhoneNumberY:      bcpc.PhoneNumberY,
		EmailX:            bcpc.EmailX,
		EmailY:            bcpc.EmailY,
		PostalCodeX:       bcpc.PostalCodeX,
		PostalCodeY:       bcpc.PostalCodeY,
		AddressX:          bcpc.AddressX,
		AddressY:          bcpc.AddressY,
		QrcodeX:           bcpc.QRCodeX,
		QrcodeY:           bcpc.QRCodeY,
	}
}
