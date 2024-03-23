package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysLogin struct{}

func (e *SysLogin) UseLogin(router *gin.RouterGroup) {

	{
		router.POST("/login", system.SystemApi.Login)
	}
}
