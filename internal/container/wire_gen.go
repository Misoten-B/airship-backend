// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/create"
	"github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_id"
	"github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_id_public"
	"github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_userid"
	"github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/create"
	fetchbyid2 "github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/fetch_by_id"
	fetchbyuserid2 "github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/fetch_by_userid"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	service3 "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	service2 "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	"github.com/Misoten-B/airship-backend/internal/drivers/config"
	"github.com/Misoten-B/airship-backend/internal/infrastructure/ar_assets"
	"github.com/Misoten-B/airship-backend/internal/infrastructure/three_dimentional_model"
	"github.com/Misoten-B/airship-backend/internal/infrastructure/voice"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeCreateARAssetsUsecaseForDev() *usecase.ARAssetsUsecaseImpl {
	mockARAssetsRepository := service.NewMockARAssetsRepository()
	mockQRCodeImageStorage := service.NewMockQRCodeImageStorage()
	mockVoiceModelAdapter := service2.NewMockVoiceModelAdapter()
	mockVoiceRepository := service2.NewMockVoiceRepository()
	voiceServiceImpl := service2.NewVoiceServiceImpl(mockVoiceRepository)
	mockThreeDimentionalModelRepository := service3.NewMockThreeDimentionalModelRepository()
	threeDimentionalModelServiceImpl := service3.NewThreeDimentionalModelServiceImpl(mockThreeDimentionalModelRepository)
	arAssetsUsecaseImpl := usecase.NewARAssetsUsecaseImpl(mockARAssetsRepository, mockQRCodeImageStorage, mockVoiceModelAdapter, voiceServiceImpl, threeDimentionalModelServiceImpl)
	return arAssetsUsecaseImpl
}

func InitializeCreateARAssetsUsecaseForProd(db *gorm.DB, config2 *config.Config) *usecase.ARAssetsUsecaseImpl {
	gormARAssetsRepository := arassets.NewGormARAssetsRepository(db)
	azureBlobDriver := drivers.NewAzureBlobDriver(config2)
	azureQRCodeImageStorage := arassets.NewAzureQRCodeImageStorage(azureBlobDriver)
	externalAPIVoiceModelAdapter := voice.NewExternalAPIVoiceModelAdapter()
	gormVoiceRepository := voice.NewGormVoiceRepository(db)
	voiceServiceImpl := service2.NewVoiceServiceImpl(gormVoiceRepository)
	gormThreeDimentionalModelRepository := threedimentionalmodel.NewGormThreeDimentionalModelRepository(db)
	threeDimentionalModelServiceImpl := service3.NewThreeDimentionalModelServiceImpl(gormThreeDimentionalModelRepository)
	arAssetsUsecaseImpl := usecase.NewARAssetsUsecaseImpl(gormARAssetsRepository, azureQRCodeImageStorage, externalAPIVoiceModelAdapter, voiceServiceImpl, threeDimentionalModelServiceImpl)
	return arAssetsUsecaseImpl
}

func InitializeFetchByIDARAssetsUsecaseForDev() *fetchbyid.Interactor {
	mockARAssetsRepository := service.NewMockARAssetsRepository()
	mockQRCodeImageStorage := service.NewMockQRCodeImageStorage()
	mockSpeakingAudioStorage := service2.NewMockSpeakingAudioStorage()
	mockThreeDimentionalModelStorage := service3.NewMockThreeDimentionalModelStorage()
	interactor := fetchbyid.NewInteractor(mockARAssetsRepository, mockQRCodeImageStorage, mockSpeakingAudioStorage, mockThreeDimentionalModelStorage)
	return interactor
}

func InitializeFetchByIDARAssetsUsecaseForProd(db *gorm.DB, config2 *config.Config) *fetchbyid.Interactor {
	gormARAssetsRepository := arassets.NewGormARAssetsRepository(db)
	azureBlobDriver := drivers.NewAzureBlobDriver(config2)
	azureQRCodeImageStorage := arassets.NewAzureQRCodeImageStorage(azureBlobDriver)
	azureSpeakingAudioStorage := voice.NewAzureSpeakingAudioStorage(azureBlobDriver)
	azureThreeDimentionalModelStorage := threedimentionalmodel.NewAzureThreeDimentionalModelStorage(azureBlobDriver)
	interactor := fetchbyid.NewInteractor(gormARAssetsRepository, azureQRCodeImageStorage, azureSpeakingAudioStorage, azureThreeDimentionalModelStorage)
	return interactor
}

