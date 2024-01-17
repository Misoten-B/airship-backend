package arassets

type ReadModel struct {
	id                   string
	uID                  string
	speakingDescription  string
	speakingAudioPath    string
	threeDimentionalID   string
	threeDimentionalPath string
	qrcodeIconImagePath  string
	isCreated            bool
}

func NewReadModel(
	id string,
	uID string,
	speakingDescription string,
	speakingAudioPath string,
	threeDimentionalID string,
	threeDimentionalPath string,
	qrcodeIconImagePath string,
	isCreated bool,
) ReadModel {
	return ReadModel{
		id:                   id,
		uID:                  uID,
		speakingDescription:  speakingDescription,
		speakingAudioPath:    speakingAudioPath,
		threeDimentionalID:   threeDimentionalID,
		threeDimentionalPath: threeDimentionalPath,
		qrcodeIconImagePath:  qrcodeIconImagePath,
		isCreated:            isCreated,
	}
}

func (r *ReadModel) ID() string {
	return r.id
}

func (r *ReadModel) UID() string {
	return r.uID
}

func (r *ReadModel) SpeakingDescription() string {
	return r.speakingDescription
}

func (r *ReadModel) SpeakingAudioPath() string {
	return r.speakingAudioPath
}

func (r *ReadModel) ThreeDimentionalID() string {
	return r.threeDimentionalID
}

func (r *ReadModel) ThreeDimentionalPath() string {
	return r.threeDimentionalPath
}

func (r *ReadModel) QrcodeIconImagePath() string {
	return r.qrcodeIconImagePath
}

func (r *ReadModel) IsCreated() bool {
	return r.isCreated
}
