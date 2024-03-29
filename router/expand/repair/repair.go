package repair

import (
	// "back-end/controller"
	"back-end/controller/test"
	"github.com/gin-gonic/gin"
)

type Repair struct{}

func (e *Repair) UseRepair(router *gin.RouterGroup) {
	Repair := router.Group("/Repair")
	// repairRouterApi:=controller.Repair_api
	repairRouterApi:=test.Repair_api
	{
		Repair.GET("getRepair/:Page/:PageSize", repairRouterApi.QueryRepairApi)
		Repair.DELETE("deleteById", repairRouterApi.DeleteRepairApi)
		Repair.POST("/createRepair", repairRouterApi.CreateRepairApi)
		Repair.PUT("/putRepair", repairRouterApi.UpdateRepairApi)
	}
}
