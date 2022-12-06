package command

// DATA OBJECT TO PASS THROUGH
type Request struct {
	test string
}

func (r *Request) Request(test string) {
	r.test = test
}

func (r *Request) getTest() string {
	return r.test
}
