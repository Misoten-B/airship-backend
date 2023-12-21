package dto

import "github.com/Misoten-B/airship-backend/internal/drivers/database/model"

type BusinessCardPartsCoordinateResponse struct {
	ID                string `json:"id" binding:"required"`
	DisplayNameX      int    `json:"displayNameX" binding:"required"`
	DisplayNameY      int    `json:"displayNameY" binding:"required"`
	CompanyNameX      int    `json:"companyNameX" binding:"required"`
	CompanyNameY      int    `json:"companyNameY" binding:"required"`
	DepartmentX       int    `json:"departmentX" binding:"required"`
	DepartmentY       int    `json:"departmentY" binding:"required"`
	OfficialPositionX int    `json:"officialPositionX" binding:"required"`
	OfficialPositionY int    `json:"officialPositionY" binding:"required"`
	PhoneNumberX      int    `json:"phoneNumberX" binding:"required"`
	PhoneNumberY      int    `json:"phoneNumberY" binding:"required"`
	EmailX            int    `json:"emailX" binding:"required"`
	EmailY            int    `json:"emailY" binding:"required"`
	PostalCodeX       int    `json:"postalCodeX" binding:"required"`
	PostalCodeY       int    `json:"postalCodeY" binding:"required"`
	AddressX          int    `json:"addressX" binding:"required"`
	AddressY          int    `json:"addressY" binding:"required"`
	QrcodeX           int    `json:"qrcodeX" binding:"required"`
	QrcodeY           int    `json:"qrcodeY" binding:"required"`
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
