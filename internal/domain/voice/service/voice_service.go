package service

import (
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type VoiceService struct {
	repository VoiceRepository
}

func NewVoiceService(repository VoiceRepository) *VoiceService {
	return &VoiceService{
		repository: repository,
	}
}

// IsModelGenerated は音声モデルが生成済みかどうかを判定します。
// 既に生成が完了している場合はtrueを返します。
func (s *VoiceService) IsModelGenerated(userID id.ID) (bool, error) {
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
