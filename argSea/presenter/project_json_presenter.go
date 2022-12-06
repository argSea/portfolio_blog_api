package presenter

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
)

type projectJSONPresenter struct {
	status   string
	code     int
	message  string
	projects []interface{}
}

func NewProjectPresenter() core.UserPresenter {
	up := projectJSONPresenter{
		status: "ok",
		code:   200,
	}

	return &up
}

func (p *projectJSONPresenter) Present(model interface{}) interface{} {
	presented := struct {
		Status   string        `json:"status"`
		Code     int           `json:"code"`
		Message  string        `json:"message,omitempty"`
		Projects []interface{} `json:"resumes"`
	}{
		Status:   p.status,
		Code:     p.code,
		Message:  p.message,
		Projects: p.projects,
	}

	presented.Projects = append(presented.Projects, model)

	return presented
}
