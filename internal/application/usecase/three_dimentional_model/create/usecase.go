package create

import (
	"mime/multipart"
	"net/http"

	"github.com/Misoten-B/airship-backend/internal/customerror"
	tdmdomain "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model"
	tdmservice "github.com/Misoten-B/airship-backend/internal/domain/three_dimentional_model/service"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type Usecase interface {
	Execute(input Input) (Output, error)
}

type Input struct {
	UserID     string
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type Output struct {
	ID string
}

// Interactor はARアセット作成ユースケースの実装です。
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

// Execute はARアセット作成ユースケースを実行します。
func (i *Interactor) Execute(input Input) (Output, error) {
	var output Output

	//	バリデーション&オブジェクト生成
	userID := id.ReconstructID(input.UserID)

	file := file.NewMyFile(input.File, input.FileHeader)
	modelFile, err := tdmdomain.NewThreeDimensionalModelFile(file)
	if err != nil {
		return output, customerror.NewApplicationErrorWithoutDetails(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	threeDimentionalModel, err := tdmdomain.NewThreeDimensionalModel(
		userID,
		modelFile.Path(),
	)
	if err != nil {
		return output, customerror.NewApplicationErrorWithoutDetails(
			err.Error(),
			http.StatusBadRequest,
		)
	}

	//  3Dモデルの保存
	err = i.threeDimentionalModelStorage.Save(modelFile)
	if err != nil {
		msg := "failed to save 3d model file"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	//  永続化
	err = i.threeDimentionalModelRepository.Save(threeDimentionalModel)
	if err != nil {
		msg := "failed to save 3d model"
		return output, customerror.NewApplicationError(
			err,
			msg,
			http.StatusInternalServerError,
		)
	}

	return Output{
		ID: threeDimentionalModel.ID().String(),
	}, nil
}
