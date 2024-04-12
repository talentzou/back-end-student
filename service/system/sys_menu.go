package system

import (
	"back-end/global"
	"back-end/model/system"
	"fmt"
	"time"
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

// 获取所有菜单树
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

// 添加角色菜单关联
func (userService *MenuService) AddRelateMenu(roleId int, MenuIdList []int) error {
	fmt.Println("删除的角色为-----", roleId)
	fmt.Println("添加的角色菜单的参数+++++", MenuIdList)
	var role_menu []system.RoleMenus
	err := global.Global_Db.Where("role_id=?", roleId).Find(&role_menu).Error
	if err != nil {
		return err
	}
	var menus []system.RoleMenus
	for _, m := range MenuIdList {
		isFound := false
		for _, v := range role_menu {
			if v.MenuId == m {
				isFound = true
				break
			}
		}
		if !isFound {
			menus = append(menus, system.RoleMenus{RoleId: roleId, MenuId: m, CreatedAt: time.Now()})
		}
	}
	err = global.Global_Db.Create(&menus).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除角色菜单关联
func (userService *MenuService) DeleteRelateMenu(roleId int, MenuIdList []int) error {
	var role_menu []system.RoleMenus
	var menusIdList []int
	fmt.Println("删除的角色菜单为------", roleId)
	fmt.Println("添加的角色菜单的参数+++++", MenuIdList, len(MenuIdList))
	if len(MenuIdList) == 0 {
		err := global.Global_Db.Where("role_id=?", roleId).Delete(&system.RoleMenus{}).Error
		if err != nil {
			return err
		}
		return nil
	}
	err := global.Global_Db.Where("role_id=?", roleId).Find(&role_menu).Error
	if err != nil {
		return err
	}
	for _, r := range role_menu {
		isFound := false
		for _, v := range MenuIdList {
			if r.MenuId == v {
				isFound = true
				break
			}
		}
		if !isFound {
			menusIdList = append(menusIdList, r.MenuId)
		}
	}
	err = global.Global_Db.Where("menu_id IN ?", menusIdList).Delete(&system.RoleMenus{}).Error

	if err != nil {
		return err
	}
	return nil
}
