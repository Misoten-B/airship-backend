package voice

import (
	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/file"
)

type File struct {
	file file.File
}

func NewFile(file file.File) (*File, error) {
	return &File{
		file: file,
	}, nil
}

func (v *File) File() file.File {
	return v.file
}

type Voice struct {
	id        shared.ID
	voice     File
	modelPath string
}

func NewVoice(voice File, modelPath string) (*Voice, error) {
	id, err := shared.NewID()
	if err != nil {
		return nil, err
	}

	return &Voice{
		id:        id,
		voice:     voice,
		modelPath: modelPath,
	}, nil
}

func (v *Voice) ID() shared.ID {
	return v.id
}

func (v *Voice) Voice() File {
	return v.voice
}

func (v *Voice) ModelPath() string {
	return v.modelPath
}
