package dorm

import "github.com/gin-gonic/gin"

type Stay struct{}

func (s *Stay) Api_Stay(router *gin.RouterGroup) {
	Stay := router.Group("/Stay")
	{
		Stay.GET("getStay", func(c *gin.Context) {})
		Stay.DELETE("deleteStayById", func(c *gin.Context) {})
		Stay.POST("/createStay", func(c *gin.Context) {})
		Stay.PUT("/putStay", func(c *gin.Context) {})
	}
}
