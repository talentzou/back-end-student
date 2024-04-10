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
		router.GET("getMenu/:authorityId", system.GetAsyncMenu)

	}
	{
		menu.GET("getSelfMenu", system.MenuApi.GetMenu) // 分页获取用户菜单
	}

}
