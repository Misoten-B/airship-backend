package update

import (
	"mime/multipart"
	"net/http"

	tdmappservice "github.com/Misoten-B/airship-backend/internal/application/applicationservice/threedimensionalmodel"
	"github.com/Misoten-B/airship-backend/internal/customerror"
	tdmdomain "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type Usecase interface {
	Execute(input Input) error
}

type Input struct {
	ID         string
	UserID     string
	File       multipart.File
	FileHeader *multipart.FileHeader
}

// Interactor はARアセット更新ユースケースの実装です。
type Interactor struct {
	threeDimentionalModelStorage    tdmservice.ThreeDimentionalModelStorage
	threeDimentionalModelRepository tdmservice.ThreeDimentionalModelRepository
}

func NewInteractor(
	threeDimentionalModelStorage tdmservice.ThreeDimentionalModelStorage,
	threeDimentionalModelRepository tdmservice.ThreeDimentionalModelRepository,
) *Interactor {
	return &Interactor{
		threeDimentionalModelStorage:    threeDimentionalModelStorage,
		threeDimentionalModelRepository: threeDimentionalModelRepository,
	}
}

// Execute はARアセット更新ユースケースを実行します。
func (i *Interactor) Execute(input Input) error {
	// バリデーション & オブジェクト生成
	tdmID := id.ReconstructID(input.ID)
	userID := id.ReconstructID(input.UserID)

	file := file.NewMyFile(input.File, input.FileHeader)

	// 3Dモデルの取得
	threeDimensionalModel, err := i.threeDimentionalModelRepository.Find(tdmID)
	if err != nil {
		msg := "failed to find 3D model"
		return customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	// 権限確認
	if !tdmappservice.HasUpdatePermission(*threeDimensionalModel, userID) {
		msg := "you don't have permission to update this 3D model"
		return customerror.NewApplicationErrorWithoutDetails(
			msg,
			http.StatusForbidden,
		)
	}

	// 3Dモデルファイルの復元
	filePath := threeDimensionalModel.Path()
	modelFile := tdmdomain.ReconstructThreeDimensionalModelFile(file, filePath)

	// 3Dモデルの更新
	err = i.threeDimentionalModelStorage.Save(modelFile)
	if err != nil {
		msg := "failed to save 3D model file"
		return customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	return nil
}
