package out_port

type AuthRepo interface {
	Store(token string, data interface{}) error
	Get(id string) string
	Remove(id string) error
}
