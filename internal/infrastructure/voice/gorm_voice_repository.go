package voice

import (
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/drivers/database/model"
	"gorm.io/gorm"
)

type GormVoiceRepository struct {
	db *gorm.DB
}

func NewGormVoiceRepository(db *gorm.DB) *GormVoiceRepository {
	return &GormVoiceRepository{
		db: db,
	}
}

func (r *GormVoiceRepository) FetchStatus(userID shared.ID) (shared.Status, error) {
	var user model.User

	err := r.db.Select("status").First(&user, "id = ?", userID.String()).Error
	if err != nil {
		return shared.StatusInProgress{}, err
	}

	if user.Status == model.GormStatusCompleted {
		return shared.StatusCompleted{}, nil
	}
	return shared.StatusInProgress{}, nil
}
