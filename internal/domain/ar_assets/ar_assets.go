package arassets

import (
	"fmt"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/file"
	"github.com/Misoten-B/airship-backend/internal/id"
)

type ARAssets struct {
	speakingAsset           SpeakingAsset
	qRCodeImage             QRCodeImage
	threeDimentionalModelID id.ID
	accessCount             int
}

func NewARAssets(
	speakingAsset SpeakingAsset,
	qrCodeImage QRCodeImage,
	threedimentionalModelID id.ID,
) ARAssets {
	return ARAssets{
		speakingAsset:           speakingAsset,
		qRCodeImage:             qrCodeImage,
		threeDimentionalModelID: threedimentionalModelID,
		accessCount:             0,
	}
}

// ID は ARAssets の ID を返します。
// ARアセット・音声ファイルともに1対1の関係なので、内部的には SpeakingAsset の ID を返します。
func (a *ARAssets) ID() id.ID {
	return a.speakingAsset.ID()
}

func (a *ARAssets) UserID() id.ID {
	return a.speakingAsset.UserID()
}

func (a *ARAssets) QRCodeImage() QRCodeImage {
	return a.qRCodeImage
}

func (a *ARAssets) ThreeDimentionalModelID() id.ID {
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
	userID      id.ID
	description string
	audioPath   string
}

func NewSpeakingAsset(userID id.ID, description string) (SpeakingAsset, error) {
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

func (s *SpeakingAsset) UserID() id.ID {
	return s.userID
}

func (s *SpeakingAsset) Description() string {
	return s.description
}

func (s *SpeakingAsset) AudioPath() string {
	return s.audioPath
}

type QRCodeImage struct {
	name string
	file *file.File
}

func NewQRCodeImage(file *file.File) (QRCodeImage, error) {
	id, err := id.NewID()
	if err != nil {
		return QRCodeImage{}, err
	}

	name := fmt.Sprintf("%s%s", id.String(), filepath.Ext(file.FileHeader().Filename))

	return QRCodeImage{
		name: name,
		file: file,
	}, nil
}

func (q *QRCodeImage) Name() string {
	return q.name
}

func (q *QRCodeImage) File() *file.File {
	return q.file
}
