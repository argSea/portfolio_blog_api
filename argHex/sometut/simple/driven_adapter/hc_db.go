package driven_adapter

//CONCRETE IMPLEMENTATION OF DB
type HCDB struct {
}

func (h HCDB) GetUser(test string) interface{} {
	user := struct {
		UserID string
		Name   string
	}{
		UserID: "12345",
		Name:   "Test name",
	}

	return user
}
