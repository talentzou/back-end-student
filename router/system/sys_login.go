package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type Login struct{}

func (e *Login) UseLogin(router *gin.RouterGroup) {
	Login := router.Group("/base")
	{
		Login.POST("/login",system.SystemApi.Login)
	}
}
