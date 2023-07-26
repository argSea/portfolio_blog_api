package in_port

import "github.com/argSea/portfolio_blog_api/argHex/domain"

type ProjectCRUDService interface {
	Create(project domain.Project) (string, error)
	Read(id string) domain.Project
	Update(project domain.Project) error
	Delete(project domain.Project) error
}
