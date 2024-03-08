package repair

import "github.com/gin-gonic/gin"

type Repair struct{}

func (e *Repair) UseRepair(router *gin.RouterGroup) {
	Repair := router.Group("/Repair")
	{
		Repair.GET("getRepair", func(c *gin.Context) {})
		Repair.DELETE("deleteById", func(c *gin.Context) {})
		Repair.POST("/createRepair", func(c *gin.Context) {})
		Repair.PUT("/putRepair", func(c *gin.Context) {})
	}
}
