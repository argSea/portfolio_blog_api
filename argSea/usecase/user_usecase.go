package usecase

import (
	"context"

	"github.com/firelink/api.jsmith-develop.com/argSea/entity"
)

//Concrete for use case
type userCase struct {
	userRepo entity.UserRepository
}

func NewUserCase(repo entity.UserRepository) entity.UserUsecase {
	return &userCase{
		userRepo: repo,
	}
}

func (u *userCase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return u.userRepo.GetUserByID(ctx, id)
}

func (u *userCase) GetUserByUserName(ctx context.Context, userName string) (*entity.User, error) {
	return u.userRepo.GetUserByUserName(ctx, userName)
}

func (u *userCase) Save(ctx context.Context, newUser entity.User) (*entity.User, error) {
	return u.userRepo.Save(ctx, newUser)
}

func (u *userCase) Update(ctx context.Context, newUser entity.User) (*entity.User, error) {
	return u.userRepo.Update(ctx, newUser)
}

func (u *userCase) Delete(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}

// func (u *userCase) Decode(body io.ReadCloser) entity.User {
// 	//Decode
// 	newUser := entity.User{}
// 	decoder := json.NewDecoder(body)
// 	decoder.Decode(&newUser)

// 	return newUser
// }
