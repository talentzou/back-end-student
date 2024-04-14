package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"back-end/model/test/repair"
	"fmt"

	"gorm.io/gorm"
)
type RepairService struct{}

// 查寻
func (f *RepairService) QueryRepair(limit int, offset int, condition interface{}) (interface{}, int64, error) {
	var stayList []repair.Repair
	var total int64
	fmt.Println("我是维修+++++++++++++", condition)
	if condition == nil {
		db := global.Global_Db.Model(&repair.Repair{}).Limit(limit).Offset(offset)
		err := db.Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Find(&stayList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分不带参数")
		return stayList, total, nil
	}
	fmt.Println("我是维修---99---")
	// 查寻数据
	err := global.Global_Db.Model(&repair.Repair{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Where(condition).Limit(limit).Offset(offset).Find(&stayList).Count(&total).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return stayList, total, nil
}
