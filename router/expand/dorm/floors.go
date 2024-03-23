package dorm

import (
	"back-end/controller"
	"github.com/gin-gonic/gin"
)

type Floor struct{}

func (f *Floor) Api_Floor(router *gin.RouterGroup) {
	Floor := router.Group("/Floor")
	apiRouterApi := controller.Floor_api
	{
		Floor.GET("/getFloor/:Page/:PageSize", apiRouterApi.QueryFloorApi)
		Floor.DELETE("/deleteFloorById", apiRouterApi.DeleteFloorApi)
		Floor.POST("/createFloor", apiRouterApi.CreateFloorApi)
		Floor.PUT("/putFloor", apiRouterApi.UpdateFloorApi)
	}
}
