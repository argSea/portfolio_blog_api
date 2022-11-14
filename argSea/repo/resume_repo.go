package repo

import (
	"context"

	"github.com/firelink/api.jsmith-develop.com/argSea/entity"
	"github.com/firelink/api.jsmith-develop.com/argSea/helper"
	"github.com/firelink/api.jsmith-develop.com/argSea/structure/argStore"
)

//Concrete for repo
type resumeRepo struct {
	store argStore.ArgDB
}

func NewResumeRepo(store argStore.ArgDB) entity.ResumeRepository {
	return &resumeRepo{
		store: store,
	}
}

func (r *resumeRepo) GetResumeByID(ctx context.Context, id string) (*entity.Resume, error) {
	newResume := &entity.Resume{}

	finalTag := helper.GetFieldTag(*newResume, "Id", "bson")
	err := r.store.Get(ctx, finalTag, id, newResume)

	return newResume, err
}

func (r *resumeRepo) GetResumeByUserID(ctx context.Context, userID string) (*entity.Resume, error) {
	newResume := &entity.Resume{}

	finalTag := helper.GetFieldTag(*newResume, "UserID", "bson")
	err := r.store.Get(ctx, finalTag, userID, newResume)

	return newResume, err
}

func (r *resumeRepo) Save(ctx context.Context, newResume entity.Resume) (*entity.Resume, error) {
	newID, err := r.store.Write(ctx, newResume)

	if nil != err {
		return nil, err
	}

	createdResume, cErr := r.GetResumeByID(ctx, newID)

	if nil != err {
		return nil, cErr
	}

	return createdResume, nil

}

func (r *resumeRepo) Update(ctx context.Context, resumeUpdates entity.Resume) (*entity.Resume, error) {
	resumeID := resumeUpdates.Id
	resumeUpdates.Id = ""

	updateErr := r.store.Update(ctx, resumeID, resumeUpdates)

	if nil != updateErr {
		return nil, updateErr
	}

	currResume, currErr := r.GetResumeByID(ctx, resumeID)

	if nil != currErr {
		return nil, currErr
	}

	return currResume, nil
}

func (r *resumeRepo) Delete(ctx context.Context, id string) error {
	return r.store.Delete(ctx, id)
}