func InitializeFetchByIDPublicARAssetsUsecaseForDev() *fetchbyidpublic.Interactor {
	mockARAssetsRepository := service.NewMockARAssetsRepository()
	mockSpeakingAudioStorage := service2.NewMockSpeakingAudioStorage()
	mockThreeDimentionalModelStorage := service3.NewMockThreeDimentionalModelStorage()
	interactor := fetchbyidpublic.NewInteractor(mockARAssetsRepository, mockSpeakingAudioStorage, mockThreeDimentionalModelStorage)
	return interactor
}

func InitializeFetchByIDPublicARAssetsUsecaseForProd(db *gorm.DB, config2 *config.Config) *fetchbyidpublic.Interactor {
	gormARAssetsRepository := arassets.NewGormARAssetsRepository(db)
	azureBlobDriver := drivers.NewAzureBlobDriver(config2)
	azureSpeakingAudioStorage := voice.NewAzureSpeakingAudioStorage(azureBlobDriver)
	azureThreeDimentionalModelStorage := threedimentionalmodel.NewAzureThreeDimentionalModelStorage(azureBlobDriver)
	interactor := fetchbyidpublic.NewInteractor(gormARAssetsRepository, azureSpeakingAudioStorage, azureThreeDimentionalModelStorage)
	return interactor
}

func InitializeFetchByUserIDARAssetsUsecaseForDev() *fetchbyuserid.Interactor {
	mockARAssetsRepository := service.NewMockARAssetsRepository()
	mockQRCodeImageStorage := service.NewMockQRCodeImageStorage()
	mockSpeakingAudioStorage := service2.NewMockSpeakingAudioStorage()
	mockThreeDimentionalModelStorage := service3.NewMockThreeDimentionalModelStorage()
	interactor := fetchbyuserid.NewInteractor(mockARAssetsRepository, mockQRCodeImageStorage, mockSpeakingAudioStorage, mockThreeDimentionalModelStorage)
	return interactor
}

func InitializeFetchByUserIDARAssetsUsecaseForProd(db *gorm.DB, config2 *config.Config) *fetchbyuserid.Interactor {
	gormARAssetsRepository := arassets.NewGormARAssetsRepository(db)
	azureBlobDriver := drivers.NewAzureBlobDriver(config2)
	azureQRCodeImageStorage := arassets.NewAzureQRCodeImageStorage(azureBlobDriver)
	azureSpeakingAudioStorage := voice.NewAzureSpeakingAudioStorage(azureBlobDriver)
	azureThreeDimentionalModelStorage := threedimentionalmodel.NewAzureThreeDimentionalModelStorage(azureBlobDriver)
	interactor := fetchbyuserid.NewInteractor(gormARAssetsRepository, azureQRCodeImageStorage, azureSpeakingAudioStorage, azureThreeDimentionalModelStorage)
	return interactor
}

func InitializeCreateThreeDimentionalModelUsecaseForDev() *create.Interactor {
	mockThreeDimentionalModelStorage := service3.NewMockThreeDimentionalModelStorage()
	mockThreeDimentionalModelRepository := service3.NewMockThreeDimentionalModelRepository()
	interactor := create.NewInteractor(mockThreeDimentionalModelStorage, mockThreeDimentionalModelRepository)
	return interactor
}

func InitializeCreateThreeDimentionalModelUsecaseForProd(db *gorm.DB, config2 *config.Config) *create.Interactor {
	azureBlobDriver := drivers.NewAzureBlobDriver(config2)
	azureThreeDimentionalModelStorage := threedimentionalmodel.NewAzureThreeDimentionalModelStorage(azureBlobDriver)
	gormThreeDimentionalModelRepository := threedimentionalmodel.NewGormThreeDimentionalModelRepository(db)
	interactor := create.NewInteractor(azureThreeDimentionalModelStorage, gormThreeDimentionalModelRepository)
	return interactor
}

func InitializeFetchByIDThreeDimentionalModelUsecaseForDev() *fetchbyid2.Interactor {
	mockThreeDimentionalModelRepository := service3.NewMockThreeDimentionalModelRepository()
	mockThreeDimentionalModelStorage := service3.NewMockThreeDimentionalModelStorage()
	interactor := fetchbyid2.NewInteractor(mockThreeDimentionalModelRepository, mockThreeDimentionalModelStorage)
	return interactor
}

func InitializeFetchByIDThreeDimentionalModelUsecaseForProd(db *gorm.DB, config2 *config.Config) *fetchbyid2.Interactor {
	gormThreeDimentionalModelRepository := threedimentionalmodel.NewGormThreeDimentionalModelRepository(db)
	azureBlobDriver := drivers.NewAzureBlobDriver(config2)
	azureThreeDimentionalModelStorage := threedimentionalmodel.NewAzureThreeDimentionalModelStorage(azureBlobDriver)
	interactor := fetchbyid2.NewInteractor(gormThreeDimentionalModelRepository, azureThreeDimentionalModelStorage)
	return interactor
}

