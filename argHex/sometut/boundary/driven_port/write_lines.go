package driven_port

//BOUNDARY CROSSER FROM INSIDE APP TO OUTPUT VIEW
//TRIGGERED BY APP
type WriteLines interface {
	WriteLines(obj interface{})
}
