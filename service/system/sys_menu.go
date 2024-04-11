package system

import (
	"back-end/global"
	"back-end/model/system"
	// "fmt"
)

type MenuService struct{}

// 获取角色菜单树
func (userService *MenuService) GetMenuTree(RoleId int) (menus []system.MenuTree, err error) {
	menuTree, err := userService.getMenuTreeMap(RoleId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = userService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

// 菜单树map
func (userService *MenuService) getMenuTreeMap(roleId int) (treeMap map[uint][]system.MenuTree, err error) {
	// 查找角色菜单树
	var baseMenu []system.MenuTree
	treeMap = make(map[uint][]system.MenuTree)

	// 查找角色菜单信息
	var SysRoleMenus []system.RoleMenus
	err = global.Global_Db.Where("role_id = ?", roleId).Find(&SysRoleMenus).Error
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

// 添加子菜单
func (userService *MenuService) getChildrenList(menu *system.MenuTree, treeMap map[uint][]system.MenuTree) (err error) {
	menu.Children = treeMap[menu.Id]
	for i := 0; i < len(menu.Children); i++ {
		err = userService.getChildrenList(&menu.Children[i], treeMap)
	}

	return err
}


// 获取所有菜单
func (userService *MenuService) GetAllMenu() ([]system.MenuTree, error) {
	var allMenu []system.MenuTree
	treeMap := make(map[uint][]system.MenuTree)
	err := global.Global_Db.Order("id").Find(&allMenu).Error
	if err != nil {
		return nil, err
	}
	for _, v := range allMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	menus := treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = userService.getChildrenList(&menus[i], treeMap)
	}
	return menus, err
}


// // 获取角色菜单id列表
// func (userService *MenuService) GETMenuIdList(RoleId uint) ([]int, error) {
// 	var menuIdList []int
// 	var RoleMenus []system.RoleMenus
// 	err := global.Global_Db.Model(&system.RoleMenus{}).Where("role_id", RoleId).Find(&RoleMenus).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	for i := range RoleMenus {
// 		menuIdList = append(menuIdList, RoleMenus[i].MenuId)
// 	}
// 	return menuIdList, nil
// }