func InitializeFetchByUserIDThreeDimentionalModelUsecaseForDev() *fetchbyuserid2.Interactor {
	mockThreeDimentionalModelRepository := service3.NewMockThreeDimentionalModelRepository()
	mockThreeDimentionalModelStorage := service3.NewMockThreeDimentionalModelStorage()
	interactor := fetchbyuserid2.NewInteractor(mockThreeDimentionalModelRepository, mockThreeDimentionalModelStorage)
	return interactor
}

func InitializeFetchByUserIDThreeDimentionalModelUsecaseForProd(db *gorm.DB, config2 *config.Config) *fetchbyuserid2.Interactor {
	gormThreeDimentionalModelRepository := threedimentionalmodel.NewGormThreeDimentionalModelRepository(db)
	azureBlobDriver := drivers.NewAzureBlobDriver(config2)
	azureThreeDimentionalModelStorage := threedimentionalmodel.NewAzureThreeDimentionalModelStorage(azureBlobDriver)
	interactor := fetchbyuserid2.NewInteractor(gormThreeDimentionalModelRepository, azureThreeDimentionalModelStorage)
	return interactor
}

// wire.go:

// CreateARAssetsUsecaseSetForDev は開発環境用のプロバイダセットです。
// 現状、このファイルに記述されていますが、将来的にはユースケースのファクトリ部分に移動することも
// 検討しています。
var CreateARAssetsUsecaseSetForDev = wire.NewSet(usecase.NewARAssetsUsecaseImpl, MockARAssetsRepositorySet,
	MockQRCodeImageStorageSet,
	MockVoiceModelAdapterSet,
	VoiceServiceImplSet,
	MockVoiceRepositorySet,
	ThreeDimentionalModelServiceImplSet,
	MockThreeDimentionalModelRepositorySet,
)

var CreateARAssetsUsecaseSetForProd = wire.NewSet(usecase.NewARAssetsUsecaseImpl, drivers.NewAzureBlobDriver, GormARAssetsRepositorySet,
	AzureQRCodeImageStorageSet,
	ExternalAPIVoiceModelAdapterSet,
	VoiceServiceImplSet,
	GormVoiceRepositorySet,
	ThreeDimentionalModelServiceImplSet,
	GormThreeDimentionalModelRepositorySet,
)

