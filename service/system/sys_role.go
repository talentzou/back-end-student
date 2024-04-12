package system

import (
	"back-end/global"
	"back-end/model/system"
	"fmt"
	// "gorm.io/gorm"
)

type RoleService struct{}

// 获取角色
func (r *RoleService) GetRoleList() ([]system.Role, error) {
	var allRole []system.Role
	err := global.Global_Db.Model(&system.Role{}).Association("MenuTrees").Error
	if err != nil {
		fmt.Println("关联失败菜单")
	}
	err = global.Global_Db.Model(&system.Role{}).Association("SysBtns").Error
	if err != nil {
		fmt.Println("关联失败菜单")
	}

	// db := global.Global_Db.Preload("MenuTrees",func(db *gorm.DB) *gorm.DB {
	// 	return db.Preload("SysBtns")
	// }).Preload("SysBtns").Limit(6).Order("id")

	db := global.Global_Db.Preload("MenuTrees").Preload("SysBtns").Limit(5).Order("id")

	err = db.Not("id=?", 1).Find(&allRole).Error
	if err != nil {
		return nil, err
	}
	return allRole, nil
}

// 添加角色
func (r *RoleService) CreateRolesList(role system.Role) error {
	var allRole []system.Role
	db := global.Global_Db.Model(&system.Role{})
	err := db.Find(&allRole).Error
	if err != nil {
		return err
	}
	for _, v := range allRole {
		if v.RoleName == role.RoleName {
			return fmt.Errorf("该角色已存在")
		}
	}
	err = db.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除角色
func (r *RoleService) DeleteRolesList(role system.Role) error {
	var Role system.Role
	db := global.Global_Db.Model(&system.Role{})
	err := db.Where("id=?", role.Model.ID).First(&Role).Error
	if err != nil {
		return fmt.Errorf("该角色不存在,无法删除")
	}
	fmt.Println("删除到这99999999999")
	err = db.Where("id=?", role.Model.ID).Delete(&Role).Error
	if err != nil {
		return err
	}
	return nil
}
