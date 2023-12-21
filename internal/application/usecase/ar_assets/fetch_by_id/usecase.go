package fetchbyid

import (
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	vservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type Usecase interface {
	Execute(input Input) (Output, error)
}

type Input struct {
	ID     string
	UserID string
}

type Output struct {
	ID                   string
	SpeakingDescription  string
	SpeakingAudioPath    string
	ThreeDimentionalPath string
	QrcodeIconImagePath  string
	IsCompleted          bool
}

// Interactor はARアセットID検索ユースケースの実装です。
type Interactor struct {
	arAssetsRepository           service.ARAssetsRepository
	qrCodeImageStorage           service.QRCodeImageStorage
	speakingAudioStorage         vservice.SpeakingAudioStorage
	threeDimentionalModelStorage tdmservice.ThreeDimentionalModelStorage
}

func NewInteractor(
	arAssetsRepository service.ARAssetsRepository,
	qrCodeImageStorage service.QRCodeImageStorage,
	speakingAudioStorage vservice.SpeakingAudioStorage,
	threeDimentionalModelStorage tdmservice.ThreeDimentionalModelStorage,
) *Interactor {
	return &Interactor{
		arAssetsRepository:           arAssetsRepository,
		qrCodeImageStorage:           qrCodeImageStorage,
		speakingAudioStorage:         speakingAudioStorage,
		threeDimentionalModelStorage: threeDimentionalModelStorage,
	}
}

// Execute はARアセットID検索ユースケースを実行します。
func (i *Interactor) Execute(input Input) (Output, error) {
	var output Output

	// バリデーション & オブジェクト生成
	id := id.ReconstructID(input.ID)

	// リポジトリから取得
	model, err := i.arAssetsRepository.FetchByID(id)
	if err != nil {
		msg := "failed to fetch AR assets"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	// 権限確認
	if model.UID() != input.UserID {
		return output, customerror.NewApplicationErrorWithoutDetails(
			"user does not have permission to use this AR assets",
			http.StatusForbidden,
		)
	}

	// URL生成
	speakingAudioPath, err := i.speakingAudioStorage.GetAudioURL(model.SpeakingAudioPath())
	if err != nil {
		msg := "failed to get speaking audio URL"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	threeDimentionalPath, err := i.threeDimentionalModelStorage.GetModelURL(model.ThreeDimentionalPath())
	if err != nil {
		msg := "failed to get 3D model URL"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	qrcodeIconImagePath, err := i.qrCodeImageStorage.GetImageURL(model.QrcodeIconImagePath())
	if err != nil {
		msg := "failed to get QR code image URL"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	return Output{
		ID:                   model.ID(),
		SpeakingDescription:  model.SpeakingDescription(),
		SpeakingAudioPath:    speakingAudioPath,
		ThreeDimentionalPath: threeDimentionalPath,
		QrcodeIconImagePath:  qrcodeIconImagePath,
		IsCompleted:          model.IsCreated(),
	}, nil
}
