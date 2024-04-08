package main

import (
	"back-end/global"
	"back-end/model/common/response"
	"back-end/model/system"
	"back-end/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 获取菜单
func GetMenu(c *gin.Context) {
	menus, err := GetMenuTree(utils.GetUserAuthorityId(c))
	if err != nil {
		// global.GVA_LOG.Error("获取失败!", zap.Error(err))
		fmt.Println("出错了")
		response.FailWithMessage("获取失败", c)
		return
	}
	if menus == nil {
		menus = []system.MenuTree{}
	}
	response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
}

// 获取动态菜单树
func GetMenuTree(authorityId uint) (menus []system.MenuTree, err error) {
	menuTree, err := getMenuTreeMap(authorityId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

// 查找树
func getMenuTreeMap(authorityId uint) (treeMap map[uint][]system.MenuTree, err error) {
	// 查找角色菜单树
	var baseMenu []system.MenuTree
	treeMap = make(map[uint][]system.MenuTree)

	// 查找角色菜单信息
	var SysRoleMenus []system.RoleMenu
	err = global.Global_Db.Where("authority_id = ?", authorityId).Find(&SysRoleMenus).Error
	if err != nil {
		return
	}
	// 拿出角色信息菜单id
	var MenuIds []uint
	for i := range SysRoleMenus {
		MenuIds = append(MenuIds, SysRoleMenus[i].MenuId)
	}

	err = global.Global_Db.Where("id in (?)", MenuIds).Order("id").Find(&baseMenu).Error
	if err != nil {
		return
	}
	for _, v := range baseMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}
//遍历
func getChildrenList(menu *system.MenuTree, treeMap map[uint][]system.MenuTree) (err error) {
	menu.Children = treeMap[menu.ParentId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
