package userService

import "github.com/argSea/portfolio_blog_api/argHex/core"

type userCRUDService struct {
	repo core.UserRepo
}

type userOutput struct {
	Id        string `json:"userID" bson:"_id,omitempty"`
	UserName  string `json:"userName" bson:"userName,omitempty"`
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	Title     string `json:"title" bson:"title,omitempty"`
	Picture   string `json:"picture" bson:"picture,omitempty"`
	About     string `json:"about" bson:"about,omitempty"`
}

func NewUserCRUDService(repo core.UserRepo) core.UserCRUDService {
	return userCRUDService{
		repo: repo,
	}
}

func (u userCRUDService) Create(user core.User) error {
	return nil
}

func (u userCRUDService) Read(id string) interface{} {
	userI := u.repo.GetUserByID(id)
	userO := userOutput{}
	userO.Id = userI.GetID()
	userO.About = userI.GetAbout()
	userO.Email = userI.GetEmail()
	userO.FirstName = userI.GetFirstName()
	userO.LastName = userI.GetLastName()
	userO.Picture = userI.GetPicture()
	userO.Title = userI.GetTitle()
	userO.UserName = userI.GetUserName()

	return userO
}

func (u userCRUDService) Update(user core.User) error {
	return nil
}

func (u userCRUDService) Delete(user core.User) error {
	return nil
}
