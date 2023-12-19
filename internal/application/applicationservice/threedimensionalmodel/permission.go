package threedimensionalmodel

import (
	tdmdomain "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
	"github.com/Misoten-B/airship-backend/internal/id"
)

func HasUpdatePermission(threeDimentionalModel tdmdomain.ThreeDimensionalModel, userID id.ID) bool {
	if threeDimentionalModel.IsTemplate() {
		return false
	}

	return threeDimentionalModel.UserID() == userID
}
