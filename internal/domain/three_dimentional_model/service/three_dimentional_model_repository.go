package service

import (
	"log"

	threedimentionalmodel "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
	"github.com/Misoten-B/airship-backend/internal/id"
	"github.com/Misoten-B/airship-backend/internal/testdata"
)

type ThreeDimentionalModelRepository interface {
	Find(id id.ID) (*threedimentionalmodel.ThreeDimentionalModel, error)
	FindByID(id id.ID) (threedimentionalmodel.ReadModel, error)
	Save(threeDimentionalModel threedimentionalmodel.ThreeDimentionalModel) error
}

type MockThreeDimentionalModelRepository struct{}

func NewMockThreeDimentionalModelRepository() *MockThreeDimentionalModelRepository {
	return &MockThreeDimentionalModelRepository{}
}

func (r *MockThreeDimentionalModelRepository) Save(_ threedimentionalmodel.ThreeDimentionalModel) error {
	log.Println("Mock ThreeDimentionalModel Repository - Save")
	return nil
}

func (r *MockThreeDimentionalModelRepository) Find(id id.ID) (*threedimentionalmodel.ThreeDimentionalModel, error) {
	log.Println("Mock ThreeDimentionalModel Repository - Find")
	return threedimentionalmodel.ReconstructThreeDimentionalModelTemplate(id), nil
}

func (r *MockThreeDimentionalModelRepository) FindByID(id id.ID) (threedimentionalmodel.ReadModel, error) {
	log.Println("Mock ThreeDimentionalModel Repository - FindByID")

	return threedimentionalmodel.NewReadModel(
		id.String(),
		testdata.DEV_UID,
		"mock-3d-model.glb",
	), nil
}
