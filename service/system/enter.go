package system

type ServiceGroup struct {
	UserService
	MenuService
}

var SysService = new(ServiceGroup)
