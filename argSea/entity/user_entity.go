package entity

import "github.com/argSea/portfolio_blog_api/argSea/core"

type users []user

//Entity // domain
type user struct {
	//Model
	Id        string   `json:"userID" bson:"_id,omitempty"`
	UserName  string   `json:"userName" bson:"userName,omitempty"`
	Password  password `json:"password" bson:"password,omitempty"`
	FirstName string   `json:"firstName" bson:"firstName,omitempty"`
	LastName  string   `json:"lastName" bson:"lastName,omitempty"`
	Email     string   `json:"email" bson:"email,omitempty"`
	Title     string   `json:"title" bson:"title,omitempty"`
	Picture   string   `json:"picture" bson:"picture,omitempty"`
	About     string   `json:"about" bson:"about,omitempty"`
}

func NewUser() core.User {
	return &user{}
}

func (user user) GetUserID() string {
	return user.Id
}

func (user user) GetAbout() string {
	return user.About
}

func (user user) GetEmail() string {
	return user.Email
}

func (user user) GetUserName() string {
	return user.UserName
}

func (user user) GetFirstName() string {
	return user.FirstName
}

func (user user) GetLastName() string {
	return user.LastName
}

func (user user) GetTitle() string {
	return user.Title
}

func (user user) GetPicture() string {
	return user.Picture
}

type password string

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}
