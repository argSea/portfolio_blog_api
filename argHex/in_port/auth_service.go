package in_port

import "github.com/argSea/portfolio_blog_api/argHex/data_objects"

const (
	// permissions
	PERM_USER  = "user"
	PERM_ADMIN = "admin"
)

// user auth interface
type AuthService interface {
	Generate(id string) (string, error)
	Validate(token string) (data_objects.AuthValidationResponseObject, error)
	IsAuthorized(id string, token string, roles ...string) bool
}
