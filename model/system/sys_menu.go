package system

import "time"

// 菜单树表
type MenuTree struct {
	Id       uint       `json:"id" gorm:"primarykey"`
	ParentId uint       `json:"parent_id"`
	Name     string     `json:"name" gorm:"size:256"`
	Title    string     `json:"title" gorm:"size:256"`
	Children []MenuTree `json:"children" gorm:"-"`
	Roles    []Role     `json:"roles" gorm:"many2many:role_menus;joinForeignKey:MenuId;joinReferences:RoleId"`
	SysBtns  []SysBtn   `json:"btns" gorm:"foreignKey:MenuId"`
}

//joinForeignKey:MenuId;joinReferences:RoleId
// 角色菜单关联信息表
type RoleMenus struct {
	MenuId    int `gorm:"primarykey"` //菜单id
	RoleId    int `gorm:"primarykey"` //角色id
	CreatedAt time.Time
}

//菜单
/*
1.foreignKey:MenuId;references:Id
2.foreignKey:Id;associationForeignKey:MenuId
3.RoleMenus []RoleMenu `json:"-" gorm:"many2many:role_authority_menus;foreignKey:Id;associationForeignKey:MenuId"`
*/

//角色菜单
/*
// AuthorityId uint `json:"authorityId" gorm:"comment:用户角色Id"`
	// RoleName         string            `json:"roleName" gorm:"comment:角色名称"`
	// MenuTrees        []MenuTree        `json:"Menu" gorm:"many2many:role_authority_menus;foreignKey:Id;associationForeignKey:Id"`


*/
//foreignKey:MenuId;associationForeignKey:Id
// 菜单树表ForeignKey:Id;references:MenuId
