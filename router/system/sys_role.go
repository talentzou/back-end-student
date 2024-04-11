package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysRole struct{}

func (s *SysBtn) RoleRoute(router *gin.RouterGroup) {
	menu := router.Group("role")
	{
		menu.GET("getRoleList", system.RoleApi.GetRoles) // 获取用户列表
	}

}
