package businesscard

import "github.com/Misoten-B/airship-backend/internal/id"

type BusinessCard struct {
	ID                          id.ID
	UserID                      id.ID
	BusinessCardPartsCoordinate id.ID
	BusinessCardBackground      id.ID
	ARAssets                    id.ID
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
	userID id.ID,
	businessCardPartsCoordinate id.ID,
	businessCardBackground id.ID,
	arAssets id.ID,
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
	id, err := id.NewID()
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
