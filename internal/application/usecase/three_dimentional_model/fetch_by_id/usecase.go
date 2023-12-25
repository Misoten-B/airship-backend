package fetchbyid

import (
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
)

type Usecase interface {
	Execute(input Input) (Output, error)
}

type Input struct {
	ID     string
	UserID string
}

type Output struct {
	ID   string
	Path string
}

// Interactor はARアセットID検索ユースケースの実装です。
type Interactor struct {
	threeDimentionalModelRepository tdmservice.ThreeDimentionalModelRepository
	threeDimentionalModelStorage    tdmservice.ThreeDimentionalModelStorage
}

func NewInteractor(
	threeDimentionalModelRepository tdmservice.ThreeDimentionalModelRepository,
	threeDimentionalModelStorage tdmservice.ThreeDimentionalModelStorage,
) *Interactor {
	return &Interactor{
		threeDimentionalModelRepository: threeDimentionalModelRepository,
		threeDimentionalModelStorage:    threeDimentionalModelStorage,
	}
}

// Execute はARアセットID検索ユースケースを実行します。
func (i *Interactor) Execute(input Input) (Output, error) {
	var output Output

	//   バリデーション&オブジェクト生成
	tdmID := shared.ReconstructID(input.ID)

	//   IDをもとに3Dモデルを取得
	readModel, err := i.threeDimentionalModelRepository.FindByID(tdmID)
	if err != nil {
		msg := "failed to find three dimentional model"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	//   権限確認（ユーザー定義の場合はそれが自分のものか）
	if !readModel.IsTemplate() {
		if readModel.UserID() != input.UserID {
			msg := "user does not have permission to use this three dimentional model"
			return output, customerror.NewApplicationError(
				nil,
				msg,
				http.StatusForbidden,
			)
		}
	}

	//   URL生成
	path, err := i.threeDimentionalModelStorage.GetModelURL(readModel.Path())
	if err != nil {
		msg := "failed to get 3D model URL"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	return Output{
		ID:   readModel.ID(),
		Path: path,
	}, nil
}
