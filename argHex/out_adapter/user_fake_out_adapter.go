package out_adapter

import (
	"fmt"

	"github.com/argSea/portfolio_blog_api/argHex/domain"
	"github.com/argSea/portfolio_blog_api/argHex/out_port"
)

type userFakeOutAdapter struct {
}

func NewUserFakeOutAdapter() out_port.UserRepo {
	return userFakeOutAdapter{}
}

func (u userFakeOutAdapter) Login(user domain.User) (string, error) {
	return "12345", nil
}

func (u userFakeOutAdapter) Get(id string) domain.User {
	user := domain.User{}
	user.Id = "12345"
	user.UserName = "testUserName"
	user.FirstName = "testFirstName"

	fmt.Println(user)

	return user
}

func (u userFakeOutAdapter) GetByUserName(username string) domain.User {
	user := domain.User{}
	user.Id = "12345"
	user.UserName = "testUserName"
	user.FirstName = "testFirstName"

	fmt.Println(user)

	return user
}

func (u userFakeOutAdapter) Set(user domain.User) error {
	return nil
}

func (u userFakeOutAdapter) Add(user domain.User) (string, error) {
	return "", nil
}

func (u userFakeOutAdapter) Remove(user domain.User) error {
	return nil
}
