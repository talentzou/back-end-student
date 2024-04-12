package system

import (
	"back-end/global"
	"back-end/model/system"
	"fmt"
	// "fmt"
)

type BtnService struct{}

// 获取角色按钮
func (b *BtnService) GetBtnRoleTree(RoleId uint) ([]system.SysBtn, error) {
	var RoleBtns []system.RoleBtns
	err := global.Global_Db.Model(&system.RoleBtns{}).Where("role_id=?", RoleId).Find(&RoleBtns).Error
	if err != nil {
		return nil, err
	}
	var btns []int
	for i := range RoleBtns {
		btns = append(btns, RoleBtns[i].SysBtnId)
	}
	var allBtn []system.SysBtn
	err = global.Global_Db.Model(&system.SysBtn{}).Where("id IN ? ", btns).Find(&allBtn).Error
	if err != nil {
		return nil, err
	}
	return allBtn, nil
}

// 按钮菜单树
func (b *BtnService) GetBtnTreeMap() ([]system.MenuTree, error) {
	var allMenu []system.MenuTree
	// err := global.Global_Db.Association("SysBtns").Error
	// if err != nil {
	// 	fmt.Println("关联失败")
	// }
	err := global.Global_Db.Preload("SysBtns").Find(&allMenu).Error
	if err != nil {
		return nil, err
	}
	return allMenu, nil
}

// 删除角色按钮关联
func (b *BtnService) DeleteRelateBtn(roleId int, SysBtnIdList []int) error {
	var role_btn []system.RoleBtns
	var btnIdList []int
	fmt.Println("删除的角色为+++++", roleId)
	fmt.Println("添加的角色按钮的参数+++++", SysBtnIdList)
	// err := global.Global_Db.Where("role_id=?", roleId).Delete(&system.RoleBtns{}).Error
	if len(SysBtnIdList) == 0 {
		err := global.Global_Db.Where("role_id=?", roleId).Delete(&system.RoleBtns{}).Error
		if err != nil {
			return err
		}
		return nil
	}
	err := global.Global_Db.Where("role_id=?", roleId).Find(&role_btn).Error
	if err != nil {
		return err
	}
	for _, r := range role_btn {
		isFound := false
		for _, v := range btnIdList {
			if r.SysBtnId == v {
				isFound = true
				break
			}
		}
		if !isFound {
			btnIdList = append(btnIdList, r.SysBtnId)
		}
	}
	err = global.Global_Db.Where("sys_btn_id IN ?", btnIdList).Delete(&system.RoleBtns{}).Error

	if err != nil {
		return err
	}
	return nil
}

// 添加角色按钮关联
func (b *BtnService) AddRelateBtn(roleId int, SysBtnIdList []int) error {
	var role_btn []system.RoleBtns
	fmt.Println("添加的角色为+++++", roleId)
	fmt.Println("添加的角色按钮的参数+++++", SysBtnIdList)
	err := global.Global_Db.Where("role_id=?", roleId).Find(&role_btn).Error
	if err != nil {
		return err
	}
	var btns []system.RoleBtns
	for _, s := range SysBtnIdList {
		isFound := false
		for _, v := range role_btn {
			if v.SysBtnId == s {
				isFound = true
				break
			}
		}
		if !isFound {
			btns = append(btns, system.RoleBtns{RoleId: roleId, SysBtnId: s})
		}

	}

	err = global.Global_Db.Create(&btns).Error
	if err != nil {
		return err
	}
	return nil
}
