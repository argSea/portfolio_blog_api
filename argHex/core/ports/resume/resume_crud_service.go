package core

import "github.com/argSea/portfolio_blog_api/argHex/core"

//User service for CRUD
type ResumeCRUDService interface {
	Create(res core.Resume) error
	Read(id string) interface{}
	Update(res core.Resume) error
	Delete(res core.Resume) error
}
