package threedimentionalmodel

import (
	"fmt"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type ThreeDimentionalModel struct {
	id     id.ID
	userID id.ID
	file   *file.File
}

func NewThreeDimentionalModel(userID id.ID, file *file.File) (ThreeDimentionalModel, error) {
	id, err := id.NewID()
	if err != nil {
		return ThreeDimentionalModel{}, err
	}

	// AzureDriverの仕様上、一旦ここでファイル名を変更しています。
	file.FileHeader().Filename = fmt.Sprintf("%s%s", id.String(), filepath.Ext(file.FileHeader().Filename))

	return ThreeDimentionalModel{
		id:     id,
		userID: userID,
		file:   file,
	}, nil
}

func ReconstructThreeDimentionalModel(id id.ID, userID id.ID) *ThreeDimentionalModel {
	return &ThreeDimentionalModel{
		id:     id,
		userID: userID,
	}
}

func ReconstructThreeDimentionalModelTemplate(id id.ID) *ThreeDimentionalModel {
	return &ThreeDimentionalModel{
		id: id,
	}
}

func (t *ThreeDimentionalModel) IsTemplate() bool {
	return t.userID == ""
}

func (t *ThreeDimentionalModel) ID() id.ID {
	return t.id
}

func (t *ThreeDimentionalModel) UserID() id.ID {
	return t.userID
}

func (t *ThreeDimentionalModel) File() *file.File {
	return t.file
}

func (t *ThreeDimentionalModel) FileName() string {
	return t.file.FileHeader().Filename
}
