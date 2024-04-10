package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysUser struct{}

func (u *SysUser) UserRouter(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.POST("admin_register", system.Register)                //用户注册
		user.GET("getUserInfo", system.GetUserInfo)                 //获取用户个人信息
		user.POST("setSelfInfo", system.SetSelfInfo)                //用户设置个人信息
		user.GET("getUserList/:Page/:PageSize", system.GetUserList) // 分页获取用户列表
		user.DELETE("deleteUser", system.DeleteUser)                //删除用户
		user.POST("setUserInfo", system.SetUserInfo)                //系统设置信息
	}

	{
	

	}
	{
		user.GET("getSelfBtn", system.BtnApi.GetSelfBtns) //获取用户啊啊牛
	}
}
