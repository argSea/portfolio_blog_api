package usecase

import (
	"context"

	"github.com/firelink/api.jsmith-develop.com/argSea/entity"
)

//Concrete for use case
type projCase struct {
	projRepo entity.ProjectRepository
}

func NewProjectCase(repo entity.ProjectRepository) entity.ProjectUsecase {
	return &projCase{
		projRepo: repo,
	}
}

func (p *projCase) GetProjects(ctx context.Context, limit int64, offset int64, sort entity.ProjectSort) (*entity.Projects, int64, error) {
	return p.projRepo.GetProjects(ctx, limit, offset, sort)
}

func (p *projCase) GetByProjectID(ctx context.Context, id string) (*entity.Project, error) {
	return p.projRepo.GetByProjectID(ctx, id)
}

func (p *projCase) GetProjectsByUserID(ctx context.Context, userID string, limit int64, offset int64, sort entity.ProjectSort) (*entity.Projects, int64, error) {
	return p.projRepo.GetProjectsByUserID(ctx, userID, limit, offset, sort)
}

func (p *projCase) Save(ctx context.Context, newProject entity.Project) (*entity.Project, error) {
	return p.projRepo.Save(ctx, newProject)
}

func (p *projCase) Update(ctx context.Context, newProject entity.Project) (*entity.Project, error) {
	return p.projRepo.Update(ctx, newProject)
}

func (p *projCase) Delete(ctx context.Context, id string) error {
	return p.projRepo.Delete(ctx, id)
}
