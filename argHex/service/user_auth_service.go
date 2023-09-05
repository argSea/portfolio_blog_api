package service

import (
	"errors"
	"log"

	"github.com/argSea/portfolio_blog_api/argHex/domain"
	"github.com/argSea/portfolio_blog_api/argHex/in_port"
	"github.com/argSea/portfolio_blog_api/argHex/out_port"
	"golang.org/x/crypto/bcrypt"
)

type userAuthService struct {
	repo out_port.UserRepo
}

func NewUserAuthService(repo out_port.UserRepo) in_port.UserAuthService {
	return userAuthService{
		repo: repo,
	}
}

func (u userAuthService) Login(user domain.User) (string, error) {
	// get user from repo
	logged_in_user := u.repo.GetByUserName(user.UserName)

	// get salt from password
	logged_in := bcrypt.CompareHashAndPassword([]byte(logged_in_user.Password), []byte(user.Password))

	if logged_in == nil {
		log.Printf("User logged in with ID: %v\n", logged_in_user.Id)
		return user.Id, nil
	}

	log.Printf("User not logged in. err: %v", logged_in.Error())

	err := errors.New("Incorrect credentials or user does not exist")

	log.Printf("User not logged in. err: %v", err)
	return "", err
}

func (u userAuthService) Signup(user domain.User) (string, error) {
	user_id, err := u.repo.Add(user)

	if nil == err {
		log.Printf("User created with ID: %v\n", user_id)
	} else {
		log.Printf("User not created. err: %v", err)
	}

	return user_id, err
}

func (u userAuthService) hashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if nil != err {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}

	return string(bcryptPassword), nil
}

// Path: go/argHex/service/user_crud_service.go
