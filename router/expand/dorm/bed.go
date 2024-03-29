package dorm

import (
	// "back-end/controller"
	"back-end/controller/test"
	"github.com/gin-gonic/gin"
)

type Bed struct{}

func (b *Bed) Api_Bed(router *gin.RouterGroup) {
	Bed := router.Group("/Bed")
	
	// bedRouter := controller.Bed_api
	bedRouter := test.Bed_api
	{
		Bed.GET("/getBed", bedRouter.QueryBedApi)
		Bed.DELETE("/deleteBedById", bedRouter.DeleteBedApi)
		Bed.POST("/createBed", bedRouter.CreateBedApi)
		Bed.PUT("/putBed", bedRouter.UpdateBedApi)
	}

	
}
