package driver_adapter

import "github.com/argSea/portfolio_blog_api/argHex/sometut/boundary/driver_port"

//CONTROLLER
type fakeUser struct {
	driverPort driver_port.DriverPort
}

func NewFakeUser(port driver_port.DriverPort) *fakeUser {
	fu := fakeUser{
		driverPort: port,
	}

	return &fu
}

func (f *fakeUser) Run() {
	f.driverPort.React("test")
	f.driverPort.React("worse test")
}
