package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"fmt"

	"gorm.io/gorm"
)

type ExpenseService struct{}

// 查寻
func (f *ExpenseService) QueryExpense(limit int, offset int, condition interface{}) (interface{}, int64, error) {
	var stayList []dorm.Expense
	var total int64
	fmt.Println("我是水电+++++++++++++", condition)
	if condition == nil {
		db := global.Global_Db.Model(&dorm.Expense{}).Limit(limit).Offset(offset)
		err := db.Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Find(&stayList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分不带参数")
		return stayList, total, nil
	}
	fmt.Println("我是水电费---99---")
	// 查寻数据
	err := global.Global_Db.Model(&dorm.Expense{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Where(condition).Limit(limit).Offset(offset).Find(&stayList).Count(&total).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return stayList, total, nil
}
