package usecase

import (
	"github.com/Misoten-B/airship-backend/config"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	threeservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	voiceservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	arassets "github.com/Misoten-B/airship-backend/internal/infrastructure/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/infrastructure/voice"
	"gorm.io/gorm"
)

// ARAssetsConfiguration はARAssetsUsecaseImplのポインタを受け取り、それを変更する関数です。
type ARAssetsConfiguration func(aruc *ARAssetsUsecaseImpl) error

// NewARAssetsUsecase は可変長のARAssetsConfigurationを受け取り、ARAssetsUsecaseImplのポインタを返します。
func NewARAssetsUsecase(configs ...ARAssetsConfiguration) (*ARAssetsUsecaseImpl, error) {
	aruc := &ARAssetsUsecaseImpl{}

	for _, config := range configs {
		if err := config(aruc); err != nil {
			return nil, err
		}
	}

	return aruc, nil
}

func WithARAssetsRepository(arAssetsRepository service.ARAssetsRepository) ARAssetsConfiguration {
	return func(aruc *ARAssetsUsecaseImpl) error {
		aruc.arAssetsRepository = arAssetsRepository
		return nil
	}
}

func WithMockARAssetsRepository() ARAssetsConfiguration {
	ar := service.NewMockARAssetsRepository()
	return WithARAssetsRepository(ar)
}

func WithGormARAssetsRepository(db *gorm.DB) ARAssetsConfiguration {
	ar := arassets.NewGormARAssetsRepository(db)
	return WithARAssetsRepository(ar)
}

func WithQRCodeImageStorage(qrCodeImageStorage service.QRCodeImageStorage) ARAssetsConfiguration {
	return func(aruc *ARAssetsUsecaseImpl) error {
		aruc.qrCodeImageStorage = qrCodeImageStorage
		return nil
	}
}

func WithMockQRCodeImageStorage() ARAssetsConfiguration {
	qr := service.NewMockQRCodeImageStorage()
	return WithQRCodeImageStorage(qr)
}

func WithAzureQRCodeImageStorage(config *config.Config) ARAssetsConfiguration {
	qr := arassets.NewAzureQRCodeImageStorage(config)
	return WithQRCodeImageStorage(qr)
}

func WithVoiceModelAdapter(voiceModelAdapter voiceservice.VoiceModelAdapter) ARAssetsConfiguration {
	return func(aruc *ARAssetsUsecaseImpl) error {
		aruc.voiceModelAdapter = voiceModelAdapter
		return nil
	}
}

func WithMockVoiceModelAdapter() ARAssetsConfiguration {
	voice := voiceservice.NewMockVoiceModelAdapter()
	return WithVoiceModelAdapter(voice)
}

func WithExternalAPIVoiceModelAdapter() ARAssetsConfiguration {
	voice := voice.NewExternalAPIVoiceModelAdapter()
	return WithVoiceModelAdapter(voice)
}

func WithVoiceService(voiceService voiceservice.VoiceService) ARAssetsConfiguration {
	return func(aruc *ARAssetsUsecaseImpl) error {
		aruc.voiceService = voiceService
		return nil
	}
}

func WithVoiceServiceImpl(repository voiceservice.VoiceRepository) ARAssetsConfiguration {
	voice := voiceservice.NewVoiceServiceImpl(repository)
	return WithVoiceService(voice)
}

func WithThreeDimentionalModelService(
	threeDimentionalModelService threeservice.ThreeDimentionalModelService,
) ARAssetsConfiguration {
	return func(aruc *ARAssetsUsecaseImpl) error {
		aruc.threeDimentionalModelService = threeDimentionalModelService
		return nil
	}
}

func WithThreeDimentionalModelServiceImpl(
	repository threeservice.ThreeDimentionalModelRepository,
) ARAssetsConfiguration {
	three := threeservice.NewThreeDimentionalModelServiceImpl(repository)
	return WithThreeDimentionalModelService(three)
}
