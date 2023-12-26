package statusdone

import (
	"errors"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	"github.com/Misoten-B/airship-backend/internal/domain/ar_assets/service"
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
)

type Usecase interface {
	Execute(input Input) error
}

type Input struct {
	ID string
}

// Interactor はARアセット生成完了ユースケースの実装です。
type Interactor struct {
	arAssetsRepository service.ARAssetsRepository
}

func NewInteractor(
	arAssetsRepository service.ARAssetsRepository,
) *Interactor {
	return &Interactor{
		arAssetsRepository: arAssetsRepository,
	}
}

// Execute はARアセット生成完了ユースケースを実行します。
func (i *Interactor) Execute(input Input) error {
	// バリデーション & オブジェクト生成
	id := shared.ReconstructID(input.ID)
	statusDone := shared.StatusCompleted{}

	// ARアセットの取得
	_, err := i.arAssetsRepository.FetchByID(id)
	if err != nil {
		if errors.Is(err, service.ErrArAssetsNotFound) {
			return customerror.NewApplicationErrorWithoutDetails(
				"AR assets not found",
				http.StatusNotFound,
			)
		}
		return customerror.NewApplicationErrorWithoutDetails(
			"failed to fetch ar assets",
			http.StatusNotFound,
		)
	}

	// ステータスの更新
	if err = i.arAssetsRepository.PatchStatus(id, statusDone); err != nil {
		msg := "failed to patch status"
		return customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	return nil
}
