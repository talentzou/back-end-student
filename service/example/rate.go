package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"fmt"

	"gorm.io/gorm"
)

type RateService struct{}

// 查寻
func (f *RateService) QueryRate(limit int, offset int, condition interface{}) (interface{}, int64, error) {
	var rateList []dorm.Rate
	var total int64
	if condition == nil {
		db := global.Global_Db.Model(&dorm.Rate{}).Limit(limit).Offset(offset)
		err := db.Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Find(&rateList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分不带参数")
		return rateList, total, nil
	}
	// 查寻数据
	err := global.Global_Db.Model(&dorm.Rate{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Where(condition).Limit(limit).Offset(offset).Find(&rateList).Count(&total).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return rateList, total, nil
}
