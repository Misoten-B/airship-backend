package mapper

import (
	"github.com/Misoten-B/airship-backend/internal/database/model"
	domain "github.com/Misoten-B/airship-backend/internal/domain/user"
)

// ToUserDomainはORMモデルをUser集約に変換します。
func ToUserDomain(model *model.User) *domain.User {
	return &domain.User{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
	}
}

// ToUserModelはUser集約をORMモデルに変換します。
func ToUserORMModel(domain *domain.User) *model.User {
	return &model.User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
	}
}
