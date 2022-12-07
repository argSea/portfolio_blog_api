package resumeAdapters

import (
	"fmt"

	"github.com/argSea/portfolio_blog_api/argHex/core"
)

type resumeFakeOutAdapter struct {
}

func NewResumeFakeOutAdapter() core.ResumeRepo {
	return resumeFakeOutAdapter{}
}

func (u resumeFakeOutAdapter) GetByUserID(id string) core.Resume {
	resume := core.Resume{}
	resume.SetUserID("cabbage").
		AddEducation(core.CollegeExperience{}).
		AddExperience(core.Experience{}).
		SetAbout("I'm me").
		AddExtraCourse(core.Course{}).
		AddSkillSection(core.SkillSection{})

	fmt.Println(resume)

	return resume
}