var FetchByIDARAssetsUsecaseSetForDev = wire.NewSet(fetchbyid.NewInteractor, MockARAssetsRepositorySet,
	MockQRCodeImageStorageSet,
	MockSpeakingAudioStorageSet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByIDARAssetsUsecaseSetForProd = wire.NewSet(fetchbyid.NewInteractor, drivers.NewAzureBlobDriver, GormARAssetsRepositorySet,
	AzureQRCodeImageStorageSet,
	AzureSpeakingAudioStorageSet,
	AzureThreeDimentionalModelStorageSet,
)

var FetchByIDPublicARAssetsUsecaseSetForDev = wire.NewSet(fetchbyidpublic.NewInteractor, MockARAssetsRepositorySet,
	MockSpeakingAudioStorageSet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByIDPublicARAssetsUsecaseSetForProd = wire.NewSet(fetchbyidpublic.NewInteractor, drivers.NewAzureBlobDriver, GormARAssetsRepositorySet,
	AzureSpeakingAudioStorageSet,
	AzureThreeDimentionalModelStorageSet,
)

var FetchByUserIDARAssetsUsecaseSetForDev = wire.NewSet(fetchbyuserid.NewInteractor, MockARAssetsRepositorySet,
	MockQRCodeImageStorageSet,
	MockSpeakingAudioStorageSet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByUserIDARAssetsUsecaseSetForProd = wire.NewSet(fetchbyuserid.NewInteractor, drivers.NewAzureBlobDriver, GormARAssetsRepositorySet,
	AzureQRCodeImageStorageSet,
	AzureSpeakingAudioStorageSet,
	AzureThreeDimentionalModelStorageSet,
)

var CreateThreeDimentionalModelUsecaseSetForDev = wire.NewSet(create.NewInteractor, MockThreeDimentionalModelRepositorySet,
	MockThreeDimentionalModelStorageSet,
)

var CreateThreeDimentionalModelUsecaseSetForProd = wire.NewSet(create.NewInteractor, drivers.NewAzureBlobDriver, GormThreeDimentionalModelRepositorySet,
	AzureThreeDimentionalModelStorageSet,
)

var FetchByIDThreeDimentionalModelUsecaseSetForDev = wire.NewSet(fetchbyid2.NewInteractor, MockThreeDimentionalModelRepositorySet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByIDThreeDimentionalModelUsecaseSetForProd = wire.NewSet(fetchbyid2.NewInteractor, drivers.NewAzureBlobDriver, GormThreeDimentionalModelRepositorySet,
	AzureThreeDimentionalModelStorageSet,
)

var FetchByUserIDThreeDimentionalModelUsecaseSetForDev = wire.NewSet(fetchbyuserid2.NewInteractor, MockThreeDimentionalModelRepositorySet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByUserIDThreeDimentionalModelUsecaseSetForProd = wire.NewSet(fetchbyuserid2.NewInteractor, drivers.NewAzureBlobDriver, GormThreeDimentionalModelRepositorySet,
	AzureThreeDimentionalModelStorageSet,
)

var MockARAssetsRepositorySet = wire.NewSet(service.NewMockARAssetsRepository, wire.Bind(new(service.ARAssetsRepository), new(*service.MockARAssetsRepository)))

var GormARAssetsRepositorySet = wire.NewSet(arassets.NewGormARAssetsRepository, wire.Bind(new(service.ARAssetsRepository), new(*arassets.GormARAssetsRepository)))

var MockQRCodeImageStorageSet = wire.NewSet(service.NewMockQRCodeImageStorage, wire.Bind(new(service.QRCodeImageStorage), new(*service.MockQRCodeImageStorage)))

var AzureQRCodeImageStorageSet = wire.NewSet(arassets.NewAzureQRCodeImageStorage, wire.Bind(new(service.QRCodeImageStorage), new(*arassets.AzureQRCodeImageStorage)))

var MockVoiceModelAdapterSet = wire.NewSet(service2.NewMockVoiceModelAdapter, wire.Bind(new(service2.VoiceModelAdapter), new(*service2.MockVoiceModelAdapter)))

var ExternalAPIVoiceModelAdapterSet = wire.NewSet(voice.NewExternalAPIVoiceModelAdapter, wire.Bind(new(service2.VoiceModelAdapter), new(*voice.ExternalAPIVoiceModelAdapter)))

var VoiceServiceImplSet = wire.NewSet(service2.NewVoiceServiceImpl, wire.Bind(new(service2.VoiceService), new(*service2.VoiceServiceImpl)))

var MockVoiceRepositorySet = wire.NewSet(service2.NewMockVoiceRepository, wire.Bind(new(service2.VoiceRepository), new(*service2.MockVoiceRepository)))

var GormVoiceRepositorySet = wire.NewSet(voice.NewGormVoiceRepository, wire.Bind(new(service2.VoiceRepository), new(*voice.GormVoiceRepository)))

var MockSpeakingAudioStorageSet = wire.NewSet(service2.NewMockSpeakingAudioStorage, wire.Bind(new(service2.SpeakingAudioStorage), new(*service2.MockSpeakingAudioStorage)))

var AzureSpeakingAudioStorageSet = wire.NewSet(voice.NewAzureSpeakingAudioStorage, wire.Bind(new(service2.SpeakingAudioStorage), new(*voice.AzureSpeakingAudioStorage)))

var ThreeDimentionalModelServiceImplSet = wire.NewSet(service3.NewThreeDimentionalModelServiceImpl, wire.Bind(new(service3.ThreeDimentionalModelService), new(*service3.ThreeDimentionalModelServiceImpl)))

var MockThreeDimentionalModelRepositorySet = wire.NewSet(service3.NewMockThreeDimentionalModelRepository, wire.Bind(new(service3.ThreeDimentionalModelRepository), new(*service3.MockThreeDimentionalModelRepository)))

var GormThreeDimentionalModelRepositorySet = wire.NewSet(threedimentionalmodel.NewGormThreeDimentionalModelRepository, wire.Bind(new(service3.ThreeDimentionalModelRepository), new(*threedimentionalmodel.GormThreeDimentionalModelRepository)))

var MockThreeDimentionalModelStorageSet = wire.NewSet(service3.NewMockThreeDimentionalModelStorage, wire.Bind(new(service3.ThreeDimentionalModelStorage), new(*service3.MockThreeDimentionalModelStorage)))

var AzureThreeDimentionalModelStorageSet = wire.NewSet(threedimentionalmodel.NewAzureThreeDimentionalModelStorage, wire.Bind(new(service3.ThreeDimentionalModelStorage), new(*threedimentionalmodel.AzureThreeDimentionalModelStorage)))
