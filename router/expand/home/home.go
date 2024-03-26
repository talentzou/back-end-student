package home

import (
	"back-end/controller"

	"github.com/gin-gonic/gin"
)

type Home struct{}

func (e *Home) UseHome(router *gin.RouterGroup) {
	Repair := router.Group("/Home")
	{
		Repair.GET("getRepair/:Page/:PageSize",controller.GetHomeMessage)
	}
}