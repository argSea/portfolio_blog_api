package usecase

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
)

//Concrete for use case
type userCase struct {
	userRepo core.UserRepository
}

func NewUserCase(repo core.UserRepository) core.UserUsecase {
	return &userCase{
		userRepo: repo,
	}
}

func (u *userCase) GetUserByID(id string) (*entity.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userCase) GetUserByUserName(userName string) (*entity.User, error) {
	return u.userRepo.GetUserByUserName(userName)
}

func (u *userCase) Save(newUser entity.User) (*entity.User, error) {
	return u.userRepo.Save(newUser)
}

func (u *userCase) Update(newUser entity.User) (*entity.User, error) {
	return u.userRepo.Update(newUser)
}

func (u *userCase) Delete(id string) error {
	return u.userRepo.Delete(id)
}

// func (u *userCase) Decode(body io.ReadCloser) entity.User {
// 	//Decode
// 	newUser := entity.User{}
// 	decoder := json.NewDecoder(body)
// 	decoder.Decode(&newUser)

// 	return newUser
// }
