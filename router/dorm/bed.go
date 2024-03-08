package dorm

import "github.com/gin-gonic/gin"

type Bed struct{}

func (b *Bed) Api_Bed(router *gin.RouterGroup){
	  Bed:=router.Group("/Bed")
	  {
		Bed.GET("/getBed",func(c *gin.Context){})
		Bed.DELETE("/deleteBedById",func(c *gin.Context){})
		Bed.POST("/createBed",func(c *gin.Context){})
		Bed.PUT("/putBed",func(c *gin.Context){})
	  }
}