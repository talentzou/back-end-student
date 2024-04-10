package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysBtn struct{}

func (s *SysBtn) BtnRoute(router *gin.RouterGroup) {
	menu := router.Group("btn")
	{
		menu.GET("getSelfBtn", system.BtnApi.GetSelfBtns) // 分页获取用户菜单
	}

}
