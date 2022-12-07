package core

import "github.com/argSea/portfolio_blog_api/argHex/core"

//User repo to connect to a store
type UserRepo interface {
	GetUserByID(id string) core.User
}
