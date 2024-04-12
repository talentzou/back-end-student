package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysRole struct{}

func (s *SysBtn) RoleRoute(router *gin.RouterGroup) {
	menu := router.Group("role")
	{
		menu.GET("getRoleList", system.RoleApi.GetRoles)       // 获取角色列表
		menu.POST("createRoles", system.RoleApi.CreateRoles)   // 添加角色
		menu.DELETE("deleteRoles", system.RoleApi.DeleteRoles) // 删除角色
	}

}
