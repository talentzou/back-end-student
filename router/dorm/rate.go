package dorm

import "github.com/gin-gonic/gin"

type Rate struct{}

func (r *Rate) Api_Rate(router *gin.RouterGroup){
	  Rate:=router.Group("/Rate")
	  {
		Rate.GET("getRate",func(c *gin.Context){})
		Rate.DELETE("deleteRateById",func(c *gin.Context){})
		Rate.POST("/createRate",func(c *gin.Context){})
		Rate.PUT("/putRate",func(c *gin.Context){})
	  }
}