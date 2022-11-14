package repo

import (
	"context"

	"github.com/firelink/api.jsmith-develop.com/argSea/entity"
	"github.com/firelink/api.jsmith-develop.com/argSea/helper"
	"github.com/firelink/api.jsmith-develop.com/argSea/structure/argStore"
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

func (u *userRepo) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	newUser := &entity.User{}

	finalTag := helper.GetFieldTag(*newUser, "Id", "bson")
	err := u.store.Get(ctx, finalTag, id, newUser)

	return newUser, err
}

func (u *userRepo) GetUserByUserName(ctx context.Context, userName string) (*entity.User, error) {
	newUser := &entity.User{}

	finalTag := helper.GetFieldTag(*newUser, "UserName", "bson")
	err := u.store.Get(ctx, finalTag, userName, newUser)

	return newUser, err
}

func (u *userRepo) Save(ctx context.Context, newUser entity.User) (*entity.User, error) {
	newID, err := u.store.Write(ctx, newUser)

	if nil != err {
		return nil, err
	}

	createdUser, cErr := u.GetUserByID(ctx, newID)

	if nil != err {
		return nil, cErr
	}

	return createdUser, nil

}

func (u *userRepo) Update(ctx context.Context, userUpdates entity.User) (*entity.User, error) {
	userID := userUpdates.Id
	userUpdates.Id = ""

	updateErr := u.store.Update(ctx, userID, userUpdates)

	if nil != updateErr {
		return nil, updateErr
	}

	currUser, currErr := u.GetUserByID(ctx, userID)

	if nil != currErr {
		return nil, currErr
	}

	return currUser, nil
}

func (u *userRepo) Delete(ctx context.Context, id string) error {
	return u.store.Delete(ctx, id)
}
