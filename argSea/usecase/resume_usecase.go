package usecase

import (
	"context"

	"github.com/firelink/api.jsmith-develop.com/argSea/entity"
)

//Concrete for use case
type resumeCase struct {
	resumeRepo entity.ResumeRepository
}

func NewResumeCase(repo entity.ResumeRepository) entity.ResumeUseCase {
	return &resumeCase{
		resumeRepo: repo,
	}
}

func (r *resumeCase) GetResumeByID(ctx context.Context, id string) (*entity.Resume, error) {
	return r.resumeRepo.GetResumeByID(ctx, id)
}

func (r *resumeCase) GetResumeByUserID(ctx context.Context, userName string) (*entity.Resume, error) {
	return r.resumeRepo.GetResumeByUserID(ctx, userName)
}

func (r *resumeCase) Save(ctx context.Context, newResume entity.Resume) (*entity.Resume, error) {
	return r.resumeRepo.Save(ctx, newResume)
}

func (r *resumeCase) Update(ctx context.Context, newResume entity.Resume) (*entity.Resume, error) {
	return r.resumeRepo.Update(ctx, newResume)
}

func (r *resumeCase) Delete(ctx context.Context, id string) error {
	return r.resumeRepo.Delete(ctx, id)
}
