package businesscard

import "github.com/Misoten-B/airship-backend/internal/domain/shared"

type BusinessCard struct {
	ID                          shared.ID
	UserID                      shared.ID
	BusinessCardPartsCoordinate shared.ID
	BusinessCardBackground      shared.ID
	ARAssets                    shared.ID
	Name                        *string
	DisplayName                 *string
	CompanyName                 *string
	Department                  *string
	OfficialPosition            *string
	PhoneNumber                 *string
	Email                       *string
	PostalCode                  *string
	Address                     *string
	AccessCount                 int
}

func NewBusinessCard(
	userID shared.ID,
	businessCardPartsCoordinate shared.ID,
	businessCardBackground shared.ID,
	arAssets shared.ID,
	name *string,
	displayName *string,
	companyName *string,
	department *string,
	officialPosition *string,
	phoneNumber *string,
	email *string,
	postalCode *string,
	address *string,
	accessCount int,
) (*BusinessCard, error) {
	id, err := shared.NewID()
	if err != nil {
		return nil, err
	}

	return &BusinessCard{
		ID:                          id,
		UserID:                      userID,
		BusinessCardPartsCoordinate: businessCardPartsCoordinate,
		BusinessCardBackground:      businessCardBackground,
		ARAssets:                    arAssets,
		Name:                        name,
		DisplayName:                 displayName,
		CompanyName:                 companyName,
		Department:                  department,
		OfficialPosition:            officialPosition,
		PhoneNumber:                 phoneNumber,
		Email:                       email,
		PostalCode:                  postalCode,
		Address:                     address,
		AccessCount:                 accessCount,
	}, nil
}
