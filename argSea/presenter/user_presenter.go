package presenter

import "github.com/firelink/api.jsmith-develop.com/argSea/entity"

type userPresenter struct {
	Status  string       `json:"status"`
	Code    int          `json:"code"`
	Message string       `json:"message,omitempty"`
	Users   entity.Users `json:"users"`
}

func NewUserPresenter() *userPresenter {
	return &userPresenter{}
}

func (u *userPresenter) SetStatus(status string) {
	u.Status = status
}

func (u *userPresenter) SetCode(code int) {
	u.Code = code
}

func (u *userPresenter) SetMessage(message string) {
	u.Message = message
}

func (u *userPresenter) AddUser(user entity.User) {
	u.Users = append(u.Users, user)
}
