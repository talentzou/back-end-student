package test

import (
	"back-end/global"
	"back-end/model/common/response"
	"back-end/model/system"
	"back-end/model/test/dorm"
	// "fmt"

	"back-end/model/test/repair"

	"github.com/gin-gonic/gin"
)

type amount struct {
	User    int64 `json:"user"`
	Dorm    int64 `json:"dorm"`
	Student int64 `json:"student"`
	Repair  int64 `json:"repair"`
}

var total = new(amount)

func GetHomeMessage(c *gin.Context) {
	err := global.Global_Db.Model(&dorm.Dorm{}).
		Where("NOT EXISTS (SELECT * FROM stud_info WHERE dorm.id = stud_info.dorm_id)").
		Count(&total.Dorm).Error
	if err != nil {
		response.FailWithMessage("查询宿舍数量错误", c)
		return
	}

	err = global.Global_Db.Model(&repair.Repair{}).Where("repair_status = ?", "未完成").Count(&total.Repair).Error
	if err != nil {
		response.FailWithMessage("查询维修数量错误", c)
		return
	}
	err = global.Global_Db.Model(&system.SysUser{}).Count(&total.User).Error
	if err != nil {
		response.FailWithMessage("查询用户数量错误", c)
		return
	}

	err = global.Global_Db.Model(&dorm.StudInfo{}).Count(&total.Student).Error
	if err != nil {
		response.FailWithMessage("查询学生数量错误", c)
		return
	}
	// amount{
	// 	User:    total.User,
	// 	Student: total.Student,
	// 	Repair:  total.Repair,
	// 	Dorm:    total.Dorm,
	// }
	// fmt.Println("哈哈哈", total)
	response.Ok("查寻成功", map[int]interface{}{
		0: total.User,
		1: total.Dorm,
		2: total.Student,
		3: total.Repair,
	}, c)
}
