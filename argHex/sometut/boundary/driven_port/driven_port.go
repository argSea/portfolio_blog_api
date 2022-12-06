package driven_port

//BOUNDARY CROSSER FROM INSIDE APP TO DB OR STORE
//TRIGGERED BY APP
type DrivenPort interface {
	GetUser(userID string) interface{}
}
