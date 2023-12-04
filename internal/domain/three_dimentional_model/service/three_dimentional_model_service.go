package service

import "github.com/Misoten-B/airship-backend/internal/id"

type ThreeDimentionalModelService struct {
	repository ThreeDimentionalModelRepository
}

func NewThreeDimentionalModelService(repository ThreeDimentionalModelRepository) *ThreeDimentionalModelService {
	return &ThreeDimentionalModelService{
		repository: repository,
	}
}

// HasUsePermission は3Dモデルの使用権限を持っているかどうかを返します。
// 権限を持っている場合はtrueを、持っていない場合はfalseを返します。
func (s *ThreeDimentionalModelService) HasUsePermission(threeDimentionalModelID id.ID, userID id.ID) (bool, error) {
	threeDimentionalModel, err := s.repository.Find(threeDimentionalModelID)
	if err != nil {
		return false, err
	}

	return threeDimentionalModel.IsTemplate() || threeDimentionalModel.UserID() == userID, nil
}
