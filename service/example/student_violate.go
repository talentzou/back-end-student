package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"back-end/model/test/student"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ViolateService struct{}

// 添加

func (v *ViolateService) CreateViolate(_violates *[]student.StudentViolate) error {
	for _, v := range *_violates {
		var tempVio student.StudentViolate
		err := global.Global_Db.Debug().
			Where(&student.StudentViolate{
				DormId:     v.DormId,
				RecordDate: v.RecordDate.Local(),
				StudInfoId: v.StudInfoId,
			}).
			First(&tempVio).Error

		fmt.Println("找不到", err, tempVio)
		if err == nil {
			return errors.New("留宿申请已存在,无法添加")
		}
	}
	// 添加数据
	err := global.Global_Db.Create(_violates).Error
	if err != nil {
		// 处理错误
		return errors.New("添加学生违纪信息失败")
	}
	return nil
}

// 删除
func (v *ViolateService) DeleteViolate(_violates *[]student.StudentViolate) error {
	for _, v := range *_violates {
		var tempDorm dorm.Dorm
		fmt.Println("参数为", v)
		fmt.Println("参数id", v.DormId)
		err := global.Global_Db.Where("id=? ", v.DormId).First(&tempDorm).Error
		if err != nil {
			return errors.New("该宿舍不存在无法删除")
		}
	}
	// 添加数据
	err := global.Global_Db.Delete(_violates).Error
	if err != nil {
		// 处理错误
		fmt.Println("错误信息为", err)
		return errors.New("删除学生违纪信息失败")
	}
	return nil
}

// 更新
func (v *ViolateService) UpdateViolate(_violate student.StudentViolate) error {

	var tempStudent student.StudentViolate
	err := global.Global_Db.Where("id=? ", _violate.Id).First(&tempStudent).Error
	if err != nil {
		return errors.New(_violate.Violate + "数据不存在")
	}
	fmt.Println("更新", _violate.RecordDate.Local())
	fmt.Println("存在", tempStudent.RecordDate)
	var tempVio student.StudentViolate
	err = global.Global_Db.Debug().
		Not("id=?", _violate.Id).
		Where(&student.StudentViolate{
			DormId:     _violate.DormId,
			RecordDate: _violate.RecordDate.Local(),
			StudInfoId: _violate.StudInfoId,
		}).
		First(&tempVio).Error

	// fmt.Println("找不到", err, tempVio)
	if err == nil {
		return errors.New(tempVio.RecordDate.Format("2006-01-02") + "留宿申请已存在,无法更新")

	}
	err = global.Global_Db.Model(&_violate).Where("id = ?", _violate.Id).Updates(_violate).Error
	if err != nil {
		// 处理错误
		return errors.New("更新学生:" + _violate.Violate + "失败")

	}
	return nil
}

// 查询学生违纪信息
func (f *ViolateService) QueryStudentViolateList(limit int, offset int, condition []string, dormId uint) (interface{}, int64, error) {
	var studentList []student.StudentViolate
	var total int64
	mapLength := len(condition)
	fmt.Println("我是学生违纪+++++++++++++", condition, mapLength)
	if mapLength != 0 {

		// fmt.Println("进来参数ggg99999", condition[0], condition[1], condition[2])
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

		var Student []dorm.StudInfo
		var studentIdList []uint
		if condition[2] != "" {
			err = global.Global_Db.Model(&dorm.StudInfo{}).Where("student_name =?", condition[2]).Find(&Student).Error
			if err != nil {
				return nil, 0, err
			}
			for _, v := range Student {
				studentIdList = append(studentIdList, v.Id)
			}
			// fmt.Println("查寻违纪宿舍带人名", Student.StudentName, "学生id", Student.Id, "宿舍id", Dorm.Id)
		}

		db := global.Global_Db.Model(&student.StudentViolate{}).
			Preload("Dorm", func(db *gorm.DB) *gorm.DB {
				return db.Model(&dorm.Dorm{}).Debug().
					Select("dorm.*,floor.floors_name AS floors_name").
					Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
			}).
			Preload("StudInfo").
			Limit(limit).Offset(offset)

		//.Debug().Select("dorm.*,floor.floors_name AS floors_name").Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		if condition[2] != "" {
			err = db.Model(&student.StudentViolate{}).Where("dorm_id=? AND stud_info_id IN ?", Dorm.Id, studentIdList).Find(&studentList).Count(&total).Error
			fmt.Println("是否找到学生数据", err)
			fmt.Println("学生数据为", studentList)
			if err != nil {
				return nil, 0, err
			}
			return studentList, total, nil
		}

		err = db.Where("dorm_id=?", Dorm.Id).Find(&studentList).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		fmt.Println("查寻违纪宿舍带不带参数")
		return studentList, total, nil
	}
	// //////////////////////////////////////////////////
	fmt.Println("我是学生违纪---99---", dormId)
	// 查寻数据
	if dormId != 0 {
		db := global.Global_Db.Model(&student.StudentViolate{}).
			Preload("StudInfo").
			Preload("Dorm", func(db *gorm.DB) *gorm.DB {
				return db.Model(&dorm.Dorm{}).Debug().
					Select("dorm.*,floor.floors_name AS floors_name").
					Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
			})
		err := db.Where("dorm_id=?", dormId).Limit(limit).Offset(offset).Find(&studentList).Count(&total).Error
		if err != nil {
			// 处理错误
			return nil, 0, err
		}
		return studentList, total, nil
	}

	db := global.Global_Db.Model(&student.StudentViolate{}).
		Preload("StudInfo").
		Preload("Dorm", func(db *gorm.DB) *gorm.DB {
			return db.Model(&dorm.Dorm{}).Debug().
				Select("dorm.*,floor.floors_name AS floors_name").
				Joins("LEFT JOIN floor ON dorm.floor_id = floor.id")
		})
	err := db.Count(&total).Limit(limit).Offset(offset).Find(&studentList).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return studentList, total, nil
}
