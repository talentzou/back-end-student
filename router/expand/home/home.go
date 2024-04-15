package home

import (
	// "back-end/controller"
    "back-end/controller/test"
	"github.com/gin-gonic/gin"
)

type Home struct{}

func (e *Home) UseHome(router *gin.RouterGroup) {
	Repair := router.Group("/Home")
	{
		Repair.GET("getHomeMessage",test.GetHomeMessage)
	}
}