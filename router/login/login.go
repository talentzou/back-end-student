package login

import "github.com/gin-gonic/gin"

type Login struct{}

func (e *Login) UseLogin(router *gin.RouterGroup) {
	Login := router.Group("/Login")
	{
		Login.GET("getLogin", func(c *gin.Context) {})
		Login.DELETE("deleteById", func(c *gin.Context) {})
		Login.POST("/createLogin", func(c *gin.Context) {})
		Login.PUT("/putLogin", func(c *gin.Context) {})
	}
}
