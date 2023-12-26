package service

import (
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
)

type VoiceService interface {
	IsModelGenerated(userID shared.ID) (bool, error)
}

type VoiceServiceImpl struct {
	repository VoiceRepository
}

func NewVoiceServiceImpl(repository VoiceRepository) *VoiceServiceImpl {
	return &VoiceServiceImpl{
		repository: repository,
	}
}

// IsModelGenerated は音声モデルが生成済みかどうかを判定します。
// 既に生成が完了している場合はtrueを返します。
func (s *VoiceServiceImpl) IsModelGenerated(userID shared.ID) (bool, error) {
	status, err := s.repository.FetchStatus(userID)
	if err != nil {
		return false, err
	}

	completed := shared.StatusCompleted{}
	if completed.Equal(status) {
		return true, nil
	}
	return false, nil
}
