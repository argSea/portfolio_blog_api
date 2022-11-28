package core

import "github.com/argSea/portfolio_blog_api/argSea/entity"

type ResumeRepository interface {
	GetResumeByID(string) (*entity.Resume, error)
	GetResumeByUserID(string) (*entity.Resume, error)
	Save(entity.Resume) (*entity.Resume, error)
	Update(entity.Resume) (*entity.Resume, error)
	Delete(string) error
}

type ResumeUseCase interface {
	GetResumeByID(string) (*entity.Resume, error)
	GetResumeByUserID(string) (*entity.Resume, error)
	Save(entity.Resume) (*entity.Resume, error)
	Update(entity.Resume) (*entity.Resume, error)
	Delete(string) error
}
