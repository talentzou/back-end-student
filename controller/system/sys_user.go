package system

import (
	"back-end/common/response"
	"back-end/global"
	"back-end/model/system"
	"back-end/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	fmt.Println("我是用户数据")
	id := utils.GetUserID(c)
	var ResUser system.SysUser
	fmt.Println("uuid", id)
	err := global.Global_Db.Model(&system.SysUser{}).Preload("SysAuthorityBtns").Where("id=?", id).First(&ResUser).Error
	if err != nil {
		fmt.Println("获取用户信息失败")
		response.FailWithMessage("获取用户信息失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ResUser}, "获取用户信息成功", c)
}

// 设置用户信息
func SetUserInfo(c *gin.Context) {
	uuid := utils.GetUserID(c)
	var userInfo system.ChangeUserInfo
	err := c.ShouldBindJSON(&userInfo)
	fmt.Println("参数为",userInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err3 := global.Global_Db.Model(&system.SysUser{}).Where("uuid=?", uuid).Updates(&system.SysUser{
		Sex: userInfo.Sex,
		Avatar: userInfo.Avatar,
		Nickname: userInfo.NickName,
		Telephone: userInfo.Telephone,
		Dorm: userInfo.Dorm,
	}).Error
	if err3 != nil {
		response.FailWithMessage("用户信息更新失败", c)
		return
	}
	response.OkWithMessage("用户信息更新成功",c)

}
