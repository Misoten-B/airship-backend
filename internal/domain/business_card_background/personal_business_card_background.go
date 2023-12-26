package businesscardbackground

import (
	"errors"
	"regexp"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/file"
)

type PersonalBusinessCardBackground struct {
	id        shared.ID
	colorCode string
	image     file.File
}

const (
	// #000000 or #000
	hexRegexString = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
)

var (
	hexRegex    = regexp.MustCompile(hexRegexString)
	ErrBadColor = errors.New("bad color code")
)

func NewPersonalBusinessCardBackground(colorCode string, image file.File) (*PersonalBusinessCardBackground, error) {
	if !hexRegex.MatchString(colorCode) {
		return nil, ErrBadColor
	}

	id, err := shared.NewID()
	if err != nil {
		return nil, err
	}

	return &PersonalBusinessCardBackground{
		id:        id,
		colorCode: colorCode,
		image:     image,
	}, nil
}

func (p *PersonalBusinessCardBackground) ID() shared.ID {
	return p.id
}

func (p *PersonalBusinessCardBackground) ColorCode() string {
	return p.colorCode
}

func (p *PersonalBusinessCardBackground) Image() file.File {
	return p.image
}
