package fetchbyuserid

import (
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type Usecase interface {
	Execute(input Input) (Output, error)
}

type Input struct {
	UserID string
}

type item struct {
	ID   string
	Path string
}

type Output struct {
	Items []item
}

// Interactor はARアセット一覧取得ユースケースの実装です。
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

// Execute はARアセット一覧取得ユースケースを実行します。
func (i *Interactor) Execute(input Input) (Output, error) {
	var output Output

	// バリデーション&オブジェクト生成
	userID := id.ReconstructID(input.UserID)

	// userIDをもとに3Dモデルのテンプレートとユーザー定義モデルを取得
	readModels, err := i.threeDimentionalModelRepository.FindByUserID(userID)
	if err != nil {
		msg := "failed to find three dimentional model"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	if len(readModels) == 0 {
		return Output{
			Items: []item{},
		}, nil
	}

	// コンテナのURL生成
	fullPath, err := i.threeDimentionalModelStorage.GetContainerFullPath()
	if err != nil {
		msg := "failed to get container full path"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	// レスポンス生成
	items := []item{}
	for _, readModel := range readModels {
		items = append(items, item{
			ID:   readModel.ID(),
			Path: fullPath.Path(readModel.Path()),
		})
	}

	return Output{
		Items: items,
	}, nil
}
