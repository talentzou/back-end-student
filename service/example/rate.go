package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type RateService struct{}

// 查寻
func (f *RateService) QueryRate(limit int, offset int, condition interface{},dormId uint) (interface{}, int64, error) {
	var rateList []dorm.Rate
	var total int64

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
	fmt.Println("评分dormId++++++",dormId)
	if dormId != 0 {
	err := global.Global_Db.Model(&dorm.Rate{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Where("dorm_id=?", dormId).Count(&total).Limit(limit).Offset(offset).Find(&rateList).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
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
// 更新
func (f *RateService) UpdateRate(rate dorm.Rate) error {
	var tempRate dorm.Rate
	err := global.Global_Db.Debug().
	Not("id = ?", rate.Id).
	Where("dorm_id=? AND rate_date=? ", rate.DormId,rate.RateDate.Local()).
	First(&tempRate).Error
	// fmt.Println("找不到",err,tempRate)
	if err == nil {
		return errors.New(rate.RateDate.Local().Format("2006-01-02") + "数据已存在:无法更新")
	}

	err = global.Global_Db.Model(&rate).Where("id = ?", rate.Id).Updates(rate).Error
	if err != nil {
		// 处理错误
		return errors.New("更新rate失败")

	}
	return nil
}

// 添加
func (f *RateService) CreateRate(_rates *[]dorm.Rate) error {
	for _, v := range *_rates {
		// //查寻存在数据
		var tempArr dorm.Rate
		fmt.Println("添加的参数为",v)
		err := global.Global_Db.Where("dorm_id=? AND rate_date=? ", v.DormId,v.RateDate.Local()).
		First(&tempArr).Error
		if err == nil {
			return errors.New("数据"+tempArr.RateDate.Format("2006-01-02")+"已存在")
		}

		// for t := range tempArr {
		// 	// fmt.Println("添加",v.RateDate.Local() )
		// 	// fmt.Println("存在",tempArr[t].RateDate )
		// 	// fmt.Println("是否相等",v.RateDate.Local() == tempArr[t].RateDate)

		// 	if v.RateDate.Local() == tempArr[t].RateDate && v.DormId == tempArr[t].DormId {
		// 		return errors.New(tempArr[t].RateDate.Format("2006-01-02") + "的日期评分已存在")
		// 	}
		// }

	}
	// 添加数据
	err := global.Global_Db.Create(_rates).Error
	if err != nil {
		// 处理错误
		return errors.New("添加评分失败")
	}
	return nil
}

// 删除
func (f *RateService) DeleteRate(_rates *[]dorm.Rate) error {
	for _, value := range *_rates {
		err := global.Global_Db.Where("id=?", value.Id).First(&value).Error
		if err != nil {
			return errors.New("删除时间为:" + value.RateDate.Format("2006-01-02") + "宿舍数据不存在")
		}
	}
	err := global.Global_Db.Delete(_rates).Error
	if err != nil {
		// 处理错误
		return errors.New("宿舍数据删除失败")
	}
	return nil
}
