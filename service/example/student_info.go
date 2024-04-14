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
func (f *StudentService) QueryStudentViolateList(limit int, offset int, condition interface{}) (interface{}, int64, error) {
	var studentList []student.StudentViolate
	var total int64
	fmt.Println("我是学生违纪+++++++++++++", condition)
	if condition == nil {
		db := global.Global_Db.Model(&dorm.StudInfo{}).Limit(limit).Offset(offset)
		err := db.Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		}).Preload("StudentInfo").Find(&studentList).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("获取评分不带参数")
		return studentList, total, nil
	}
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
