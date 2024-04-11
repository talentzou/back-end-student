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
		menu.GET("getAllMenu", system.MenuApi.GetAllMenu) //获取所有菜单
		menu.GET("getSelfMenu", system.MenuApi.GetSelfMenu) // 获取用户个人菜单
	}

}
