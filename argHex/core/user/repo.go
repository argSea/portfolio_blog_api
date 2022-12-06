package core

//User repo to connect to a store
type UserRepo interface {
	GetUserByID(id string) User
}
