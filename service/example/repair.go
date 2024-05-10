package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"back-end/model/test/repair"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RepairService struct{}

// 查寻
func (f *RepairService) QueryRepair(limit int, offset int, condition []string, dormId uint) (interface{}, int64, error) {
	var repairList []repair.Repair
	var total int64
	fmt.Println("我是维修+++++++++++++", condition)
	if condition != nil {
		var floor dorm.Floor
		err := global.Global_Db.Model(&dorm.Floor{}).Where("floors_name=?", condition[0]).First(&floor).Error
		if err != nil {
			return nil, 0, err
		}
		// fmt.Println("宿舍楼为+++++++++",floor)
		var Dorm dorm.Dorm
		err = global.Global_Db.Model(&dorm.Dorm{}).Where("floor_id=? AND dorm_number=?", floor.Id, condition[1]).First(&Dorm).Error
		if err != nil {
			return nil, 0, err
		}
		// fmt.Println("宿舍为+++++++++",Dorm)
		db := global.Global_Db.Model(&repair.Repair{}).
			Preload("Dorm", func(db *gorm.DB) *gorm.DB {
				return db.Model(&dorm.Dorm{}).Debug().
					Select("dorm.*,floor.floors_name AS floors_name").
					Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
			}).
			Limit(limit).Offset(offset)

		if condition[2] != "" {
			fmt.Println("查寻维修带状态参数")
			err = db.Where("dorm_id=? AND repair_status=?", Dorm.Id, condition[2]).Find(&repairList).Error
			if err != nil {
				return nil, 0, err
			}
			return repairList, total, nil
		}

		err = db.Where("dorm_id=?", Dorm.Id).Find(&repairList).Count(&total).Error
		fmt.Println("找不到数据", err)
		if err != nil {
			fmt.Println("找不到数据", err)
			return nil, 0, err
		}

		// fmt.Println("获取维修不带状态")
		return repairList, total, nil
	}
	// ///////////////////////////////////////////////////////////////
	// fmt.Println("我是维修---99---")
	// 查寻数据
	fmt.Println("维修dormId++++++", dormId)
	if dormId != 0 {
		err := global.Global_Db.Model(&repair.Repair{}).
			Preload("Dorm", func(db *gorm.DB) *gorm.DB {
				return db.Model(&dorm.Dorm{}).Debug().
					Select("dorm.*,floor.floors_name AS floors_name").
					Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
			}).
			Where("dorm_id=?", dormId).
			Count(&total).
			Limit(limit).Offset(offset).Find(&repairList).Error
		if err != nil {
			// 处理错误
			return nil, 0, err
		}
		return repairList, total, nil
	}
	err := global.Global_Db.Model(&repair.Repair{}).
		Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().
				Select("dorm.*,floor.floors_name AS floors_name").
				Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).
		Where(condition).
		Count(&total).
		Limit(limit).Offset(offset).
		Find(&repairList).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return repairList, total, nil
}

// 更新
func (f *RepairService) UpdateRepair(Repair repair.Repair, roleId uint) error {
	var tempRepair repair.Repair
	err := global.Global_Db.Where("id=?", Repair.Id).First(&tempRepair).Error
	if err != nil {
		return errors.New("数据不存在,无法更新")
	}
	if Repair.RepairStatus == "已完成" {
		if roleId > 2 {
			return errors.New("状态已完成，权限不足，无法更新")
		}
		Repair.FinishDate = time.Now()
	}
	err = global.Global_Db.Model(&Repair).Where("id = ?", Repair.Id).Updates(Repair).Error
	if err != nil {
		// 处理错误
		return errors.New("更新失败")
	}
	return nil
}

// 删除
func (f *RepairService) DeleteRepair(_repairs *[]repair.Repair, roleId uint) error {
	// 遍历查寻数据是否存在
	for _, v := range *_repairs {
		var tempRepair repair.Repair
		err := global.Global_Db.Where("id=?", v.Id).First(&tempRepair).Error
		if err != nil {
			return errors.New(v.SubmitDate.Format("2006-01-02") + "至" + tempRepair.FinishDate.Format("2006-01-02"))
		}
		if tempRepair.RepairStatus == "已完成" {
			if roleId > 2 {
				return errors.New("状态已完成，权限不足，无法删除" + tempRepair.SubmitDate.Format("2006-01-02"))
			}
		}
	}
	err := global.Global_Db.Delete(_repairs).Error
	if err != nil {
		// 处理错误
		return errors.New("数据删除失败")
	}
	return nil

}

// 添加
func (f *RepairService) CreateRepair(_repairs *[]repair.Repair) error {
	// 给数据添加id
	for _, v := range *_repairs {
		var tempRepair dorm.Dorm
		err := global.Global_Db.Where("id=? ", v.DormId).First(&tempRepair).Error
		if err != nil {
			return errors.New("该宿舍不存在,无法添加")
		}
	}
	// 添加数据
	err := global.Global_Db.Create(&_repairs).Error
	if err != nil {
		// 处理错误
		return errors.New("添加维修失败")
	}
	return nil
}
