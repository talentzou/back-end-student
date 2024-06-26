package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"fmt"

	"gorm.io/gorm"
)

type StayService struct{}

// 查寻
func (f *StayService) QueryStay(limit int, offset int, condition interface{},dormId uint) (interface{}, int64, error) {
	var stayList []dorm.Stay
	var total int64
	fmt.Println("我是留宿申请+++++++++++++", condition)
	if condition == nil {
		db := global.Global_Db.Model(&dorm.Stay{}).Count(&total).Limit(limit).Offset(offset)
		err := db.Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Find(&stayList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分不带参数")
		return stayList, total, nil
	}
	//查寻数据
	fmt.Println("留宿dormId++++++",dormId)
	
	if dormId != 0 {
	err := global.Global_Db.Model(&dorm.Stay{}).Where(condition).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Where("dorm_id=?", dormId).Count(&total).Limit(limit).Offset(offset).Find(&stayList).Error
	fmt.Println("我是留宿申请---99---",err)
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return stayList, total, nil
	}

	// 查寻数据
	err := global.Global_Db.Model(&dorm.Stay{}).Where(condition).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Count(&total).Limit(limit).Offset(offset).Find(&stayList).Error
	fmt.Println("我是留宿申请---99---", err)
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return stayList, total, nil
}
