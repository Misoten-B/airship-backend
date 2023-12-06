package dto

import "github.com/Misoten-B/airship-backend/internal/database/model"

type BusinessCardPartsCoordinateResponse struct {
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
