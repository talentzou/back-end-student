package system

import (
	// "back-end/config"
	// "back-end/global"
	// sysReq "back-end/model/common/request"
	sysReq "back-end/model/common/request"
	sysRes "back-end/model/common/response"
	"back-end/model/system"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	// "strconv"
)

type menuApi struct{}

var MenuApi menuApi

// // 根据路由配置文件获取菜单
// func GetAsyncMenu(c *gin.Context) {
// 	id, b := c.Params.Get("authorityId")
// 	if !b {
// 		sysRes.FailWithMessage("缺少角色参数标识", c)
// 		return
// 	}
// 	authority, err := strconv.Atoi(id)
// 	if err != nil {
// 		sysRes.FailWithMessage("参数标识不是Number类型", c)
// 		return
// 	}
// 	fmt.Println("authority", authority)
// 	var routes []config.Route
// 	if authority == 1 {
// 		routes = global.Global_Web_Route.Admin
// 		sysRes.OkWithDetailed(sysReq.SysMenusResponse{
// 			Authority: authority,
// 			Menu:      routes,
// 		}, "获取管理员路由成功", c)
// 		return
// 	} else if authority == 2 {
// 		routes = global.Global_Web_Route.Dorm
// 		sysRes.OkWithDetailed(sysReq.SysMenusResponse{
// 			Authority: authority,
// 			Menu:      routes,
// 		}, "获取宿舍管路由成功", c)
// 		return
// 	} else {
// 		routes = global.Global_Web_Route.Student
// 		sysRes.OkWithDetailed(sysReq.SysMenusResponse{
// 			Authority: authority,
// 			Menu:      routes,
// 		}, "获取学生路由成功", c)
// 		return
// 	}
// }

// 获取角色菜单
func (m *menuApi) GetSelfMenu(c *gin.Context) {
	fmt.Println("菜单角色id为", utils.GetUserRoleId(c))
	menus, err := menuService.GetMenuTree(int(utils.GetUserRoleId(c))) //1
	if err != nil {
		fmt.Println("出错了")
		sysRes.FailWithMessage("获取菜单失败", c)
		return
	}
	if menus == nil {
		menus = []system.MenuTree{}
	}
	sysRes.OkWithDetailed(sysRes.SysMenusResponse{Menus: menus}, "获取菜单成功", c)
}

// 获取菜单树
func (m *menuApi) GetAllMenu(c *gin.Context) {
	AllMenus, err := menuService.GetAllMenu() //1
	if err != nil {
		fmt.Println("出错了")
		sysRes.FailWithMessage("获取菜单失败", c)
		return
	}
	if AllMenus == nil {
		AllMenus = []system.MenuTree{}
	}
	sysRes.OkWithDetailed(sysRes.SysMenusResponse{Menus: AllMenus}, "获取菜单成功", c)
}

// 添加角色菜单关联
func (m *menuApi) AddRelateRoleAndMenu(c *gin.Context) {
	var role_menu sysReq.ReqRelateRoleAndMenu
	err := c.ShouldBindJSON(&role_menu)
	if err != nil {
		sysRes.FailWithMessage("参数错误", c)
		return
	}
	err =  menuService.AddRelateMenu(role_menu.RoleId, role_menu.RoleMenuIdList)
	if err != nil {
		sysRes.FailWithMessage("添加失败", c)
	}
	sysRes.OkWithMessage("添加角色菜单关联成功", c)

}

// 删除角色菜单关联
func (m *menuApi) DeleteRelateRoleAndMenu(c *gin.Context) {
	var role_menu sysReq.ReqRelateRoleAndMenu
	err := c.ShouldBindJSON(&role_menu)
	if err != nil {
		sysRes.FailWithMessage("参数错误", c)
		return
	}
	err = menuService.DeleteRelateMenu(role_menu.RoleId, role_menu.RoleMenuIdList)
	if err != nil {
		sysRes.FailWithMessage("删除失败", c)
	}
	sysRes.OkWithMessage("删除角色菜单关联成功", c)
}
