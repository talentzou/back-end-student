package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"back-end/model/test/student"
	"fmt"
	"gorm.io/gorm"
)

type StudentService struct{}

// 查寻学生信息
func (f *StudentService) QueryStudentInfoList(limit int, offset int, condition interface{}) (interface{}, int64, error) {
	var studentList []dorm.StudInfo
	var total int64
	fmt.Println("我是学生信息+++++++++++++", condition)
	if condition == nil {
		db := global.Global_Db.Model(&dorm.StudInfo{}).Limit(limit).Offset(offset)
		err := db.Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Find(&studentList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分不带参数")
		return studentList, total, nil
	}
	fmt.Println("我是学生信息---99---")
	// 查寻数据
	err := global.Global_Db.Model(&dorm.StudInfo{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	}).Where(condition).Limit(limit).Offset(offset).Find(&studentList).Count(&total).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return studentList, total, nil
}

// 查询学生违纪信息
func (f *StudentService) QueryStudentViolateList(limit int, offset int, condition []string) (interface{}, int64, error) {
	var studentList []student.StudentViolate
	var total int64
	mapLength := len(condition)
	fmt.Println("我是学生违纪+++++++++++++", condition, mapLength)
	if mapLength != 0 {
		// db := global.Global_Db.Model(&dorm.StudInfo{}).Limit(limit).Offset(offset)
		// err := db.Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		// 	return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		// }).Preload("StudentInfo").Find(&studentList).Error
		// if err != nil {
		// 	return nil, 0, err
		// }
		fmt.Println("进来参数ggg99999", condition[0], condition[1], condition[2])
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
		var Student dorm.StudInfo
		err = global.Global_Db.Model(&dorm.StudInfo{}).Where("student_name =?", condition[2]).First(&Student).Error
		if err != nil {
			return nil, 0, err
		}
		/*
			, func(db *gorm.DB) *gorm.DB {
				// if condition[2] != "" {
				// 	fmt.Println("学生有参数88777")
				// 	return db.Debug().Where("student_name =?", condition[2])
				// }
				// return db
			}
		*/
		// .Or("student_name = ?", condition[1])
		db := global.Global_Db.Model(&student.StudentViolate{}).Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Preload("StudInfo").Limit(limit).Offset(offset)

		//.Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		if condition[2] != "" {
			fmt.Println("查寻违纪宿舍带人名")
			err = db.Where("dorm_id=? AND stud_info_id=?", Dorm.Id, Student.Id).Find(&studentList).Error
			if err != nil {
				return nil, 0, err
			}
			return studentList, total, nil
		}

		err = db.Where("dorm_id=?", Dorm.Id).Find(&studentList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("查寻违纪宿舍带不带参数")
		return studentList, total, nil
	}
	// //////////////////////////////////////////////////
	fmt.Println("我是学生违纪---99---")
	// 查寻数据
	db := global.Global_Db.Model(&student.StudentViolate{}).Preload("StudInfo").Preload("Dorm", func(db *gorm.DB) *gorm.DB {
		return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
	})
	err := db.Where(condition).Limit(limit).Offset(offset).Find(&studentList).Count(&total).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return studentList, total, nil
}
