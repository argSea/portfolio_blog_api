package core

import "github.com/argSea/portfolio_blog_api/argSea/entity"

//User repo interface
type UserRepository interface {
	GetUserByID(string) (*entity.User, error)
	GetUserByUserName(string) (*entity.User, error)
	Save(entity.User) (*entity.User, error)
	Update(entity.User) (*entity.User, error)
	Delete(string) error
}

//Use case for the above
type UserUsecase interface {
	GetUserByID(string) (interface{}, error)
	GetUserByUserName(string) (interface{}, error)
	Save(entity.User) (interface{}, error)
	Update(entity.User) (interface{}, error)
	Delete(string) error
	// Decode(io.ReadCloser) User
}

type UserPresenter interface {
	Present(model interface{}) interface{}
}
