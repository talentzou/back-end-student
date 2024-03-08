package dorm

import "github.com/gin-gonic/gin"

type Floor struct{}

func (f *Floor) Api_Floor(router *gin.RouterGroup) {
	Floor := router.Group("/Floor")
	{
		Floor.GET("getFloor", func(c *gin.Context) {})
		Floor.DELETE("deleteFloorById", func(c *gin.Context) {})
		Floor.POST("/createFloor", func(c *gin.Context) {})
		Floor.PUT("/putFloor", func(c *gin.Context) {})
	}
}
