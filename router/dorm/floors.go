package dorm

import (
	"back-end/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Floor struct{}

func (f *Floor) Api_Floor(router *gin.RouterGroup) {
	Floor := router.Group("/Floor")
	apiRouterApi := controller.Dorm_api
	fmt.Println("wo执行到333333")
	{
		Floor.GET("getFloor", apiRouterApi.QueryApi)
		Floor.DELETE("deleteFloorById", func(c *gin.Context) {})
		Floor.POST("/createFloor", func(c *gin.Context) {})
		Floor.PUT("/putFloor", func(c *gin.Context) {})
	}
}
