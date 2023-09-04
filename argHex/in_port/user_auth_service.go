package in_port

import "github.com/argSea/portfolio_blog_api/argHex/domain"

// user auth interface
type UserAuthService interface {
	Login(user domain.User) (string, error)
	Signup(user domain.User) (string, error)
}
