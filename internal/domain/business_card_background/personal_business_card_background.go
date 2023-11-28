package businesscardbackground

import (
	"errors"
	"regexp"

	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type PersonalBusinessCardBackground struct {
	id        id.ID
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

	id, err := id.NewID()
	if err != nil {
		return nil, err
	}

	return &PersonalBusinessCardBackground{
		id:        id,
		colorCode: colorCode,
		image:     image,
	}, nil
}

func (p *PersonalBusinessCardBackground) ID() id.ID {
	return p.id
}

func (p *PersonalBusinessCardBackground) ColorCode() string {
	return p.colorCode
}

func (p *PersonalBusinessCardBackground) Image() file.File {
	return p.image
}
