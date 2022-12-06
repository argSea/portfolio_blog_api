package userAdapters

import (
	"fmt"

	core "github.com/argSea/portfolio_blog_api/argHex/core/user"
)

type userFakeOutAdapter struct {
}

func NewUserFakeOutAdapter() core.UserRepo {
	return userFakeOutAdapter{}
}

func (u userFakeOutAdapter) GetUserByID(id string) core.User {
	user := core.User{}
	user.SetID("12345").
		SetUserName("testUserName").
		SetFirstName("testFirstName")

	fmt.Println(user)

	return user
}
