package argStore

type ArgDB interface {
	Get(field string, value interface{}, decoder interface{}) (interface{}, error)
	GetMany(field string, value interface{}, limit int64, offset int64, sort interface{}, decoder interface{}) (int64, error)
	GetAll(limit int64, offset int64, sort interface{}, decoder interface{}) (int64, error)
	Write(data interface{}) (string, error)
	Update(key string, newData interface{}) error
	Delete(key string) error
}
