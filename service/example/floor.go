package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
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
		err := global.Global_Db.Find(&floorList).Error
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
		return nil,err
	}
	return floors,nil
}
