package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysMenu struct{}

func (s *SysMenu) MenuRoute(router *gin.RouterGroup) {
	menu := router.Group("menu")
	{
		//通过配置文件获取
		// router.GET("getMenu/:authorityId", system.GetAsyncMenu)

	}
	{
		menu.GET("getAllMenu", system.MenuApi.GetAllMenu)                              //获取所有菜单
		menu.GET("getSelfMenu", system.MenuApi.GetSelfMenu)                            // 获取用户个人菜单
		menu.POST("addRelateRoleAndMenu", system.MenuApi.AddRelateRoleAndMenu)         //添加角色菜单关联
		menu.DELETE("deleteRelateRoleAndMenu", system.MenuApi.DeleteRelateRoleAndMenu) //删除角色菜单关联
	}

}
