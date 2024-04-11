package system

import (
	"back-end/global"
	"back-end/model/system"
)

type BtnService struct{}

// 获取角色按钮
func (b *BtnService) GetBtnTreeMap(RoleId uint) ([]system.SysBtn, error) {
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

// 获取角色按钮id列表
func (b *BtnService) GETBtnIdList(RoleId uint) ([]int, error) {
	var btnIdList []int
	var RoleBtns []system.RoleBtns
	err := global.Global_Db.Model(&system.RoleBtns{}).Where("role_id", RoleId).Find(&RoleBtns).Error
	if err != nil {
		return nil, err
	}
	for i := range RoleBtns {
		btnIdList = append(btnIdList, RoleBtns[i].SysBtnId)
	}
	return btnIdList, nil
}
