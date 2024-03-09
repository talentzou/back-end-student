package dorm

import "github.com/gin-gonic/gin"

type Dorm struct{}

func (d *Dorm) Api_Dorm(router *gin.RouterGroup){
	  dorm:=router.Group("/Dormitory")
	  {
		dorm.GET("getDorm",func(c *gin.Context){})
		dorm.DELETE("deleteDorm_ById",func(c *gin.Context){})
		dorm.POST("/createDorm",func(c *gin.Context){})
		dorm.PUT("/putDorm",func(c *gin.Context){})
	  }
}