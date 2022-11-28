package repo

import (
	"github.com/argSea/portfolio_blog_api/argSea/core"
	"github.com/argSea/portfolio_blog_api/argSea/entity"
	"github.com/argSea/portfolio_blog_api/argSea/helper"
	"github.com/argSea/portfolio_blog_api/argSea/structure/argStore"
)

//Concrete for repo
type resumeRepo struct {
	store argStore.ArgDB
}

func NewResumeRepo(store argStore.ArgDB) core.ResumeRepository {
	return &resumeRepo{
		store: store,
	}
}

func (r *resumeRepo) GetResumeByID(id string) (*entity.Resume, error) {
	newResume := entity.Resume{}

	finalTag := helper.GetFieldTag(newResume, "Id", "bson")
	data, err := r.store.Get(finalTag, id, newResume)
	resume := data.(entity.Resume)

	return &resume, err
}

func (r *resumeRepo) GetResumeByUserID(userID string) (*entity.Resume, error) {
	newResume := entity.Resume{}

	finalTag := helper.GetFieldTag(newResume, "UserID", "bson")
	data, err := r.store.Get(finalTag, userID, newResume)
	resume := data.(entity.Resume)

	return &resume, err
}

func (r *resumeRepo) Save(newResume entity.Resume) (*entity.Resume, error) {
	newID, err := r.store.Write(newResume)

	if nil != err {
		return nil, err
	}

	createdResume, cErr := r.GetResumeByID(newID)

	if nil != err {
		return nil, cErr
	}

	return createdResume, nil

}

func (r *resumeRepo) Update(resumeUpdates entity.Resume) (*entity.Resume, error) {
	resumeID := resumeUpdates.Id
	resumeUpdates.Id = ""

	updateErr := r.store.Update(resumeID, resumeUpdates)

	if nil != updateErr {
		return nil, updateErr
	}

	currResume, currErr := r.GetResumeByID(resumeID)

	if nil != currErr {
		return nil, currErr
	}

	return currResume, nil
}

func (r *resumeRepo) Delete(id string) error {
	return r.store.Delete(id)
}
