package core

type ResumeRepo interface {
	GetByUserID(id string) Resume
}
