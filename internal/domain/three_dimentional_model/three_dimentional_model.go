package threedimentionalmodel

import (
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type ThreeDimentionalModel struct {
	id     id.ID
	userID id.ID
	file   file.File
}

func NewThreeDimentionalModel(userID id.ID, file file.File) (*ThreeDimentionalModel, error) {
	id, err := id.NewID()
	if err != nil {
		return nil, err
	}

	return &ThreeDimentionalModel{
		id:     id,
		userID: userID,
		file:   file,
	}, nil
}

func (t *ThreeDimentionalModel) ID() id.ID {
	return t.id
}

func (t *ThreeDimentionalModel) UserID() id.ID {
	return t.userID
}

func (t *ThreeDimentionalModel) File() file.File {
	return t.file
}
