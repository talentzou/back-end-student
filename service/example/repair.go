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
func (f *RepairService) QueryRepair(limit int, offset int, condition []string) (interface{}, int64, error) {
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
		db := global.Global_Db.Model(&repair.Repair{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Limit(limit).Offset(offset)

		if condition[2] != "" {
			fmt.Println("查寻维修带状态参数")
			err = db.Where("dorm_id=? AND repair_status=?", Dorm.Id, condition[2]).Find(&repairList).Error
			if err != nil {
				return nil, 0, err
			}
			return repairList, total, nil
		}

		err = db.Where("dorm_id=?", Dorm.Id).Find(&repairList).Count(&total).Error
		fmt.Println("找不到数据",err)
		if err != nil {
			fmt.Println("找不到数据",err)
			return nil, 0, err
		}

		// fmt.Println("获取维修不带状态")
		return repairList, total, nil
	}
	// ///////////////////////////////////////////////////////////////
	fmt.Println("我是维修---99---")
	// 查寻数据
	err := global.Global_Db.Model(&repair.Repair{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Where(condition).Count(&total).Limit(limit).Offset(offset).Find(&repairList).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return repairList, total, nil
}
