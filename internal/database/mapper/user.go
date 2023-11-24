package mapper

import (
	"github.com/Misoten-B/airship-backend/internal/database/model"
	domain "github.com/Misoten-B/airship-backend/internal/domain/user"
)

// ToUserModelはUser集約をORMモデルに変換します。
func ToUserORMModel(domain *domain.User) *model.User {
	return &model.User{
		ID:        domain.ID().String(),
		IsToured:  domain.IsToured(),
		CreatedAt: domain.CreatedAt(),
		DeletedAt: domain.DeletedAt(),
	}
}
