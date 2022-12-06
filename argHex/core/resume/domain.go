package core

import "github.com/argSea/portfolio_blog_api/argHex/core/core_shared_structs"

type Resume struct {
	id            string                            //`json:"projectID" bson:"_id,omitempty"`
	userID        string                            //`json:"userID" bson:"userID,omitempty"`
	about         string                            //`json:"about" bson:"about,omitempty"`
	experiences   core_shared_structs.Experiences   //`json:"experiences" bson:"experiences,omitempty"`
	education     core_shared_structs.Education     //`json:"education" bson:"education,omitempty"`
	extraCourses  core_shared_structs.Courses       //`json:"extraCourses" bson:"extraCourses,omitempty"`
	skillSections core_shared_structs.SkillSections //`json:"skills" bson:"skills,omitempty"`
}

func (res *Resume) GetID() string {
	return res.id
}

func (res *Resume) GetUserID() string {
	return res.userID
}

func (res *Resume) GetAbout() string {
	return res.about
}

func (res *Resume) GetExperiences() core_shared_structs.Experiences {
	return res.experiences
}

func (res *Resume) GetEducation() core_shared_structs.Education {
	return res.education
}

func (res *Resume) GetExtraCourses() core_shared_structs.Courses {
	return res.extraCourses
}

func (res *Resume) GetSkillSections() core_shared_structs.SkillSections {
	return res.skillSections
}

func (res *Resume) SetID(id string) *Resume {
	res.id = id

	return res
}

func (res *Resume) SetUserID(userID string) *Resume {
	res.userID = userID

	return res
}

func (res *Resume) SetAbout(about string) *Resume {
	res.about = about

	return res
}

func (res *Resume) AddExperience(experience core_shared_structs.Experience) *Resume {
	res.experiences = append(res.experiences, experience)

	return res
}

func (res *Resume) SetExperiences(experiences core_shared_structs.Experiences) *Resume {
	res.experiences = experiences

	return res
}

func (res *Resume) AddEducation(ce core_shared_structs.CollegeExperience) *Resume {
	res.education = append(res.education, ce)

	return res
}

func (res *Resume) SetEducation(education core_shared_structs.Education) *Resume {
	res.education = education

	return res
}

func (res *Resume) AddExtraCourse(course core_shared_structs.Course) *Resume {
	res.extraCourses = append(res.extraCourses, course)

	return res
}

func (res *Resume) SetExtraCourses(courses core_shared_structs.Courses) *Resume {
	res.extraCourses = courses

	return res
}

func (res *Resume) AddSkillSection(skillSection core_shared_structs.SkillSection) *Resume {
	res.skillSections = append(res.skillSections, skillSection)

	return res
}

func (res *Resume) SetSkillSection(skillSections core_shared_structs.SkillSections) *Resume {
	res.skillSections = skillSections

	return res
}
