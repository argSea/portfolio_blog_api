package adapters

import (
	"fmt"

	core "github.com/argSea/portfolio_blog_api/argHex/core/user"
)

type resumeFakeOutAdapter struct {
}

func NewResumeFakeOutAdapter() core.UserRepo {
	return resumeFakeOutAdapter{}
}

func (u resumeFakeOutAdapter) GetUserByID(id string) core.User {
	user := core.User{}
	user.SetID("12345").
		SetUserName("testUserName").
		SetFirstName("testFirstName")

	fmt.Println(user)

	return user
}
