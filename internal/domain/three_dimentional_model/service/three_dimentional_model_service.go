package service

import "github.com/Misoten-B/airship-backend/internal/id"

type ThreeDimentionalModelService interface {
	HasUsePermission(threeDimentionalModelID id.ID, userID id.ID) (bool, error)
}

type ThreeDimentionalModelServiceImpl struct {
	repository ThreeDimentionalModelRepository
}

func NewThreeDimentionalModelServiceImpl(repository ThreeDimentionalModelRepository) *ThreeDimentionalModelServiceImpl {
	return &ThreeDimentionalModelServiceImpl{
		repository: repository,
	}
}

// HasUsePermission は3Dモデルの使用権限を持っているかどうかを返します。
// 権限を持っている場合はtrueを、持っていない場合はfalseを返します。
func (s *ThreeDimentionalModelServiceImpl) HasUsePermission(threeDimentionalModelID id.ID, userID id.ID) (bool, error) {
	threeDimentionalModel, err := s.repository.Find(threeDimentionalModelID)
	if err != nil {
		return false, err
	}

	return threeDimentionalModel.IsTemplate() || threeDimentionalModel.UserID() == userID, nil
}
