package core

import "github.com/argSea/portfolio_blog_api/argHex/core"

type ResumeRepo interface {
	GetByUserID(id string) core.Resume
}
