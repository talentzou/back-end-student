package system

type ServiceGroup struct {
	UserService
	MenuService
	BtnService
	RoleService
}

var SysService = new(ServiceGroup)
