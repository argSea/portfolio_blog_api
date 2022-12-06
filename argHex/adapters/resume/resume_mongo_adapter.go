package resumeAdapters

import core "github.com/argSea/portfolio_blog_api/argHex/core/resume"

type resumeMongoAdapter struct {
}

func NewUserMongoAdapter() core.ResumeRepo {
	r := resumeMongoAdapter{}

	return r
}

func (res resumeMongoAdapter) GetByUserID(id string) core.Resume {
	return core.Resume{}
}
