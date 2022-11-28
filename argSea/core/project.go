package core

import "github.com/argSea/portfolio_blog_api/argSea/entity"

//User repo interface
type ProjectRepository interface {
	GetProjects(int64, int64, entity.ProjectSort) (*entity.Projects, int64, error)
	GetByProjectID(string) (*entity.Project, error)
	GetProjectsByUserID(string, int64, int64, entity.ProjectSort) (*entity.Projects, int64, error)
	Save(entity.Project) (*entity.Project, error)
	Update(entity.Project) (*entity.Project, error)
	Delete(string) error
}

//Use case for the above
type ProjectUsecase interface {
	GetProjects(int64, int64, entity.ProjectSort) (*entity.Projects, int64, error)
	GetByProjectID(string) (*entity.Project, error)
	GetProjectsByUserID(string, int64, int64, entity.ProjectSort) (*entity.Projects, int64, error)
	Save(entity.Project) (*entity.Project, error)
	Update(entity.Project) (*entity.Project, error)
	Delete(string) error
}
