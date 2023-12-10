//go:build wireinject
// +build wireinject

package container

import (
	"github.com/Misoten-B/airship-backend/config"
	arcreateusecase "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets/create"
	ardomain "github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	tdmdomain "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	vdomain "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	arinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/ar_assets"
	tdminfra "github.com/Misoten-B/airship-backend/internal/infrastructure/three_dimentional_model"
	vinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/voice"
	"github.com/google/wire"
	"gorm.io/gorm"
)

/** Usecase Provider Set **/

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
	GormARAssetsRepositorySet,
	AzureQRCodeImageStorageSet,
	ExternalAPIVoiceModelAdapterSet,
	VoiceServiceImplSet,
	GormVoiceRepositorySet,
	ThreeDimentionalModelServiceImplSet,
	GormThreeDimentionalModelRepositorySet,
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

/** Injectors **/

func InitializeCreateARAssetsUsecaseForDev() *arcreateusecase.ARAssetsUsecaseImpl {
	wire.Build(CreateARAssetsUsecaseSetForDev)

	return &arcreateusecase.ARAssetsUsecaseImpl{}
}

func InitializeCreateARAssetsUsecaseForProd(db *gorm.DB, config *config.Config) *arcreateusecase.ARAssetsUsecaseImpl {
	wire.Build(CreateARAssetsUsecaseSetForProd)

	return &arcreateusecase.ARAssetsUsecaseImpl{}
}
