package system

type ServiceGroup struct {
	UserService
}

var SysService = new(ServiceGroup)
