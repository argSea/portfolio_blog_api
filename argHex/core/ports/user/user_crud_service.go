package core

import "github.com/argSea/portfolio_blog_api/argHex/core"

//User service for CRUD
type UserCRUDService interface {
	Create(user core.User) error
	Read(id string) interface{}
	Update(user core.User) error
	Delete(user core.User) error
}
