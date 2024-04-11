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
	SysBtn
	SysRole
}

func (S SystemRouteGroup) InitializeSys(R *gin.RouterGroup) {

	S.SysLogin.UseLogin(R)
	S.SysUpload.SystemUploadImg(R)
	// 鉴权
	PrivateGroup := R.Group("sys_jwt")
	PrivateGroup.Use(middleware.JwtAuth())
	S.SysUser.UserRouter(PrivateGroup)
	S.SysMenu.MenuRoute(PrivateGroup)
	S.SysBtn.BtnRoute(PrivateGroup)
	S.SysBtn.RoleRoute(PrivateGroup)

}
