package controller

import (
	"back-end/common/response"
	"back-end/global"
	"back-end/model/apidorm"
	"back-end/model/apirepair"
	"back-end/model/system"

	"github.com/gin-gonic/gin"
)

type amount struct {
	user_amount       int64
	empty_dorm_amount int64
	student_amount    int64
	repair_amount     int64
}

var total = new(amount)

func GetHomeMessage(c *gin.Context) {
	var bed apidorm.Bed_api
	var repair apirepair.Repair_dorm
	var user system.SysUser
	// var dorm
	err1 := global.Global_Db.Model(&bed).Where("bed_status = ?", "有人").Count(&total.student_amount).Error
	err2 := global.Global_Db.Model(&repair).Where("repair_status = ?", "未完成").Count(&total.repair_amount).Error
	err3 := global.Global_Db.Model(&user).Count(&total.user_amount).Error
	if err1 != nil || err2 != nil || err3 != nil {
		response.FailWithMessage("查寻数量错误", c)
	}
	response.Ok("查寻成功",total, c)
}
