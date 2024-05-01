package example

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"errors"
	"strings"

	// "back-end/model/test/student"
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
		db := global.Global_Db.Model(&dorm.StudInfo{}).Count(&total).Limit(limit).Offset(offset)
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
	}).Where(condition).Count(&total).Limit(limit).Offset(offset).Find(&studentList).Error
	if err != nil {
		// 处理错误
		return nil, 0, err
	}
	return studentList, total, nil
}

// 更新
func (f *StudentService) UpdateStudentInfo(_stud_info dorm.StudInfo) error {
	var tempStudent dorm.StudInfo

	err := global.Global_Db.Where("id=?", _stud_info.Id).First(&tempStudent).Error
	if err != nil {
		return errors.New("更新的学生:" + tempStudent.StudentName + "数据不存在")
	}
	var tempDorm dorm.Dorm
	err = global.Global_Db.Where("id=?", _stud_info.DormId).Preload("Floor").First(&tempDorm).Error
	if err != nil {
		return errors.New("更新宿舍不存在")
	}
	fmt.Println("宿舍id", _stud_info.DormId)
	// 宿舍数量
	count := global.Global_Db.Model(&dorm.Dorm{Id: _stud_info.DormId}).Association("StudInfos").Count()
	// fmt.Println("Floor宿舍类型", tempDorm.Floor.FloorsType, "sex", _stud_info.Sex, "count", count, "tempDorm.Capacity", tempDorm.Capacity)
	// fmt.Println("+++", strings.Contains(tempDorm.Floor.FloorsType,_stud_info.Sex))
	isHas := strings.Contains(tempDorm.Floor.FloorsType, _stud_info.Sex)
	if !isHas {
		return errors.New("宿舍楼类型与学生性别不一致")
	}
	if _stud_info.DormId == tempStudent.DormId {
	} else {
		if count >= int64(tempDorm.Capacity) {
			return errors.New("该宿舍容量已达到最大值")
		}
	}
	// if strings.ContainsRune()

	err = global.Global_Db.Model(&dorm.StudInfo{}).Where("id= ?", _stud_info.Id).Updates(_stud_info).Error
	if err != nil {
		// 处理错误
		return errors.New("更新失败")
	}
	return nil
}

// 添加
func (f *StudentService) CreateStudentInfo(_stud_info *[]dorm.StudInfo) error {

	for _, v := range *_stud_info {
		var dorm_message dorm.Dorm
		err := global.Global_Db.Model(&dorm.Dorm{}).Where("id=?", v.DormId).Preload("Floor").First(&dorm_message).Error
		if err != nil {
			continue
		}
		// 查寻宿舍学生数量
		count := global.Global_Db.Model(&dorm.Dorm{Id: v.DormId}).Association("StudInfos").Count()
		if count >= int64(dorm_message.Capacity) {
			return errors.New("该宿舍:" + dorm_message.DormNumber + "容量已达到最大值")
		}
		//判断性别
		isHas := strings.Contains(dorm_message.Floor.FloorsType, v.Sex)
		if !isHas {
			return errors.New("宿舍楼类型与学生性别不一致")
		}

		//查询存在数据
		var tempStud dorm.StudInfo
		err = global.Global_Db.Where("student_number=?", v.StudentNumber).First(&tempStud).Error
		if err != nil {
			continue
		}
		if tempStud.StudentNumber == v.StudentNumber {
			return errors.New("该学号:" + v.StudentNumber + "学生已存在")
		}

	}
	// 添加数据
	err := global.Global_Db.Create(_stud_info).Error
	if err != nil {
		// 处理错误
		return errors.New("添加学生信息失败")
	}

	return nil
}

// 删除
func (f *StudentService) DeleteStudentInfo(_stud_infos *[]dorm.StudInfo) error {
	// 遍历查寻数据是否存在
	for _, value := range *_stud_infos {
		var student dorm.StudInfo
		err := global.Global_Db.Model(&student).Where("Id=?", value.Id).First(&student).Error
		if err != nil {
			return errors.New("删除学号为:" + value.StudentNumber + "数据不存在")
		}
	}
	err := global.Global_Db.Delete(_stud_infos).Error
	if err != nil {
		// 处理错误
		return errors.New("删除失败")
	}
	return nil
}
