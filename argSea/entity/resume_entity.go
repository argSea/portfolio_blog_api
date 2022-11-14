package entity

import "context"

type Resume struct {
	Id            string        `json:"projectID" bson:"_id,omitempty"`
	UserID        string        `json:"userID" bson:"userID,omitempty"`
	About         string        `json:"about" bson:"about,omitempty"`
	Experiences   Experiences   `json:"experiences" bson:"experiences,omitempty"`
	Education     *[]Education  `json:"education" bson:"education,omitempty"`
	ExtraCourses  Courses       `json:"extraCourses" bson:"extraCourses,omitempty"`
	SkillSections SkillSections `json:"skills" bson:"skills,omitempty"`
}

type ResumeRepository interface {
	GetResumeByID(context.Context, string) (*Resume, error)
	GetResumeByUserID(context.Context, string) (*Resume, error)
	Save(context.Context, Resume) (*Resume, error)
	Update(context.Context, Resume) (*Resume, error)
	Delete(context.Context, string) error
}

type ResumeUseCase interface {
	GetResumeByID(context.Context, string) (*Resume, error)
	GetResumeByUserID(context.Context, string) (*Resume, error)
	Save(context.Context, Resume) (*Resume, error)
	Update(context.Context, Resume) (*Resume, error)
	Delete(context.Context, string) error
}
