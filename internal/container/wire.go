//go:build wireinject
// +build wireinject

package container

import (
	"github.com/Misoten-B/airship-backend/config"
	arcreateusecase "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/create"
	arfetchbyidusecase "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_id"
	arfetchbyidpubusecase "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_id_public"
	arfetchbyuseridusecase "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/fetch_by_userid"
	createtdmusecase "github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/create"
	tdmfetchbyidusecase "github.com/Misoten-B/airship-backend/internal/application/usecase/three_dimentional_model/fetch_by_id"
	ardomain "github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	tdmdomain "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	vdomain "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/drivers"
	arinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/ar_assets"
	tdminfra "github.com/Misoten-B/airship-backend/internal/infrastructure/three_dimentional_model"
	vinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/voice"
	"github.com/google/wire"
	"gorm.io/gorm"
)

/** Usecase Provider Set **/

/** ARAssets **/

// CreateARAssetsUsecaseSetForDev は開発環境用のプロバイダセットです。
// 現状、このファイルに記述されていますが、将来的にはユースケースのファクトリ部分に移動することも
// 検討しています。
var CreateARAssetsUsecaseSetForDev = wire.NewSet(
	arcreateusecase.NewARAssetsUsecaseImpl,
	MockARAssetsRepositorySet,
	MockQRCodeImageStorageSet,
	MockVoiceModelAdapterSet,
	VoiceServiceImplSet,
	MockVoiceRepositorySet,
	ThreeDimentionalModelServiceImplSet,
	MockThreeDimentionalModelRepositorySet,
)

var CreateARAssetsUsecaseSetForProd = wire.NewSet(
	arcreateusecase.NewARAssetsUsecaseImpl,
	drivers.NewAzureBlobDriver,
	GormARAssetsRepositorySet,
	AzureQRCodeImageStorageSet,
	ExternalAPIVoiceModelAdapterSet,
	VoiceServiceImplSet,
	GormVoiceRepositorySet,
	ThreeDimentionalModelServiceImplSet,
	GormThreeDimentionalModelRepositorySet,
)

