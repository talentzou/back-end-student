package system

type ServiceGroup struct {
	UserService
	MenuService
	BtnService
}

var SysService = new(ServiceGroup)
