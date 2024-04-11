package system

import (
	"back-end/global"
	"back-end/model/system"
	"fmt"

	// "gorm.io/gorm"
)

type RoleService struct{}

func (r *RoleService) GetRoleList() ([]system.Role, error) {
	var allRole []system.Role
	err := global.Global_Db.Model(&system.Role{}).Association("MenuTrees").Error
	if err != nil {
		fmt.Println("关联失败菜单")
	}
	// err = global.Global_Db.Model(&system.MenuTree{}).Association("SysBtns").Error
	// if err != nil {
	// 	fmt.Println("关联失败菜单")
	// }.Preload("SysBtns")
	db := global.Global_Db.Preload("MenuTrees").Model(&system.Role{}).Limit(6).Order("id")

	err = db.Not("id=?", 1).Find(&allRole).Error
	if err != nil {
		return nil, err
	}
	return allRole, nil
}
/* 
, func(db *gorm.DB) *gorm.DB {
		return db.Joins("")
	}

*/