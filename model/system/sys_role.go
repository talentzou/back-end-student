package system

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	SysUserId uint       `json:"userId"` //用户表外键
	RoleName  string     `json:"roleName" gorm:"comment:角色名称"`
	MenuTrees []MenuTree `json:"Menu" gorm:"many2many:role_menus;joinForeignKey:RoleId;joinReferences:MenuId"`
	SysBtns   []SysBtn   `json:"sysBtns" gorm:"many2many:role_btns;joinForeignKey:RoleId;joinReferences:MenuId"`
	Btn_List  []int      `json:"list" gorm:"-"`
}

//joinForeignKey:RoleId;joinReferences:MenuId
//;joinForeignKey:RoleId;joinReferences:BtnId"
