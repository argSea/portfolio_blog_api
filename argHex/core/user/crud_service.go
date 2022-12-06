package core

//User service for CRUD
type UserCRUDService interface {
	Create(user User) error
	Read(id string) interface{}
	Update(user User) error
	Delete(user User) error
}
