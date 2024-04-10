package system

import (
	"back-end/global"
	"back-end/model/system"
	// "fmt"
)

type MenuService struct{}

// 获取动态菜单树
func (userService *MenuService) GetMenuTree(RoleId int) (menus []system.MenuTree, err error) {
	menuTree, err := userService.getMenuTreeMap(RoleId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = userService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

// 查找树
func (userService *MenuService) getMenuTreeMap(roleId int) (treeMap map[uint][]system.MenuTree, err error) {
	// 查找角色菜单树
	var baseMenu []system.MenuTree
	treeMap = make(map[uint][]system.MenuTree)

	// 查找角色菜单信息
	var SysRoleMenus []system.RoleMenus
	err = global.Global_Db.Where("role_id = ?",roleId).Find(&SysRoleMenus).Error
	if err != nil {
		return
	}
	// 拿出角色信息菜单id
	var MenuIds []int
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

// 遍历
func (userService *MenuService) getChildrenList(menu *system.MenuTree, treeMap map[uint][]system.MenuTree) (err error) {
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		err = userService.getChildrenList(&menu.Children[i], treeMap)
	}

	return err
}
