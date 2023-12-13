package threedimentionalmodel_test

import (
	"testing"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	threedimentionalmodel "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
	"github.com/Misoten-B/airship-backend/internal/id"
)

func TestThreeDimensionalModel(t *testing.T) {
	t.Parallel()

	tdmID := id.ReconstructID("test-id")
	userID := id.ReconstructID("test-user-id")
	filePath := shared.NewFilePath("test.gltf")

	newPersonal, err := threedimentionalmodel.NewThreeDimensionalModel(userID, filePath)
	if err != nil {
		t.Fatal(err)
	}

	if newPersonal.IsTemplate() {
		t.Fatal("newPersonal should not be template")
	}

	personal := threedimentionalmodel.ReconstructThreeDimensionalModel(tdmID, userID, filePath)

	if personal.IsTemplate() {
		t.Fatal("personal should not be template")
	}

	template := threedimentionalmodel.ReconstructThreeDimensionalModelTemplate(tdmID, filePath)

	if !template.IsTemplate() {
		t.Fatal("template should be template")
	}
}
