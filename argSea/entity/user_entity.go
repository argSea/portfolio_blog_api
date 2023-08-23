package entity

type Users []User

//Entity // domain
type User struct {
	//Model
	Id            string          `json:"userID" bson:"_id,omitempty"`
	UserName      string          `json:"userName" bson:"userName,omitempty"`
	Password      password        `json:"password" bson:"password,omitempty"`
	FirstName     string          `json:"firstName" bson:"firstName,omitempty"`
	LastName      string          `json:"lastName" bson:"lastName,omitempty"`
	Email         string          `json:"email" bson:"email,omitempty"`
	Title         string          `json:"title" bson:"title,omitempty"`
	Picture       string          `json:"picture" bson:"picture,omitempty"`
	About         string          `json:"about" bson:"about,omitempty"`
	TechInterests []TechInterests `json:"techInterests" bson:"techInterests,omitempty"`
}

// func (u *User) GetID() string {
// 	return u.id
// }

// func (u *User) GetUserName() string {
// 	return u.userName
// }

// func (u *User) GetPassword() password {
// 	return u.password
// }

// func (u *User) GetFirstName() string {
// 	return u.firstName
// }

// func (u *User) GetLastName() string {
// 	return u.lastName
// }

// func (u *User) GetEmail() string {
// 	return u.email
// }

// func (u *User) GetTitle() string {
// 	return u.title
// }

// func (u *User) GetPicture() string {
// 	return u.picture
// }

// func (u *User) GetAbout() string {
// 	return u.about
// }

type password string

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}
