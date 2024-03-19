package system

import "github.com/gin-gonic/gin"

type Login struct{}

func (e *Login) UseLogin(router *gin.RouterGroup) {
	Login := router.Group("/base")
	{
		Login.POST("/login", func(c *gin.Context) {})
	}
}
