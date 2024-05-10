package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type StayService struct{}

// 查寻
func (f *StayService) QueryStay(limit int, offset int, condition interface{}, dormId uint) (interface{}, int64, error) {
	var stayList []dorm.Stay
	var total int64
	fmt.Println("我是留宿申请+++++++++++++", condition)
	if condition == nil {
		db := global.Global_Db.Model(&dorm.Stay{}).Count(&total).Limit(limit).Offset(offset)
		err := db.Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().
				Select("dorm.*,floor.floors_name AS floors_name").
				Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Find(&stayList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分不带参数")
		return stayList, total, nil
	}
	//查寻数据
	fmt.Println("留宿dormId++++++", dormId)

	if dormId != 0 {
		err := global.Global_Db.Model(&dorm.Stay{}).
			Where(condition).
			Preload("Dorm", func(db *gorm.DB) *gorm.DB {
				return db.Model(&dorm.Dorm{}).Debug().
					Select("dorm.*,floor.floors_name AS floors_name").
					Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
			}).
			Where("dorm_id=?", dormId).
			Count(&total).
			Limit(limit).Offset(offset).Find(&stayList).Error
		fmt.Println("我是留宿申请---99---", err)
		if err != nil {
			// 处理错误
			return nil, 0, err
		}
		return stayList, total, nil
	}

	// 查寻数据
	err := global.Global_Db.Model(&dorm.Stay{}).
		Where(condition).
		Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().
				Select("dorm.*,floor.floors_name AS floors_name").
				Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).
		Count(&total).
		Limit(limit).Offset(offset).
		Find(&stayList).Error
	fmt.Println("我是留宿申请---99---", err)
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return stayList, total, nil
}

// 添加
func (f *StayService) CreateStay(_stays *[]dorm.Stay) error {
	for _, v := range *_stays {
		var tempStay dorm.Stay
		err := global.Global_Db.Debug().
			// Where("dorm_id=? AND start_time=? AND end_time=? ", v.DormId, v.StayTime.StartTime.Local(), v.StayTime.EndTime.Local()).
			Where(&dorm.Stay{
				DormId: v.DormId,
				StayTime: dorm.StayTime{
					StartTime: v.StayTime.StartTime.Local(),
					EndTime:   v.StayTime.EndTime.Local(),
				},
				StudentName: v.StudentName,
			}).
			First(&tempStay).Error
		// fmt.Println("找不到", err, tempStay)
		if err == nil {
			return errors.New("学生:" + v.StudentName + "留宿申请已存在")
		}

	}
	err := global.Global_Db.Create(&_stays).Error
	if err != nil {
		// 处理错误
		return errors.New("添加失败")
	}

	return nil
}

// 删除
func (f *StayService) DeleteStay(_stays *[]dorm.Stay, roleId uint) error {
	// 遍历查寻数据是否存在
	fmt.Println("执行删除前")
	for _, value := range *_stays {
		fmt.Println("执行删除++++++")
		var tempStay dorm.Stay
		err := global.Global_Db.Where("id=?", value.Id).First(&tempStay).Error
		if err != nil {
			return errors.New("删除日期为:" + value.StayTime.StartTime.Format("2006-01-02") + "至" + value.StayTime.EndTime.Format("2006-01-02") + "的数据不存在:")
		}
		fmt.Println("状态意见",value.Opinions)
		if tempStay.Opinions == "不同意" || tempStay.Opinions == "同意" {
			if roleId > 2 {
				return errors.New("状态发生改变，权限不足，无法删除")
			}
		}
	}
	fmt.Println("执行删除-----")
	err := global.Global_Db.Delete(_stays).Error
	if err != nil {
		// 处理错误
		return errors.New("删除失败")
	}
	return nil
}

// 更新
func (f *StayService) UpdateStay(stay dorm.Stay, roleId uint) error {
	// 查寻数据是否存在
	var tempRate dorm.Stay
	err := global.Global_Db.Where("id=?", stay.Id).First(&tempRate).Error
	if err != nil {
		return errors.New(stay.StayCause + ":数据不存在:无法更新")
	}
	var tempStay dorm.Stay
	err = global.Global_Db.Debug().Not("id=?", stay.Id).
		Where(&dorm.Stay{
			DormId: stay.DormId,
			StayTime: dorm.StayTime{
				StartTime: stay.StayTime.StartTime.Local(),
				EndTime:   stay.StayTime.EndTime.Local(),
			},
			StudentName: stay.StudentName,
		}).
		First(&tempStay).Error

	// fmt.Println("找不到", err, tempStay)
	if err == nil {
		return errors.New("学生:" + stay.StudentName + "留宿申请已存在,无法更新")
	}

	if stay.Opinions == "不同意" || stay.Opinions == "同意" {
		if roleId > 2 {
			return errors.New("状态发生改变，权限不足，无法更新")
		}
	}

	err = global.Global_Db.Model(&stay).Where("id = ?", stay.Id).Updates(stay).Error
	if err != nil {
		// 处理错误
		return errors.New("更新失败")

	}
	return nil
}
