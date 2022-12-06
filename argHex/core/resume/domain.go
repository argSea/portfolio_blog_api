package core

import extra "github.com/argSea/portfolio_blog_api/argHex/core/extra/"

type Resume struct {
	Id            string            //`json:"projectID" bson:"_id,omitempty"`
	UserID        string            //`json:"userID" bson:"userID,omitempty"`
	About         string            //`json:"about" bson:"about,omitempty"`
	Experiences   extra.Experiences //`json:"experiences" bson:"experiences,omitempty"`
	Education     *[]Education      //`json:"education" bson:"education,omitempty"`
	ExtraCourses  Courses           //`json:"extraCourses" bson:"extraCourses,omitempty"`
	SkillSections SkillSections     //`json:"skills" bson:"skills,omitempty"`
}
