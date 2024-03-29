package dorm

import (
	// "back-end/controller"
	"back-end/controller/test"
	"github.com/gin-gonic/gin"
)

type Dorm struct{}

func (d *Dorm) Api_Dorm(router *gin.RouterGroup) {
	dorm := router.Group("/dorm")
	// dormRoute:=controller.Dorm_api
	dormRoute := test.Dorm_api
	{
		dorm.GET("getDorm/:Page/:PageSize", dormRoute.QueryDormApi)
		dorm.DELETE("deleteDorm_ById", dormRoute.DeleteDormApi)
		dorm.POST("/createDorm", dormRoute.CreateDormApi)
		dorm.PUT("/putDorm", dormRoute.UpdateDormApi)
	}
	

}
