package system

import (
	"back-end/global"
	"fmt"
	"back-end/model/system"

	"github.com/gin-gonic/gin"
	// "fmt"
)

type MenuService struct{}

// 获取动态菜单树
func (userService *MenuService) GetMenuTree(authorityId uint, c *gin.Context) (menus []system.MenuTree, err error) {
	menuTree, err := userService.getMenuTreeMap(authorityId)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = userService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

// 查找树
func (userService *MenuService) getMenuTreeMap(authorityId uint) (treeMap map[uint][]system.MenuTree, err error) {
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

// 遍历
func (userService *MenuService) getChildrenList(menu *system.MenuTree, treeMap map[uint][]system.MenuTree) (err error) {
	menu.Children = treeMap[menu.Id]
	fmt.Println("hahaha", len(menu.Children), menu.Children)
	for i := 0; i < len(menu.Children); i++ {
		err = userService.getChildrenList(&menu.Children[i], treeMap)
	}

	return err
}
