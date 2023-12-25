package service

import (
	"log"

	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/testdata"
)

type ARAssetsRepository interface {
	Save(arassets arassets.ARAssets) error
	FetchByID(id shared.ID) (arassets.ReadModel, error)
	FetchByUserID(userID shared.ID) ([]arassets.ReadModel, error)
	PatchStatus(id shared.ID, status shared.Status) error
}

type MockARAssetsRepository struct{}

func NewMockARAssetsRepository() *MockARAssetsRepository {
	return &MockARAssetsRepository{}
}

func (r *MockARAssetsRepository) Save(_ arassets.ARAssets) error {
	log.Println("Mock ARAssets Repository - Save")
	return nil
}

func (r *MockARAssetsRepository) FetchByID(id shared.ID) (arassets.ReadModel, error) {
	log.Printf("Mock ARAssets Repository - FetchByID: %s", id)

	return arassets.NewReadModel(
		id.String(),
		testdata.DEV_UID,
		"test description",
		"mock-example.mp3",
		"mock-example.glb",
		"mock-example.png",
		true,
	), nil
}

func (r *MockARAssetsRepository) FetchByUserID(userID shared.ID) ([]arassets.ReadModel, error) {
	log.Printf("Mock ARAssets Repository - FetchByUserID: %s", userID)

	return []arassets.ReadModel{
		arassets.NewReadModel(
			"1",
			testdata.DEV_UID,
			"test description",
			"mock-example.mp3",
			"mock-example.glb",
			"mock-example.png",
			true,
		),
	}, nil
}

func (r *MockARAssetsRepository) PatchStatus(id shared.ID, status shared.Status) error {
	log.Printf("Mock ARAssets Repository - PatchStatus: id: %s, status: %d", id, status.Status())

	return nil
}
