package usecase

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
)

//Concrete for use case
type resumeCase struct {
	resumeRepo core.ResumeRepository
}

func NewResumeCase(repo core.ResumeRepository) core.ResumeUseCase {
	return &resumeCase{
		resumeRepo: repo,
	}
}

func (r *resumeCase) GetResumeByID(id string) (*entity.Resume, error) {
	return r.resumeRepo.GetResumeByID(id)
}

func (r *resumeCase) GetResumeByUserID(userName string) (*entity.Resume, error) {
	return r.resumeRepo.GetResumeByUserID(userName)
}

func (r *resumeCase) Save(newResume entity.Resume) (*entity.Resume, error) {
	return r.resumeRepo.Save(newResume)
}

func (r *resumeCase) Update(newResume entity.Resume) (*entity.Resume, error) {
	return r.resumeRepo.Update(newResume)
}

func (r *resumeCase) Delete(id string) error {
	return r.resumeRepo.Delete(id)
}
