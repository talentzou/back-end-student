package student

import (
	"back-end/controller"
	"github.com/gin-gonic/gin"
)

type Student struct{}

func (e *Student) Api_Student(router *gin.RouterGroup) {
	Student := router.Group("/Std")
	studInfoRouterApi := controller.StudInfo_api

	{
		Student.GET("/getStudent", studInfoRouterApi.QueryStudInfoApi)
		Student.DELETE("/deleteById", studInfoRouterApi.DeleteStudInfoApi)
		Student.POST("/createStudent", studInfoRouterApi.CreateStudInfoApi)
		Student.PUT("/putStudent", studInfoRouterApi.UpdateStudInfoApi)
	}
}
