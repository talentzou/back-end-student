package system

import (
	"back-end/global"
	"back-end/model/common/request"
	"back-end/model/common/response"
	"back-end/model/system"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	fmt.Println("参数为", userInfo)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err3 := global.Global_Db.Model(&system.SysUser{}).Where("uuid=?", uuid).Updates(&system.SysUser{
		Sex:       userInfo.Sex,
		Avatar:    userInfo.Avatar,
		Nickname:  userInfo.NickName,
		Telephone: userInfo.Telephone,
		Dorm:      userInfo.Dorm,
	}).Error
	if err3 != nil {
		response.FailWithMessage("用户信息更新失败", c)
		return
	}
	response.OkWithMessage("用户信息更新成功", c)

}

// 获取用户列表
func GetUserList(c *gin.Context) {
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	// 分页
	offset := PageSize * (Page - 1)
	limit := PageSize
	if er1 != nil && er2 != nil {
		fmt.Println("分页数错误", er1.Error(), er2.Error())
	}
	list, total, err := userService.GetUserInfoList(offset, limit)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageInfo{
		List:     list,
		Total:    total,
		Page:     Page,
		PageSize: PageSize,
	}, "获取成功", c)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		// global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// 添加或注册用户
func Register(c *gin.Context) {
	var r system.SysUser
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{UserName: r.UserName, Nickname: r.Nickname, Password: r.Password, Avatar: r.Avatar, Authority: r.Authority, Telephone: r.Telephone}
	userReturn, err := userService.Register(*user)
	if err != nil {
		response.FailWithDetailed(gin.H{"user": userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"user": userReturn}, "注册成功", c)
}
