package threedimentionalmodel

type ReadModel struct {
	id     string
	userID string
	path   string
}

func NewReadModel(
	id string,
	userID string,
	path string,
) ReadModel {
	return ReadModel{
		id:     id,
		userID: userID,
		path:   path,
	}
}

func NewTemplateReadModel(
	id string,
	path string,
) ReadModel {
	return ReadModel{
		id:     id,
		userID: "",
		path:   path,
	}
}

func (r *ReadModel) ID() string {
	return r.id
}

func (r *ReadModel) UserID() string {
	return r.userID
}

func (r *ReadModel) Path() string {
	return r.path
}

func (r *ReadModel) IsTemplate() bool {
	return r.userID == ""
}
