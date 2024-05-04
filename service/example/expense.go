package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ExpenseService struct{}

// 查寻
func (f *ExpenseService) QueryExpense(limit int, offset int, condition interface{}, dormId uint) (interface{}, int64, error) {
	var stayList []dorm.Expense
	var total int64
	fmt.Println("我是水电+++++++++++++", condition)
	if condition != nil {
		condition := condition.([]string)
		fmt.Println("进来参数ggg", condition)
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

		db := global.Global_Db.Model(&dorm.Expense{}).
			Where("dorm_id=?", Dorm.Id).
			Preload("Dorm", func(db *gorm.DB) *gorm.DB {
				return db.Model(&dorm.Dorm{}).Debug().
					Select("dorm.*,floor.floors_name AS floors_name").
					Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
			}).
			Count(&total).Limit(limit).Offset(offset)

		//.Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		err = db.Find(&stayList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分不带参数")
		return stayList, total, nil
	}

	fmt.Println("水电费dormId++++++", dormId)
	if dormId != 0 {
		// 查寻数据
		err := global.Global_Db.Model(&dorm.Expense{}).
			Preload("Dorm", func(db *gorm.DB) *gorm.DB {
				return db.Model(&dorm.Dorm{}).Debug().
					Select("dorm.*,floor.floors_name AS floors_name").
					Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
			}).Where("dorm_id=?", dormId).Count(&total).Limit(limit).Offset(offset).Find(&stayList).Error
		if err != nil {
			// 处理错误
			return nil, 0, err
		}
		return stayList, total, nil
	}
	// 查寻数据
	err := global.Global_Db.Model(&dorm.Expense{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Where(condition).Count(&total).Limit(limit).Offset(offset).Find(&stayList).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return stayList, total, nil
}

// 创建
func (f *ExpenseService) CreateExpense(expenseList *[]dorm.Expense) error {

	for _, v := range *expenseList {
		var tempDorm dorm.Dorm
		err := global.Global_Db.Where("id=? ", v.DormId).First(&tempDorm).Error
		if err != nil {
			return errors.New("宿舍不存在,无法添加")
		}
		var tempRate dorm.Expense
		err = global.Global_Db.Debug().
			Where("dorm_id=? AND payment_time=? ", v.DormId, v.PaymentTime.Local()).
			First(&tempRate).Error
		fmt.Println("找不到", err, tempRate)
		if err == nil {
			return errors.New(tempRate.PaymentTime.Format("2006-01-02") + "数据已存在:无法更新")
		}
	}
	// 添加数据
	err := global.Global_Db.Create(&expenseList).Error
	if err != nil {
		// 处理错误
		return errors.New("添加失败")
	}
	return nil
}

// 删除
func (f *ExpenseService) DeleteExpense(expenseList *[]dorm.Expense) error {
	for _, value := range *expenseList {
		err := global.Global_Db.Where("id=?", value.Id).First(&value).Error
		if err != nil {
			return errors.New(value.PaymentTime.Format("2006-01-02") + "费用数据不存在")
		}
	}
	err := global.Global_Db.Delete(expenseList).Error
	if err != nil {
		return errors.New("删除数据失败")
	}
	return nil
}

// 更新
func (f *ExpenseService) UpdateFloor(expense dorm.Expense) error {
	var temp dorm.Expense
	err := global.Global_Db.Model(&temp).Where("id=?", expense.Id).First(&temp).Error
	if err != nil {
		return errors.New(expense.PaymentTime.Format("2006-01-02") + "费用数据不存在")
	}
	var tempRate dorm.Expense
	err = global.Global_Db.Debug().
		Not("id = ?", expense.Id).
		Where("dorm_id=? AND payment_time=? ", expense.DormId, expense.PaymentTime.Local()).
		First(&tempRate).Error
	fmt.Println("找不到", err, tempRate)
	if err == nil {
		return errors.New(temp.PaymentTime.Format("2006-01-02") + "数据已存在:无法更新")
	}

	err = global.Global_Db.Model(&expense).Where("id = ?", expense.Id).Updates(expense).Error
	if err != nil {
		// 处理错误
		return errors.New("更新失败")

	}
	return nil
}
