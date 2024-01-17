package fetchbyuserid

import (
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	vservice "github.com/Misoten-B/airship-backend/internal/domain/voice/service"
)

type Usecase interface {
	Execute(input Input) (Output, error)
}

type Input struct {
	UserID string
}

type item struct {
	ID                   string
	SpeakingDescription  string
	SpeakingAudioPath    string
	ThreeDimentionalID   string
	ThreeDimentionalPath string
	QrcodeIconImagePath  string
	IsCompleted          bool
}

type Output struct {
	Items []item
}

// Interactor はARアセット一覧取得ユースケースの実装です。
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

// Execute はARアセット一覧取得ユースケースを実行します。
func (i *Interactor) Execute(input Input) (Output, error) {
	var output Output

	// バリデーション & オブジェクト生成
	userID := shared.ReconstructID(input.UserID)

	// userIDをもとに一覧取得
	models, err := i.arAssetsRepository.FetchByUserID(userID)
	if err != nil {
		msg := "failed to fetch AR assets"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	// 取得結果が空の場合は空スライスを返す
	if len(models) == 0 {
		return Output{
			Items: []item{},
		}, nil
	}

	// 各URLのパスを取得する
	qrCodeImageFullPath, err := i.qrCodeImageStorage.GetContainerFullPath()
	if err != nil {
		msg := "failed to get qr code image full path"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	speakingAudioFullPath, err := i.speakingAudioStorage.GetContainerFullPath()
	if err != nil {
		msg := "failed to get speaking audio full path"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	threeDimentionalModelFullPath, err := i.threeDimentionalModelStorage.GetContainerFullPath()
	if err != nil {
		msg := "failed to get three dimentional model full path"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	items := []item{}
	for _, model := range models {
		var qrCodeImagePath string
		if model.QrcodeIconImagePath() != "" {
			qrCodeImagePath = qrCodeImageFullPath.Path(model.QrcodeIconImagePath())
		}

		element := item{
			ID:                   model.ID(),
			SpeakingDescription:  model.SpeakingDescription(),
			SpeakingAudioPath:    speakingAudioFullPath.Path(model.SpeakingAudioPath()),
			ThreeDimentionalID:   model.ThreeDimentionalID(),
			ThreeDimentionalPath: threeDimentionalModelFullPath.Path(model.ThreeDimentionalPath()),
			QrcodeIconImagePath:  qrCodeImagePath,
			IsCompleted:          model.IsCreated(),
		}
		items = append(items, element)
	}

	return Output{
		Items: items,
	}, nil
}
