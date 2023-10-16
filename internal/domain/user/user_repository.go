package domain

type UserRepository interface {
	Create(id string, modelKey string) User
	ReadById(id string) User
	Update(id string) User
	Delete(id string) User
}
