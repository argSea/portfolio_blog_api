package argStore

type Blackgate struct {
	user   string
	pass   string
	dbName string
}

func NewBlackgate(user string, pass string, db string) *Blackgate {
	return &Blackgate{user: user, pass: pass, dbName: db}
}
