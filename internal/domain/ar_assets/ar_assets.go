package arassets

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/id"
)

type ARAssets struct {
	speakingAsset           SpeakingAsset
	qRCodeImage             QRCodeImage
	threeDimentionalModelID string
	accessCount             int
}

func NewARAssets(
	speakingAsset SpeakingAsset,
	qrCodeImage QRCodeImage,
	threedimentionalModelID string,
) ARAssets {
	return ARAssets{
		speakingAsset:           speakingAsset,
		qRCodeImage:             qrCodeImage,
		threeDimentionalModelID: threedimentionalModelID,
		accessCount:             0,
	}
}

// ID は ARAssets の ID を返します。
// ARアセット・音声ともに一意なので、内部的には SpeakingAsset の ID を返します。
func (a *ARAssets) ID() id.ID {
	return a.speakingAsset.ID()
}

func (a *ARAssets) UserID() string {
	return a.speakingAsset.UserID()
}

func (a *ARAssets) QRCodeImage() QRCodeImage {
	return a.qRCodeImage
}

func (a *ARAssets) ThreeDimentionalModelID() string {
	return a.threeDimentionalModelID
}

func (a *ARAssets) SpeakingAsset() SpeakingAsset {
	return a.speakingAsset
}

func (a *ARAssets) AccessCount() int {
	return a.accessCount
}

type SpeakingAsset struct {
	id          id.ID
	userID      string
	description string
	audioPath   string
}

func NewSpeakingAsset(userID string, description string) (SpeakingAsset, error) {
	id, err := id.NewID()
	if err != nil {
		return SpeakingAsset{}, err
	}

	audioPath := fmt.Sprintf("%s.wav", id.String())

	return SpeakingAsset{
		id:          id,
		userID:      userID,
		description: description,
		audioPath:   audioPath,
	}, nil
}

func (s *SpeakingAsset) ID() id.ID {
	return s.id
}

func (s *SpeakingAsset) UserID() string {
	return s.userID
}

func (s *SpeakingAsset) Description() string {
	return s.description
}

func (s *SpeakingAsset) AudioPath() string {
	return s.audioPath
}

type QRCodeImage struct {
	id         id.ID
	file       multipart.File
	fileHeader *multipart.FileHeader
}

func NewQRCodeImage(file multipart.File, fileHeader *multipart.FileHeader) (QRCodeImage, error) {
	id, err := id.NewID()
	if err != nil {
		return QRCodeImage{}, err
	}

	return QRCodeImage{
		id:         id,
		file:       file,
		fileHeader: fileHeader,
	}, nil
}

func (q *QRCodeImage) Name() string {
	return fmt.Sprintf("%s%s", q.id.String(), filepath.Ext(q.fileHeader.Filename))
}

func (q *QRCodeImage) File() multipart.File {
	return q.file
}
