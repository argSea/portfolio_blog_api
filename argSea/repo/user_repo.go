package repo

import (
	"fmt"

	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
	"github.com/argSea/portfolio_blog_api/argSea/helper"
	"github.com/argSea/portfolio_blog_api/argSea/structure/argStore"
)

//Concrete for repo
type userRepo struct {
	store argStore.ArgDB
}

func NewUserRepo(store argStore.ArgDB) core.UserRepository {
	return &userRepo{
		store: store,
	}
}

func (u *userRepo) GetUserByID(id string) (*entity.User, error) {
	newUser := entity.User{}

	finalTag := helper.GetFieldTag(newUser, "Id", "bson")
	data, err := u.store.Get(finalTag, id, newUser)
	user := data.(entity.User)
	fmt.Println(user)

	return &user, err
}

func (u *userRepo) GetUserByUserName(userName string) (*entity.User, error) {
	newUser := entity.User{}

	finalTag := helper.GetFieldTag(newUser, "UserName", "bson")
	data, err := u.store.Get(finalTag, userName, newUser)
	user := data.(entity.User)

	return &user, err
}

func (u *userRepo) Save(newUser entity.User) (*entity.User, error) {
	newID, err := u.store.Write(newUser)

	if nil != err {
		return nil, err
	}

	createdUser, cErr := u.GetUserByID(newID)

	if nil != err {
		return nil, cErr
	}

	return createdUser, nil

}

func (u *userRepo) Update(userUpdates entity.User) (*entity.User, error) {
	userID := userUpdates.Id
	userUpdates.Id = ""

	updateErr := u.store.Update(userID, userUpdates)

	if nil != updateErr {
		return nil, updateErr
	}

	currUser, currErr := u.GetUserByID(userID)

	if nil != currErr {
		return nil, currErr
	}

	return currUser, nil
}

func (u *userRepo) Delete(id string) error {
	return u.store.Delete(id)
}
