package threedimentionalmodel

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

// ThreeDimensionalModel は3Dモデルのエンティティです。
// userIDによって、個人用の3Dモデルか、テンプレートかを判別します。
type ThreeDimensionalModel struct {
	id     id.ID
	userID id.ID
	path   shared.FilePath
}

// NewThreeDimensionalModel は新規の3Dモデルを作成する際に使用します。
func NewThreeDimensionalModel(userID id.ID, path shared.FilePath) (ThreeDimensionalModel, error) {
	id, err := id.NewID()
	if err != nil {
		return ThreeDimensionalModel{}, err
	}

	return ThreeDimensionalModel{
		id:     id,
		userID: userID,
		path:   path,
	}, nil
}

// ReconstructThreeDimensionalModel は個人用の3Dモデルを再構築する際に使用します。
func ReconstructThreeDimensionalModel(id id.ID, userID id.ID, path shared.FilePath) *ThreeDimensionalModel {
	return &ThreeDimensionalModel{
		id:     id,
		userID: userID,
		path:   path,
	}
}

// ReconstructThreeDimensionalModelTemplate はテンプレートの3Dモデルを再構築する際に使用します。
func ReconstructThreeDimensionalModelTemplate(
	id id.ID,
	path shared.FilePath,
) *ThreeDimensionalModel {
	return &ThreeDimensionalModel{
		id:   id,
		path: path,
	}
}

// IsTemplate はテンプレートかどうかを判別します。
// テンプレートの場合はtrueを返します。
func (t *ThreeDimensionalModel) IsTemplate() bool {
	return t.userID == ""
}

func (t *ThreeDimensionalModel) ID() id.ID {
	return t.id
}

func (t *ThreeDimensionalModel) UserID() id.ID {
	return t.userID
}

func (t *ThreeDimensionalModel) Path() shared.FilePath {
	return t.path
}

// ThreeDimensionalModelFile は3Dモデルのファイルを表す構造体です。
type ThreeDimensionalModelFile struct {
	file *file.File
	path shared.FilePath
}

func NewThreeDimensionalModelFile(
	file *file.File,
) (ThreeDimensionalModelFile, error) {
	uniqueID, err := id.NewID()
	if err != nil {
		return ThreeDimensionalModelFile{}, err
	}

	// ファイル名は<uniqueID>.<拡張子>となるようにしています。
	fileName := fmt.Sprintf("%s%s", uniqueID.String(), filepath.Ext(file.FileHeader().Filename))
	path := shared.NewFilePath(fileName)

	return ThreeDimensionalModelFile{
		file: file,
		path: path,
	}, nil
}

// ReconstructThreeDimensionalModelFile は既にファイルが存在する場合に使用します。
func ReconstructThreeDimensionalModelFile(
	file *file.File,
	path shared.FilePath,
) ThreeDimensionalModelFile {
	return ThreeDimensionalModelFile{
		file: file,
		path: path,
	}
}

func (t *ThreeDimensionalModelFile) File() multipart.File {
	return t.file.File()
}

func (t *ThreeDimensionalModelFile) Path() shared.FilePath {
	return t.path
}
