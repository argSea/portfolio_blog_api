package argStore

import "context"

type ArgDB interface {
	Get(context.Context, string, interface{}, interface{}) error
	GetMany(context.Context, string, interface{}, int64, int64, interface{}, interface{}) (int64, error)
	GetAll(context.Context, int64, int64, interface{}, interface{}) (int64, error)
	Write(context.Context, interface{}) (string, error)
	Update(context.Context, string, interface{}) error
	Delete(context.Context, string) error
}
