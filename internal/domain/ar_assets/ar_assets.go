package arassets

import (
	"fmt"
	"path/filepath"

	"github.com/Misoten-B/airship-backend/internal/domain/shared"
	"github.com/Misoten-B/airship-backend/internal/file"
)

type ARAssets struct {
	speakingAsset           SpeakingAsset
	qRCodeImage             QRCodeImage
	threeDimentionalModelID shared.ID
	accessCount             int
}

func NewARAssets(
	speakingAsset SpeakingAsset,
	qrCodeImage QRCodeImage,
	threedimentionalModelID shared.ID,
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
func (a *ARAssets) ID() shared.ID {
	return a.speakingAsset.ID()
}

func (a *ARAssets) UserID() shared.ID {
	return a.speakingAsset.UserID()
}

func (a *ARAssets) QRCodeImage() QRCodeImage {
	return a.qRCodeImage
}

func (a *ARAssets) ThreeDimentionalModelID() shared.ID {
	return a.threeDimentionalModelID
}

func (a *ARAssets) SpeakingAsset() SpeakingAsset {
	return a.speakingAsset
}

func (a *ARAssets) AccessCount() int {
	return a.accessCount
}

type SpeakingAsset struct {
	id          shared.ID
	userID      shared.ID
	description string
	audioPath   string
}

func NewSpeakingAsset(userID shared.ID, description string) (SpeakingAsset, error) {
	id, err := shared.NewID()
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

func (s *SpeakingAsset) ID() shared.ID {
	return s.id
}

func (s *SpeakingAsset) UserID() shared.ID {
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
	id, err := shared.NewID()
	if err != nil {
		return QRCodeImage{}, err
	}

	// AzureDriverの仕様上、一旦ここでファイル名を変更しています。
	name := fmt.Sprintf("%s%s", id.String(), filepath.Ext(file.FileHeader().Filename))
	file.FileHeader().Filename = name

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
