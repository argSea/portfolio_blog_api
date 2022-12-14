package entity

type Users []User

//Entity // domain
type User struct {
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

type password string

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}
