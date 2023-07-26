package domain

type Projects []Project

type ProjectSort struct {
	Id       int `bson:"_id,omitempty"`
	Priority int `bson:"priority,omitempty"`
}

//Entity // domain
type Project struct {
	//Model
	Id                string      `json:"projectID" bson:"_id,omitempty"`
	UserIDs           []string    `json:"userIDs" bson:"userIDs,omitempty"`
	Type              string      `json:"projectType" bson:"projectType,omitempty"`
	Name              string      `json:"name" bson:"name,omitempty"`
	ShortName         *string     `json:"shortName" bson:"shortName,omitempty"`
	Icon              *string     `json:"icon" bson:"icon,omitempty"`
	Slug              string      `json:"slug" bson:"slug,omitempty"`
	RepoURL           *string     `json:"repoURL" bson:"repoURL,omitempty"`
	Description       *string     `json:"description" bson:"description,omitempty"`
	Skills            *[]string   `json:"skills" bson:"skills,omitempty"`
	Roles             *[]string   `json:"roles" bson:"roles,omitempty"`
	Priority          int         `json:"priority" bson:"priority,omitempty"`
	IsActive          bool        `json:"isActive" bson:"isActive,omitempty"`
	IsReleased        bool        `json:"isReleased" bson:"isReleased,omitempty"`
	BookID            *string     `json:"bookID" bson:"bookID,omitempty"`
	RelatedCourse     *Course     `json:"relatedCourse" bson:"relatedCourse,omitempty"`
	RelatedExperience *Experience `json:"relatedExperience" bson:"relatedExperience,omitempty"`
	Links             Links       `json:"links" bson:"links,omitempty"`
	Snippets          Snippets    `json:"snippets" bson:"snippets,omitempty"`
	Features          Features    `json:"features" bson:"features,omitempty"`
}
