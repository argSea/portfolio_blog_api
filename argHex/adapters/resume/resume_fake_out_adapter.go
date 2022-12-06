package resumeAdapters

import (
	"fmt"

	"github.com/argSea/portfolio_blog_api/argHex/core/core_shared_structs"
	core "github.com/argSea/portfolio_blog_api/argHex/core/resume"
)

type resumeFakeOutAdapter struct {
}

func NewResumeFakeOutAdapter() core.ResumeRepo {
	return resumeFakeOutAdapter{}
}

func (u resumeFakeOutAdapter) GetByUserID(id string) core.Resume {
	resume := core.Resume{}
	resume.SetUserID("cabbage").
		AddEducation(core_shared_structs.CollegeExperience{}).
		AddExperience(core_shared_structs.Experience{}).
		SetAbout("I'm me").
		AddExtraCourse(core_shared_structs.Course{}).
		AddSkillSection(core_shared_structs.SkillSection{})

	fmt.Println(resume)

	return resume
}
