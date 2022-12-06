package usecase

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
)

//Concrete for use case
type userAPICase struct {
	userRepo      core.UserRepository
	userPresenter core.UserPresenter
}

func NewAPIUserCase(repo core.UserRepository, pres core.UserPresenter) core.UserUsecase {
	return &userAPICase{
		userRepo:      repo,
		userPresenter: pres,
	}
}

func (u *userAPICase) GetUserByID(id string) (interface{}, error) {
	u_data, err := u.userRepo.GetUserByID(id)

	if nil != err {
		return nil, err
	}

	view := u.userPresenter.Present(u_data)
	return view, nil
}

func (u *userAPICase) GetUserByUserName(userName string) (interface{}, error) {
	u_data, err := u.userRepo.GetUserByUserName(userName)

	if nil != err {
		return nil, err
	}

	view := u.userPresenter.Present(u_data)
	return view, nil
}

func (u *userAPICase) Save(newUser entity.User) (interface{}, error) {
	u_data, err := u.userRepo.Save(newUser)

	if nil != err {
		return nil, err
	}

	view := u.userPresenter.Present(u_data)
	return view, nil
}

func (u *userAPICase) Update(newUser entity.User) (interface{}, error) {
	u_data, err := u.userRepo.Update(newUser)

	if nil != err {
		return nil, err
	}

	view := u.userPresenter.Present(u_data)
	return view, nil
}

func (u *userAPICase) Delete(id string) error {
	return u.userRepo.Delete(id)
}

// func (u *userCase) Decode(body io.ReadCloser) entity.User {
// 	//Decode
// 	newUser := entity.User{}
// 	decoder := json.NewDecoder(body)
// 	decoder.Decode(&newUser)

// 	return newUser
// }
