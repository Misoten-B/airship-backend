//go:build wireinject
// +build wireinject

package container

import (
	"github.com/Misoten-B/airship-backend/config"
	arusecase "github.com/Misoten-B/airship-backend/internal/application/usecase/ar_assets"
	ardomain "github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	tdmdomain "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	vdomain "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	arinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/ar_assets"
	tdminfra "github.com/Misoten-B/airship-backend/internal/infrastructure/three_dimentional_model"
	vinfra "github.com/Misoten-B/airship-backend/internal/infrastructure/voice"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// CreateARAssetsUsecaseForDev は開発環境用のARAssetsUsecaseを生成します。
// 現状、このファイルに記述されていますが、将来的にはユースケースのファクトリ部分に移動することも
// 検討しています。
var CreateARAssetsUsecaseForDev = wire.NewSet(
	arusecase.NewARAssetsUsecaseImpl,
	ardomain.NewMockARAssetsRepository,
	wire.Bind(new(ardomain.ARAssetsRepository), new(*ardomain.MockARAssetsRepository)),
	ardomain.NewMockQRCodeImageStorage,
	wire.Bind(new(ardomain.QRCodeImageStorage), new(*ardomain.MockQRCodeImageStorage)),
	vdomain.NewMockVoiceModelAdapter,
	wire.Bind(new(vdomain.VoiceModelAdapter), new(*vdomain.MockVoiceModelAdapter)),
	vdomain.NewVoiceServiceImpl,
	wire.Bind(new(vdomain.VoiceService), new(*vdomain.VoiceServiceImpl)),
	vdomain.NewMockVoiceRepository,
	wire.Bind(new(vdomain.VoiceRepository), new(*vdomain.MockVoiceRepository)),
	tdmdomain.NewThreeDimentionalModelServiceImpl,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelService), new(*tdmdomain.ThreeDimentionalModelServiceImpl)),
	tdmdomain.NewMockThreeDimentionalModelRepository,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelRepository), new(*tdmdomain.MockThreeDimentionalModelRepository)),
)

var CreateARAssetsUsecaseForProd = wire.NewSet(
	arusecase.NewARAssetsUsecaseImpl,
	arinfra.NewGormARAssetsRepository,
	wire.Bind(new(ardomain.ARAssetsRepository), new(*arinfra.GormARAssetsRepository)),
	arinfra.NewAzureQRCodeImageStorage,
	wire.Bind(new(ardomain.QRCodeImageStorage), new(*arinfra.AzureQRCodeImageStorage)),
	vinfra.NewExternalAPIVoiceModelAdapter,
	wire.Bind(new(vdomain.VoiceModelAdapter), new(*vinfra.ExternalAPIVoiceModelAdapter)),
	vdomain.NewVoiceServiceImpl,
	wire.Bind(new(vdomain.VoiceService), new(*vdomain.VoiceServiceImpl)),
	vinfra.NewGormVoiceRepository,
	wire.Bind(new(vdomain.VoiceRepository), new(*vinfra.GormVoiceRepository)),
	tdmdomain.NewThreeDimentionalModelServiceImpl,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelService), new(*tdmdomain.ThreeDimentionalModelServiceImpl)),
	tdminfra.NewGormThreeDimentionalModelRepository,
	wire.Bind(new(tdmdomain.ThreeDimentionalModelRepository), new(*tdminfra.GormThreeDimentionalModelRepository)),
)

func InitializeCreateARAssetsUsecaseForDev() *arusecase.ARAssetsUsecaseImpl {
	wire.Build(CreateARAssetsUsecaseForDev)

	return &arusecase.ARAssetsUsecaseImpl{}
}

func InitializeCreateARAssetsUsecaseForProd(db *gorm.DB, config *config.Config) *arusecase.ARAssetsUsecaseImpl {
	wire.Build(CreateARAssetsUsecaseForProd)

	return &arusecase.ARAssetsUsecaseImpl{}
}
