package usecase

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
)

//Concrete for use case
type resumeAPICase struct {
	resumeRepo      core.ResumeRepository
	resumePresenter core.ResumePresenter
}

func NewAPIResumeCase(repo core.ResumeRepository, pres core.ResumePresenter) core.ResumeUseCase {
	return &resumeAPICase{
		resumeRepo:      repo,
		resumePresenter: pres,
	}
}

func (r *resumeAPICase) GetResumeByID(id string) (interface{}, error) {
	r_data, err := r.resumeRepo.GetResumeByID(id)

	if nil != err {
		return nil, err
	}

	view := r.resumePresenter.Present(r_data)

	return view, nil
}

func (r *resumeAPICase) GetResumeByUserID(userName string) (interface{}, error) {
	r_data, err := r.resumeRepo.GetResumeByUserID(userName)

	if nil != err {
		return nil, err
	}

	view := r.resumePresenter.Present(r_data)

	return view, nil
}

func (r *resumeAPICase) Save(newResume entity.Resume) (interface{}, error) {
	return r.resumeRepo.Save(newResume)
}

func (r *resumeAPICase) Update(newResume entity.Resume) (interface{}, error) {
	return r.resumeRepo.Update(newResume)
}

func (r *resumeAPICase) Delete(id string) error {
	return r.resumeRepo.Delete(id)
}
