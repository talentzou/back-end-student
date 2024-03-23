package system

import (
	"back-end/common/response"
	"back-end/global"
	"back-end/model/system"
	"back-end/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	fmt.Println("我是用户数据")
	uuid := utils.GetUserUUID(c)
	var ResUser system.SysUser
	fmt.Println("uuid",uuid)
	err := global.Global_Db.Model(&system.SysUser{}).Where("uuid=?", uuid).First(&ResUser).Error
	if err != nil {
		fmt.Println("获取用户信息失败")
		response.FailWithMessage("获取用户信息失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ResUser}, "获取用户信息成功", c)
}
