package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysMenu struct{}

func (s *SysMenu) GetMenu(router *gin.RouterGroup) {
	{
		router.GET("getMenu/:authorityId", system.GetAsyncMenu)
	}

}
