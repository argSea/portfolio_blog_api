package core

type Resume struct {
	id            string        //`json:"projectID" bson:"_id,omitempty"`
	userID        string        //`json:"userID" bson:"userID,omitempty"`
	about         string        //`json:"about" bson:"about,omitempty"`
	experiences   Experiences   //`json:"experiences" bson:"experiences,omitempty"`
	education     Education     //`json:"education" bson:"education,omitempty"`
	extraCourses  Courses       //`json:"extraCourses" bson:"extraCourses,omitempty"`
	skillSections SkillSections //`json:"skills" bson:"skills,omitempty"`
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

func (res *Resume) GetExperiences() Experiences {
	return res.experiences
}

func (res *Resume) GetEducation() Education {
	return res.education
}

func (res *Resume) GetExtraCourses() Courses {
	return res.extraCourses
}

func (res *Resume) GetSkillSections() SkillSections {
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

func (res *Resume) AddExperience(experience Experience) *Resume {
	res.experiences = append(res.experiences, experience)

	return res
}

func (res *Resume) SetExperiences(experiences Experiences) *Resume {
	res.experiences = experiences

	return res
}

func (res *Resume) AddEducation(ce CollegeExperience) *Resume {
	res.education = append(res.education, ce)

	return res
}

func (res *Resume) SetEducation(education Education) *Resume {
	res.education = education

	return res
}

func (res *Resume) AddExtraCourse(course Course) *Resume {
	res.extraCourses = append(res.extraCourses, course)

	return res
}

func (res *Resume) SetExtraCourses(courses Courses) *Resume {
	res.extraCourses = courses

	return res
}

func (res *Resume) AddSkillSection(skillSection SkillSection) *Resume {
	res.skillSections = append(res.skillSections, skillSection)

	return res
}

func (res *Resume) SetSkillSection(skillSections SkillSections) *Resume {
	res.skillSections = skillSections

	return res
}
