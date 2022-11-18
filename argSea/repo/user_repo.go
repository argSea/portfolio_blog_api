package repo

import (
	"context"
	"time"

	"github.com/argSea/portfolio_blog_api/argSea/entity"
	"github.com/argSea/portfolio_blog_api/argSea/helper"
	"github.com/argSea/portfolio_blog_api/argSea/structure/argStore"
)

//Concrete for repo
type userRepo struct {
	store argStore.ArgDB
}

func NewUserRepo(store argStore.ArgDB) entity.UserRepository {
	return &userRepo{
		store: store,
	}
}

func (u *userRepo) GetUserByID(id string) (*entity.User, error) {
	newUser := &entity.User{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	finalTag := helper.GetFieldTag(*newUser, "Id", "bson")
	err := u.store.Get(ctx, finalTag, id, newUser)

	return newUser, err
}

func (u *userRepo) GetUserByUserName(userName string) (*entity.User, error) {
	newUser := &entity.User{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	finalTag := helper.GetFieldTag(*newUser, "UserName", "bson")
	err := u.store.Get(ctx, finalTag, userName, newUser)

	return newUser, err
}

func (u *userRepo) Save(newUser entity.User) (*entity.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)
	newID, err := u.store.Write(ctx, newUser)

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
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)

	updateErr := u.store.Update(ctx, userID, userUpdates)

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
	ctx, _ := context.WithTimeout(context.Background(), time.Second+10)
	return u.store.Delete(ctx, id)
}
