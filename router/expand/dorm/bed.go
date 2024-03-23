package dorm

import ("github.com/gin-gonic/gin"
"back-end/controller")

type Bed struct{}

func (b *Bed) Api_Bed(router *gin.RouterGroup){
	  Bed:=router.Group("/Bed")
	  bedRouter:=controller.Bed_api
	  {
		Bed.GET("/getBed",bedRouter.QueryBedApi)
		Bed.DELETE("/deleteBedById",bedRouter.DeleteBedApi)
		Bed.POST("/createBed",bedRouter.CreateBedApi)
		Bed.PUT("/putBed",bedRouter.UpdateBedApi)
	  }
}