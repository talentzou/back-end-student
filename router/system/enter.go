package system

import (
	"back-end/middleware"
	"github.com/gin-gonic/gin"
)

type SystemRouteGroup struct {
	SysLogin
	SysMenu
	SysUpload
	SysUser
}

func (S SystemRouteGroup) InitializeSys(R *gin.RouterGroup) {

	S.SysLogin.UseLogin(R)
	S.SysUpload.SystemUploadImg(R)
	// 鉴权
	PrivateGroup := R.Group("sys_jwt")
	PrivateGroup.Use(middleware.JwtAuth())
	S.SysMenu.GetMenu(PrivateGroup)
	S.SysUser.UserRouter(PrivateGroup)
}
