package driver_port

//BOUNDARY CROSSER FROM CONTROLLER TO INSIDE APP
//CORRESPONDS TO USE CASE
type DriverPort interface {
	React(command interface{})
}
