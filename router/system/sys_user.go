package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysUser struct{}

func (u *SysUser) UserRouter(r *gin.RouterGroup) {
	user := r.Group("user")
	user.GET("getUserInfo", system.GetUserInfo)
}

