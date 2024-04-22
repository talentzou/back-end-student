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

	if condition != nil {
		condition := condition.([]string)
		fmt.Println("进来参数ggg",condition)
		var floor dorm.Floor
		err := global.Global_Db.Model(&dorm.Floor{}).Where("floors_name=?", condition[0]).First(&floor).Error
		if err != nil {
			return nil, 0, err
		}
		var Dorm dorm.Dorm
		err = global.Global_Db.Model(&dorm.Dorm{}).Where("floor_id=? AND dorm_number=?", floor.Id, condition[1]).First(&Dorm).Error
		if err != nil {
			return nil, 0, err
		}

		db := global.Global_Db.Model(&dorm.Rate{}).Where("dorm_id=?", Dorm.Id).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Count(&total).Limit(limit).Offset(offset)

		//.Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		err = db.Find(&rateList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分带参数++++++++++++")
		return rateList, total, nil
	}

	// 查寻数据
	err := global.Global_Db.Model(&dorm.Rate{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Count(&total).Limit(limit).Offset(offset).Find(&rateList).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return rateList, total, nil
}
