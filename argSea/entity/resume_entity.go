package entity

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
	GetResumeByID(string) (*Resume, error)
	GetResumeByUserID(string) (*Resume, error)
	Save(Resume) (*Resume, error)
	Update(Resume) (*Resume, error)
	Delete(string) error
}

type ResumeUseCase interface {
	GetResumeByID(string) (*Resume, error)
	GetResumeByUserID(string) (*Resume, error)
	Save(Resume) (*Resume, error)
	Update(Resume) (*Resume, error)
	Delete(string) error
}
