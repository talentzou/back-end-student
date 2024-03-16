package dorm

import (
	"back-end/controller"
	"github.com/gin-gonic/gin"
)

type Rate struct{}

func (r *Rate) Api_Rate(router *gin.RouterGroup) {
	Rate := router.Group("/Rate")
	rateRouterApi := controller.Rate_api
	{
		Rate.GET("getRate/:Page/:PageSize", rateRouterApi.QueryRateApi)
		Rate.DELETE("deleteRateById", rateRouterApi.DeleteRateApi)
		Rate.POST("/createRate", rateRouterApi.CreateRateApi)
		Rate.PUT("/putRate", rateRouterApi.UpdateRateApi)
	}
}
