package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"errors"
	"fmt"
)

type FloorService struct{}

// 查寻
func (f *FloorService) QueryFloor(limit int, offset int, params interface{}) (interface{}, int64, error) {
	fmt.Println("楼搜索参数为", params)
	var floorList []dorm.Floor
	var total int64
	db := global.Global_Db.Model(&dorm.Floor{}).Limit(limit).Offset(offset)
	if params == nil {
		fmt.Println("-----------获取楼层不带参数")
		err := global.Global_Db.Find(&floorList).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		// fmt.Println("获取楼层不带参数", floorList)
		return floorList, total, nil
	}
	fmt.Println("后面++++++++++")
	err := db.Where("floors_name=?", params).Or("floors_type = ?", params).Find(&floorList).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return floorList, total, nil
}

// 获取宿舍楼相关宿舍
func (f *FloorService) GetFloorDorm() ([]dorm.Floor, error) {
	var floors []dorm.Floor
	err := global.Global_Db.Model(&dorm.Floor{}).Preload("Dorms").Find(&floors).Error
	if err != nil {
		return nil, err
	}
	return floors, nil
}

// 更新宿舍
func (f *FloorService) UpdateFloor(_floor dorm.Floor) error {
	var floors dorm.Floor
	err := global.Global_Db.Where("floors_name = ?", _floor.FloorsName).First(&floors).Error
	if err == nil {
		if floors.Id == _floor.Id {
		} else {
			return fmt.Errorf("宿舍楼已存在")
		}
	}
	err = global.Global_Db.Model(&dorm.Floor{}).Where("id= ?", _floor.Id).Updates(_floor).Error
	if err != nil {
		return fmt.Errorf("更新失败")
	}
	return nil
}

// 删除宿舍
func (f *FloorService) DeleteFloor(_floors *[]dorm.Floor) error {
	for _, value := range *_floors {
		var floor dorm.Floor
		fmt.Println("刪除的数据为", value.Id)
		err := global.Global_Db.Where("id= ?", value.Id).First(&floor).Error
		fmt.Println(floor)
		if err != nil {
			// fmt.Println("错误为", err)
			return errors.New("删除的数据不存在:" + value.FloorsName)
		}
	}

	err := global.Global_Db.Delete(*_floors).Error
	if err != nil {
		// 处理错误
		// response.FailWithMessage("该数据不存在,无法删除:"+del.FloorsName, c)
		return errors.New("删除数据失败")
	}
	return nil
}

// 添加宿舍楼
func (f *FloorService) CreateFloor(_floors *[]dorm.Floor) error {
	var temp dorm.Floor
	fmt.Println("进入宿舍楼业务")
	for _, v := range *_floors {
		err := global.Global_Db.Model(&dorm.Floor{}).Where("floors_name", v.FloorsName).First(&temp).Error
		if err == nil {
			return errors.New("该楼：" + v.FloorsName + "已存在")
		}
	}
	err := global.Global_Db.Create(_floors).Error
	if err != nil {
		// 处理错误
		// response.FailWithMessage("添加失败", c)
		return errors.New("添加失败")
	}
	return nil
}
