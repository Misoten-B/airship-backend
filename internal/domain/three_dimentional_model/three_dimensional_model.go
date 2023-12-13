package threedimentionalmodel

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type ThreeDimensionalModel struct{}

func NewThreeDimensionalModel() ThreeDimensionalModel {
	return ThreeDimensionalModel{}
}

func ReconstructThreeDimensionalModel() *ThreeDimensionalModel {
	return &ThreeDimensionalModel{}
}

// ThreeDimensionalModelFile は3Dモデルのファイルを表す構造体です。
type ThreeDimensionalModelFile struct {
	file *file.File
	path shared.FilePath
}

func NewThreeDimensionalModelFile(
	file *file.File,
) (ThreeDimensionalModelFile, error) {
	uid, err := id.NewID()
	if err != nil {
		return ThreeDimensionalModelFile{}, err
	}

	fileName := fmt.Sprintf("%s%s", uid.String(), filepath.Ext(file.FileHeader().Filename))
	path := shared.NewFilePath(fileName)

	return ThreeDimensionalModelFile{
		file: file,
		path: path,
	}, nil
}

// NewThreeDimensionalModelFileFromPath は既にファイルが存在する場合に使用します。
func NewThreeDimensionalModelFileFromPath(
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
