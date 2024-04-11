package system

import (
	service "back-end/service/system"
)

var (
	userService = service.SysService.UserService
	menuService = service.SysService.MenuService
	btnService  = service.SysService.BtnService
	roleService=service.SysService.RoleService
)
