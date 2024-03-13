package repair

import (
	"back-end/controller"

	"github.com/gin-gonic/gin"
)

type Repair struct{}

func (e *Repair) UseRepair(router *gin.RouterGroup) {
	Repair := router.Group("/Repair")
	repairRouterApi:=controller.Repair_api

	{
		Repair.GET("getRepair", repairRouterApi.QueryRepairApi)
		Repair.DELETE("deleteById", repairRouterApi.DeleteRepairApi)
		Repair.POST("/createRepair", repairRouterApi.CreateRepairApi)
		Repair.PUT("/putRepair", repairRouterApi.UpdateRepairApi)
	}
}
