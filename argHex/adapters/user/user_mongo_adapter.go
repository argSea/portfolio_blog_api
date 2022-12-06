package userAdapters

import core "github.com/argSea/portfolio_blog_api/argHex/core/user"

type UserMongoAdapter struct {
}

func NewUserMongoAdapter() core.UserRepo {
	u := UserMongoAdapter{}

	return u
}

func (u UserMongoAdapter) GetUserByID(id string) core.User {
	return core.User{}
}
