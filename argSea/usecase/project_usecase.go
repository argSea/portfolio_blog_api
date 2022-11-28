package usecase

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
)

//Concrete for use case
type projCase struct {
	projRepo core.ProjectRepository
}

func NewProjectCase(repo core.ProjectRepository) core.ProjectUsecase {
	return &projCase{
		projRepo: repo,
	}
}

func (p *projCase) GetProjects(limit int64, offset int64, sort entity.ProjectSort) (*entity.Projects, int64, error) {
	return p.projRepo.GetProjects(limit, offset, sort)
}

func (p *projCase) GetByProjectID(id string) (*entity.Project, error) {
	return p.projRepo.GetByProjectID(id)
}

func (p *projCase) GetProjectsByUserID(userID string, limit int64, offset int64, sort entity.ProjectSort) (*entity.Projects, int64, error) {
	return p.projRepo.GetProjectsByUserID(userID, limit, offset, sort)
}

func (p *projCase) Save(newProject entity.Project) (*entity.Project, error) {
	return p.projRepo.Save(newProject)
}

func (p *projCase) Update(newProject entity.Project) (*entity.Project, error) {
	return p.projRepo.Update(newProject)
}

func (p *projCase) Delete(id string) error {
	return p.projRepo.Delete(id)
}
