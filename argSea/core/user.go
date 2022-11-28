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
	GetUserByID(string) (*entity.User, error)
	GetUserByUserName(string) (*entity.User, error)
	Save(entity.User) (*entity.User, error)
	Update(entity.User) (*entity.User, error)
	Delete(string) error
	// Decode(io.ReadCloser) User
}

type UserPresenter interface {
	Present() *entity.User
}
