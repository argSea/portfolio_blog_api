package repo

import (
	"context"

	"github.com/firelink/api.jsmith-develop.com/argSea/entity"
	"github.com/firelink/api.jsmith-develop.com/argSea/helper"
	"github.com/firelink/api.jsmith-develop.com/argSea/structure/argStore"
)

//Concrete for repo
type projectRepo struct {
	store argStore.ArgDB
}

func NewProjectRepo(store argStore.ArgDB) entity.ProjectRepository {
	return &projectRepo{
		store: store,
	}
}

func (p *projectRepo) GetProjects(ctx context.Context, limit int64, offset int64, sort entity.ProjectSort) (*entity.Projects, int64, error) {
	projects := &entity.Projects{}

	count, err := p.store.GetAll(ctx, limit, offset, sort, projects)

	if nil != err {
		return nil, 0, err
	}

	return projects, count, nil
}

func (p *projectRepo) GetByProjectID(ctx context.Context, id string) (*entity.Project, error) {
	newProject := &entity.Project{}

	finalTag := helper.GetFieldTag(*newProject, "Id", "bson")
	err := p.store.Get(ctx, finalTag, id, newProject)

	return newProject, err
}

func (p *projectRepo) GetProjectsByUserID(ctx context.Context, userID string, limit int64, offset int64, sort entity.ProjectSort) (*entity.Projects, int64, error) {
	projects := &entity.Projects{}
	project := &entity.Project{}

	finalTag := helper.GetFieldTag(*project, "UserIDs", "bson")
	count, err := p.store.GetMany(ctx, finalTag, userID, limit, offset, sort, projects)

	if nil != err {
		return nil, 0, err
	}

	return projects, count, err
}

func (p *projectRepo) Save(ctx context.Context, newProject entity.Project) (*entity.Project, error) {
	newID, err := p.store.Write(ctx, newProject)

	if nil != err {
		return nil, err
	}

	createdProject, cErr := p.GetByProjectID(ctx, newID)

	if nil != err {
		return nil, cErr
	}

	return createdProject, nil

}

func (p *projectRepo) Update(ctx context.Context, projectUpdates entity.Project) (*entity.Project, error) {
	projectID := projectUpdates.Id
	projectUpdates.Id = ""

	updateErr := p.store.Update(ctx, projectID, projectUpdates)

	if nil != updateErr {
		return nil, updateErr
	}

	currUser, currErr := p.GetByProjectID(ctx, projectID)

	if nil != currErr {
		return nil, currErr
	}

	return currUser, nil
}

func (p *projectRepo) Delete(ctx context.Context, id string) error {
	return p.store.Delete(ctx, id)
}