var FetchByIDARAssetsUsecaseSetForDev = wire.NewSet(
	arfetchbyidusecase.NewInteractor,
	MockARAssetsRepositorySet,
	MockQRCodeImageStorageSet,
	MockSpeakingAudioStorageSet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByIDARAssetsUsecaseSetForProd = wire.NewSet(
	arfetchbyidusecase.NewInteractor,
	drivers.NewAzureBlobDriver,
	GormARAssetsRepositorySet,
	AzureQRCodeImageStorageSet,
	AzureSpeakingAudioStorageSet,
	AzureThreeDimentionalModelStorageSet,
)

var FetchByIDPublicARAssetsUsecaseSetForDev = wire.NewSet(
	arfetchbyidpubusecase.NewInteractor,
	MockARAssetsRepositorySet,
	MockSpeakingAudioStorageSet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByIDPublicARAssetsUsecaseSetForProd = wire.NewSet(
	arfetchbyidpubusecase.NewInteractor,
	drivers.NewAzureBlobDriver,
	GormARAssetsRepositorySet,
	AzureSpeakingAudioStorageSet,
	AzureThreeDimentionalModelStorageSet,
)

var FetchByUserIDARAssetsUsecaseSetForDev = wire.NewSet(
	arfetchbyuseridusecase.NewInteractor,
	MockARAssetsRepositorySet,
	MockQRCodeImageStorageSet,
	MockSpeakingAudioStorageSet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByUserIDARAssetsUsecaseSetForProd = wire.NewSet(
	arfetchbyuseridusecase.NewInteractor,
	drivers.NewAzureBlobDriver,
	GormARAssetsRepositorySet,
	AzureQRCodeImageStorageSet,
	AzureSpeakingAudioStorageSet,
	AzureThreeDimentionalModelStorageSet,
)

/** Three Dimentional Model **/

var CreateThreeDimentionalModelUsecaseSetForDev = wire.NewSet(
	createtdmusecase.NewInteractor,
	MockThreeDimentionalModelRepositorySet,
	MockThreeDimentionalModelStorageSet,
)

var CreateThreeDimentionalModelUsecaseSetForProd = wire.NewSet(
	createtdmusecase.NewInteractor,
	drivers.NewAzureBlobDriver,
	GormThreeDimentionalModelRepositorySet,
	AzureThreeDimentionalModelStorageSet,
)

var FetchByIDThreeDimentionalModelUsecaseSetForDev = wire.NewSet(
	tdmfetchbyidusecase.NewInteractor,
	MockThreeDimentionalModelRepositorySet,
	MockThreeDimentionalModelStorageSet,
)

var FetchByIDThreeDimentionalModelUsecaseSetForProd = wire.NewSet(
	tdmfetchbyidusecase.NewInteractor,
	drivers.NewAzureBlobDriver,
	GormThreeDimentionalModelRepositorySet,
	AzureThreeDimentionalModelStorageSet,
)

/* Interface Binding */

/** ARAssets **/

var MockARAssetsRepositorySet = wire.NewSet(
	ardomain.NewMockARAssetsRepository,
	wire.Bind(new(ardomain.ARAssetsRepository), new(*ardomain.MockARAssetsRepository)),
)

var GormARAssetsRepositorySet = wire.NewSet(
	arinfra.NewGormARAssetsRepository,
	wire.Bind(new(ardomain.ARAssetsRepository), new(*arinfra.GormARAssetsRepository)),
)

var MockQRCodeImageStorageSet = wire.NewSet(
	ardomain.NewMockQRCodeImageStorage,
	wire.Bind(new(ardomain.QRCodeImageStorage), new(*ardomain.MockQRCodeImageStorage)),
)

var AzureQRCodeImageStorageSet = wire.NewSet(
	arinfra.NewAzureQRCodeImageStorage,
	wire.Bind(new(ardomain.QRCodeImageStorage), new(*arinfra.AzureQRCodeImageStorage)),
)

/** Voice **/

var MockVoiceModelAdapterSet = wire.NewSet(
	vdomain.NewMockVoiceModelAdapter,
	wire.Bind(new(vdomain.VoiceModelAdapter), new(*vdomain.MockVoiceModelAdapter)),
)

var ExternalAPIVoiceModelAdapterSet = wire.NewSet(
	vinfra.NewExternalAPIVoiceModelAdapter,
	wire.Bind(new(vdomain.VoiceModelAdapter), new(*vinfra.ExternalAPIVoiceModelAdapter)),
)

var VoiceServiceImplSet = wire.NewSet(
	vdomain.NewVoiceServiceImpl,
	wire.Bind(new(vdomain.VoiceService), new(*vdomain.VoiceServiceImpl)),
)

var MockVoiceRepositorySet = wire.NewSet(
	vdomain.NewMockVoiceRepository,
	wire.Bind(new(vdomain.VoiceRepository), new(*vdomain.MockVoiceRepository)),
)

var GormVoiceRepositorySet = wire.NewSet(
	vinfra.NewGormVoiceRepository,
	wire.Bind(new(vdomain.VoiceRepository), new(*vinfra.GormVoiceRepository)),
)

var MockSpeakingAudioStorageSet = wire.NewSet(
	vdomain.NewMockSpeakingAudioStorage,
	wire.Bind(new(vdomain.SpeakingAudioStorage), new(*vdomain.MockSpeakingAudioStorage)),
)

var AzureSpeakingAudioStorageSet = wire.NewSet(
	vinfra.NewAzureSpeakingAudioStorage,
	wire.Bind(new(vdomain.SpeakingAudioStorage), new(*vinfra.AzureSpeakingAudioStorage)),
)

/** ThreeDimentionalModel **/

var ThreeDimentionalModelServiceImplSet = wire.NewSet(
	tdmdomain.NewThreeDimentionalModelServiceImpl,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelService), new(*tdmdomain.ThreeDimentionalModelServiceImpl)),
)

var MockThreeDimentionalModelRepositorySet = wire.NewSet(
	tdmdomain.NewMockThreeDimentionalModelRepository,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelRepository), new(*tdmdomain.MockThreeDimentionalModelRepository)),
)

