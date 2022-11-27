package core

type User interface {
	GetUserID() string
	GetUserName() string
	GetFirstName() string
	GetLastName() string
	GetEmail() string
	GetTitle() string
	GetPicture() string
	GetAbout() string
}

//User repo interface
type UserRepository interface {
	GetUserByID(string) (*User, error)
	GetUserByUserName(string) (*User, error)
	Save(User) (*User, error)
	Update(User) (*User, error)
	Delete(string) error
}

//Use case for the above
type UserUsecase interface {
	GetUserByID(string) (*User, error)
	GetUserByUserName(string) (*User, error)
	Save(User) (*User, error)
	Update(User) (*User, error)
	Delete(string) error
	// Decode(io.ReadCloser) User
}

type UserPresenter interface {
	Present() *User
}

type password string

func (password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}
