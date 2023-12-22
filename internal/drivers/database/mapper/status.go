package mapper

import (
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
)

func ToGormStatus(status shared.Status) int {
	switch status.(type) {
	case shared.StatusError:
		return model.GormStatusError
	case shared.StatusNone:
		return model.GormStatusNone
	case shared.StatusInProgress:
		return model.GormStatusInProgress
	case shared.StatusCompleted:
		return model.GormStatusCompleted
	default:
		return model.GormStatusInProgress
	}
}
