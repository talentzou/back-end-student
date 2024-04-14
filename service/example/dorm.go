package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"errors"
	"fmt"

	// "gorm.io/gorm"
	// "fmt"
)

type DormService struct {
}

// 查寻
func (D *DormService) QueryDorm(limit int, offset int, condition map[string]interface{}) (interface{}, int64, error) {
	var dormList []dorm.Dorm
	var total int64
	if condition == nil {
		err := global.Global_Db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id").
			Preload("StudInfos").Find(&dormList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取宿舍带学生参数")
		return dormList, total, nil
	}

	fmt.Println("我来查询了")
	global.Global_Db.Limit(limit).Offset(offset)
	err := global.Global_Db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id").
		Find(&dormList).Count(&total).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return dormList, total, nil
}

// 更新
func (D *DormService) UpdateDorm(Dorm dorm.Dorm) error {
	var exist_dorm dorm.Dorm
	db := global.Global_Db.Model(&dorm.Dorm{})
	err := db.Where("floors_id=? AND dorm_number=?", Dorm.FloorId, Dorm.DormNumber).First(&exist_dorm)
	if err == nil {
		return errors.New("记录已存在")
	}
	err = global.Global_Db.Model(&dorm.Dorm{}).Where("id = ?", Dorm.Id).Updates(Dorm)
	if err.Error != nil {
		// 处理错误
		return errors.New("更新失败")
	}
	return nil
}

// 删除
func (D *DormService) DeleteDorm(dormList []dorm.Dorm) error {
	// 遍历查寻数据是否存在
	for _, value := range dormList {
		err := global.Global_Db.Where("id=?", value.Id).First(&value).Error
		if err != nil {
			return err
		}
	}
	err := global.Global_Db.Delete(&dormList).Error
	if err != nil {
		return err
	}

	return nil
}

// 创建
func (D *DormService) CreateDorm(dormList []dorm.Dorm) error {
	// 遍历查寻数据是否存在
	for _, v := range dormList {
		// 查询宿舍楼存在数据
		var tempFloor dorm.Floor
		err := global.Global_Db.Where("id=?", v.FloorId).First(&tempFloor).Error
		if err != nil {
			return errors.New("该宿舍楼" + v.FloorsName + "不存在,无法添加")
		}
		// 查询宿舍存在数据
		var tempArr dorm.Dorm
		err = global.Global_Db.Where("dorm_number=? AND floor_id=?", v.DormNumber, v.FloorId).First(&tempArr).Error
		if err != nil {
			continue
		} else {
			return errors.New("该宿舍:" + v.DormNumber + "已存在")
		}
	}
	// 添加数据
	err := global.Global_Db.Model(&dorm.Dorm{}).Create(&dormList).Error
	if err != nil {
		// 处理错误
		return errors.New("添加失败")
	}
	return nil

}
