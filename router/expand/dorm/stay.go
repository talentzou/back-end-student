package dorm

import (
	// "back-end/controller"
	"back-end/controller/test"
	"github.com/gin-gonic/gin"
)

type Stay struct{}

func (s *Stay) Api_Stay(router *gin.RouterGroup) {
	Stay := router.Group("/Stay")
	// stayRouterApi := controller.Stay_api
	stayRouterApi := test.Stay_api
	{
		Stay.GET("getStay/:Page/:PageSize", stayRouterApi.QueryFloorApi)
		Stay.DELETE("deleteStayById", stayRouterApi.DeleteFloorApi)
		Stay.POST("/createStay", stayRouterApi.CreateFloorApi)
		Stay.PUT("/putStay", stayRouterApi.UpdateFloorApi)
	}
}
