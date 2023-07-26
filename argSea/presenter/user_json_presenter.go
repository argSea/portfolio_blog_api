package presenter

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
)

type userJSONPresenter struct {
	status  string
	code    int
	message string
	users   []interface{}
}

func NewUserPresenter() core.UserPresenter {
	up := userJSONPresenter{
		status: "ok",
		code:   200,
	}

	return &up
}

func (p *userJSONPresenter) Present(model interface{}) interface{} {
	presented := struct {
		Status  string        `json:"status"`
		Code    int           `json:"code"`
		Message string        `json:"message,omitempty"`
		Users   []interface{} `json:"users"`
	}{
		Status:  p.status,
		Code:    p.code,
		Message: p.message,
		Users:   p.users,
	}

	presented.Users = append(presented.Users, model)

	return presented
}
