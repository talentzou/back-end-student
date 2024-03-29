package notice

import (
	// "back-end/controller"
	"back-end/controller/test"
	"github.com/gin-gonic/gin"
)

type Notice struct{}

func (e *Notice) UseNotice(router *gin.RouterGroup) {
	Repair := router.Group("/Notice")
	// repairRouterApi := controller.Notice_api
	repairRouterApi := test.Notice_api
	{
		Repair.GET("getNotice/:Page/:PageSize", repairRouterApi.QueryNoticeApi)
		Repair.DELETE("deleteById", repairRouterApi.DeleteNoticeApi)
		Repair.POST("/createNotice", repairRouterApi.CreateNoticeApi)
		Repair.PUT("/putNotice", repairRouterApi.UpdateNoticeApi)
	}
}
