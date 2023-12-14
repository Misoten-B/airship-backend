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
	FindByUserID(userID id.ID) ([]threedimentionalmodel.ReadModel, error)
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
	log.Printf("Mock ThreeDimentionalModel Repository - Find: %s", id)
	return threedimentionalmodel.ReconstructThreeDimentionalModelTemplate(id), nil
}

func (r *MockThreeDimentionalModelRepository) FindByID(id id.ID) (threedimentionalmodel.ReadModel, error) {
	log.Printf("Mock ThreeDimentionalModel Repository - FindByID: %s", id)

	return threedimentionalmodel.NewReadModel(
		id.String(),
		testdata.DEV_UID,
		"mock-3d-model.glb",
	), nil
}

func (r *MockThreeDimentionalModelRepository) FindByUserID(userID id.ID) ([]threedimentionalmodel.ReadModel, error) {
	log.Printf("Mock ThreeDimentionalModel Repository - FindByUserID: %s", userID)

	return []threedimentionalmodel.ReadModel{
		threedimentionalmodel.NewReadModel(
			"1",
			testdata.DEV_UID,
			"mock-3d-model.glb",
		),
		threedimentionalmodel.NewTemplateReadModel(
			"2",
			"mock-3d-model-template.glb",
		),
	}, nil
}
