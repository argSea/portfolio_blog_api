package presenter

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
)

type resumeJSONPresenter struct {
	status  string
	code    int
	message string
	resumes []interface{}
}

func NewResumePresenter() core.UserPresenter {
	up := resumeJSONPresenter{
		status: "ok",
		code:   200,
	}

	return &up
}

func (r *resumeJSONPresenter) Present(model interface{}) interface{} {
	presented := struct {
		Status  string        `json:"status"`
		Code    int           `json:"code"`
		Message string        `json:"message,omitempty"`
		Resumes []interface{} `json:"resumes"`
	}{
		Status:  r.status,
		Code:    r.code,
		Message: r.message,
		Resumes: r.resumes,
	}

	presented.Resumes = append(presented.Resumes, model)

	return presented
}
