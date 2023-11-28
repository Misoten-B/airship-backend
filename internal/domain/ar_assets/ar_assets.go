package arassets

import (
	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type ARAssets struct {
	ID             id.ID
	UserID         string
	AudioModelPath string
	QRCodeImage    *file.File
	AudioText      string
	AudioID        string
	AudioPath      string
}

func NewARAssets(
	userID string,
	audioModelPath string,
	qrCodeImage *file.File,
	audioText string,
) (*ARAssets, error) {
	id, err := id.NewID()
	if err != nil {
		return nil, err
	}

	return &ARAssets{
		ID:             id,
		UserID:         userID,
		AudioModelPath: audioModelPath,
		QRCodeImage:    qrCodeImage,
		AudioText:      audioText,
	}, nil
}

func (a *ARAssets) UpdateARAssets(audioID string, audioPath string) {
	a.AudioID = audioID
	a.AudioPath = audioPath
}
