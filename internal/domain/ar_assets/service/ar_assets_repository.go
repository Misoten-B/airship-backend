package service

import (
	"log"

	arassets "github.com/Misoten-B/airship-backend/internal/domain/ar_assets"
)

type ARAssetsRepository interface {
	Save(arassets arassets.ARAssets) error
}

type MockARAssetsRepository struct{}

func NewMockARAssetsRepository() *MockARAssetsRepository {
	return &MockARAssetsRepository{}
}

func (r *MockARAssetsRepository) Save(_ arassets.ARAssets) error {
	log.Println("Mock ARAssets Repository - Save")
	return nil
}
