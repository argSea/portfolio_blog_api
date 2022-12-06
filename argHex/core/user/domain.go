package core

type Users []User

//Entity // domain
type User struct {
	//Model
	id        string
	userName  string
	password  password
	firstName string
	lastName  string
	email     string
	title     string
	picture   string
	about     string
}

func (u *User) SetID(id string) *User {
	u.id = id

	return u
}

func (u *User) SetUserName(userName string) *User {
	u.userName = userName

	return u
}

func (u *User) SetFirstName(name string) *User {
	u.firstName = name

	return u
}

func (u *User) SetLastName(name string) *User {
	u.lastName = name

	return u
}

func (u *User) SetEmail(email string) *User {
	u.email = email

	return u
}

func (u *User) SetTitle(title string) *User {
	u.title = title

	return u
}

func (u *User) SetPicture(pic string) *User {
	u.picture = pic

	return u
}

func (u *User) SetAbout(ab string) *User {
	u.about = ab

	return u
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetUserName() string {
	return u.userName
}

func (u *User) GetPassword() password {
	return u.password
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetTitle() string {
	return u.title
}

func (u *User) GetPicture() string {
	return u.picture
}

func (u *User) GetAbout() string {
	return u.about
}

type password string

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}
