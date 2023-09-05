package service

import (
	"log"

	"github.com/argSea/portfolio_blog_api/argHex/domain"
	"github.com/argSea/portfolio_blog_api/argHex/in_port"
	"github.com/argSea/portfolio_blog_api/argHex/out_port"
	"github.com/argSea/portfolio_blog_api/argHex/utility"
	"golang.org/x/crypto/bcrypt"
)

type userCRUDService struct {
	repo out_port.UserRepo
}

func NewUserCRUDService(repo out_port.UserRepo) in_port.UserCRUDService {
	return userCRUDService{
		repo: repo,
	}
}

func (u userCRUDService) Create(user domain.User) (string, error) {
	user_id, err := u.repo.Add(user)

	if nil == err {
		log.Printf("User created with ID: %v\n", user_id)
	} else {
		log.Printf("User not created. err: %v", err)
	}

	return user_id, err
}

func (u userCRUDService) Read(id string) domain.User {
	userI := u.repo.Get(id)

	return userI
}

func (u userCRUDService) Update(user domain.User) error {
	err := u.repo.Set(user)

	// compare passwords with bcrypt
	check_pass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))

	if check_pass != nil {
		// hash password
		new_pass, err := utility.HashPassword(string(user.Password))

		if nil != err {
			log.Printf("Error hashing password: %v\n", err)
			return err
		}

		user.Password = domain.Password(new_pass)
	} else {
		user.Password = ""
	}

	if nil == err {
		log.Printf("User updated, user: %v\n", user)
	} else {
		log.Printf("User not updated, error: %v\n", err)
	}

	return err
}

func (u userCRUDService) Delete(user domain.User) error {
	err := u.repo.Remove(user)

	if nil == err {
		log.Printf("User with ID %v deleted successfully\n", user.Id)
	} else {
		log.Printf("User with ID %v could not be deleted, possible user doesn't exist? err: %v\n", user.Id, err)
	}

	return err
}