var GormThreeDimentionalModelRepositorySet = wire.NewSet(
	tdminfra.NewGormThreeDimentionalModelRepository,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelRepository), new(*tdminfra.GormThreeDimentionalModelRepository)),
)

var MockThreeDimentionalModelStorageSet = wire.NewSet(
	tdmdomain.NewMockThreeDimentionalModelStorage,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelStorage), new(*tdmdomain.MockThreeDimentionalModelStorage)),
)

var AzureThreeDimentionalModelStorageSet = wire.NewSet(
	tdminfra.NewAzureThreeDimentionalModelStorage,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelStorage), new(*tdminfra.AzureThreeDimentionalModelStorage)),
)

/** Injectors **/

func InitializeCreateARAssetsUsecaseForDev() *arcreateusecase.ARAssetsUsecaseImpl {
	wire.Build(CreateARAssetsUsecaseSetForDev)

	return &arcreateusecase.ARAssetsUsecaseImpl{}
}

func InitializeCreateARAssetsUsecaseForProd(db *gorm.DB, config *config.Config) *arcreateusecase.ARAssetsUsecaseImpl {
	wire.Build(CreateARAssetsUsecaseSetForProd)

	return &arcreateusecase.ARAssetsUsecaseImpl{}
}

func InitializeFetchByIDARAssetsUsecaseForDev() *arfetchbyidusecase.Interactor {
	wire.Build(FetchByIDARAssetsUsecaseSetForDev)

	return &arfetchbyidusecase.Interactor{}
}

func InitializeFetchByIDARAssetsUsecaseForProd(db *gorm.DB, config *config.Config) *arfetchbyidusecase.Interactor {
	wire.Build(FetchByIDARAssetsUsecaseSetForProd)

	return &arfetchbyidusecase.Interactor{}
}

func InitializeFetchByIDPublicARAssetsUsecaseForDev() *arfetchbyidpubusecase.Interactor {
	wire.Build(FetchByIDPublicARAssetsUsecaseSetForDev)

	return &arfetchbyidpubusecase.Interactor{}
}

func InitializeFetchByIDPublicARAssetsUsecaseForProd(
	db *gorm.DB, config *config.Config,
) *arfetchbyidpubusecase.Interactor {
	wire.Build(FetchByIDPublicARAssetsUsecaseSetForProd)

	return &arfetchbyidpubusecase.Interactor{}
}

func InitializeFetchByUserIDARAssetsUsecaseForDev() *arfetchbyuseridusecase.Interactor {
	wire.Build(FetchByUserIDARAssetsUsecaseSetForDev)

	return &arfetchbyuseridusecase.Interactor{}
}

func InitializeFetchByUserIDARAssetsUsecaseForProd(
	db *gorm.DB, config *config.Config,
) *arfetchbyuseridusecase.Interactor {
	wire.Build(FetchByUserIDARAssetsUsecaseSetForProd)

	return &arfetchbyuseridusecase.Interactor{}
}

/** Three Dimentional Model **/

func InitializeCreateThreeDimentionalModelUsecaseForDev() *createtdmusecase.Interactor {
	wire.Build(CreateThreeDimentionalModelUsecaseSetForDev)

	return &createtdmusecase.Interactor{}
}

func InitializeCreateThreeDimentionalModelUsecaseForProd(
	db *gorm.DB, config *config.Config,
) *createtdmusecase.Interactor {
	wire.Build(CreateThreeDimentionalModelUsecaseSetForProd)

	return &createtdmusecase.Interactor{}
}

func InitializeFetchByIDThreeDimentionalModelUsecaseForDev() *tdmfetchbyidusecase.Interactor {
	wire.Build(FetchByIDThreeDimentionalModelUsecaseSetForDev)

	return &tdmfetchbyidusecase.Interactor{}
}

func InitializeFetchByIDThreeDimentionalModelUsecaseForProd(
	db *gorm.DB, config *config.Config,
) *tdmfetchbyidusecase.Interactor {
	wire.Build(FetchByIDThreeDimentionalModelUsecaseSetForProd)

	return &tdmfetchbyidusecase.Interactor{}
}
