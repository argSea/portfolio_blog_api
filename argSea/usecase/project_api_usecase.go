package usecase

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
)

//Concrete for use case
type projAPICase struct {
	projRepo      core.ProjectRepository
	projPresenter core.ProjectPresenter
}

func NewAPIProjectCase(repo core.ProjectRepository, pres core.ProjectPresenter) core.ProjectUsecase {
	return &projAPICase{
		projRepo:      repo,
		projPresenter: pres,
	}
}

func (p *projAPICase) GetProjects(limit int64, offset int64, sort entity.ProjectSort) (interface{}, error) {
	p_data, count, err := p.projRepo.GetProjects(limit, offset, sort)

	if nil != err {
		return nil, err
	}

	count = count + 1
	view := p.projPresenter.Present(p_data)

	return view, nil
}

func (p *projAPICase) GetByProjectID(id string) (interface{}, error) {
	return p.projRepo.GetByProjectID(id)
}

func (p *projAPICase) GetProjectsByUserID(userID string, limit int64, offset int64, sort entity.ProjectSort) (interface{}, error) {
	p_data, count, err := p.projRepo.GetProjectsByUserID(userID, limit, offset, sort)

	if nil != err {
		return nil, err
	}

	count = count + 1
	view := p.projPresenter.Present(p_data)

	return view, nil
}

func (p *projAPICase) Save(newProject entity.Project) (interface{}, error) {
	return p.projRepo.Save(newProject)
}

func (p *projAPICase) Update(newProject entity.Project) (interface{}, error) {
	return p.projRepo.Update(newProject)
}

func (p *projAPICase) Delete(id string) error {
	return p.projRepo.Delete(id)
}
