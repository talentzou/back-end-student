package dorm

import (
	"back-end/controller"
	"github.com/gin-gonic/gin"
)

type Floor struct{}

func (f *Floor) Api_Floor(router *gin.RouterGroup) {
	Floor := router.Group("/Floor")
	apiRouterApi := controller.Dorm_api
	{
		Floor.GET("/getFloor", apiRouterApi.QueryApi)
		Floor.DELETE("/deleteFloorById", apiRouterApi.DeleteApi)
		Floor.POST("/createFloor", apiRouterApi.CreateApi)
		Floor.PUT("/putFloor", apiRouterApi.UpdateApi)
	}
}
