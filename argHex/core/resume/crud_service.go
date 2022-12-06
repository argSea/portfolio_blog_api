package core

//User service for CRUD
type ResumeCRUDService interface {
	Create(res Resume) error
	Read(id string) interface{}
	Update(res Resume) error
	Delete(res Resume) error
}
