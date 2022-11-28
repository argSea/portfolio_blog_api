package repo

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
	"github.com/argSea/portfolio_blog_api/argSea/helper"
	"github.com/argSea/portfolio_blog_api/argSea/structure/argStore"
)

//Concrete for repo
type projectRepo struct {
	store argStore.ArgDB
}

func NewProjectRepo(store argStore.ArgDB) core.ProjectRepository {
	return &projectRepo{
		store: store,
	}
}

func (p *projectRepo) GetProjects(limit int64, offset int64, sort entity.ProjectSort) (*entity.Projects, int64, error) {
	projects := &entity.Projects{}

	count, err := p.store.GetAll(limit, offset, sort, projects)

	if nil != err {
		return nil, 0, err
	}

	return projects, count, nil
}

func (p *projectRepo) GetByProjectID(id string) (*entity.Project, error) {
	newProject := &entity.Project{}

	finalTag := helper.GetFieldTag(*newProject, "Id", "bson")
	data, err := p.store.Get(finalTag, id, newProject)
	project := data.(entity.Project)

	return &project, err
}

func (p *projectRepo) GetProjectsByUserID(userID string, limit int64, offset int64, sort entity.ProjectSort) (*entity.Projects, int64, error) {
	projects := &entity.Projects{}
	project := &entity.Project{}

	finalTag := helper.GetFieldTag(*project, "UserIDs", "bson")
	count, err := p.store.GetMany(finalTag, userID, limit, offset, sort, projects)

	if nil != err {
		return nil, 0, err
	}

	return projects, count, err
}

func (p *projectRepo) Save(newProject entity.Project) (*entity.Project, error) {
	newID, err := p.store.Write(newProject)

	if nil != err {
		return nil, err
	}

	createdProject, cErr := p.GetByProjectID(newID)

	if nil != err {
		return nil, cErr
	}

	return createdProject, nil

}

func (p *projectRepo) Update(projectUpdates entity.Project) (*entity.Project, error) {
	projectID := projectUpdates.Id
	projectUpdates.Id = ""

	updateErr := p.store.Update(projectID, projectUpdates)

	if nil != updateErr {
		return nil, updateErr
	}

	currUser, currErr := p.GetByProjectID(projectID)

	if nil != currErr {
		return nil, currErr
	}

	return currUser, nil
}

func (p *projectRepo) Delete(id string) error {
	return p.store.Delete(id)
}
