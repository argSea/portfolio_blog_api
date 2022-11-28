package argStore

type testStore struct {
}

func NewTestStore() ArgDB {
	return &testStore{}
}

func (t testStore) Get(field string, value interface{}, decoder interface{}) error {
	decoder = value

	return nil
}

func (t testStore) GetMany(field string, value interface{}, limit int64, offset int64, sort interface{}, decoder interface{}) (int64, error) {
	return 1, nil
}

func (t testStore) GetAll(limit int64, offset int64, sort interface{}, decoder interface{}) (int64, error) {
	return 1, nil
}

func (t testStore) Write(data interface{}) (string, error) {
	return "12345", nil
}

func (t testStore) Update(key string, newData interface{}) error {
	return nil
}

func (t testStore) Delete(key string) error {
	return nil
}
