package student

import (
	// "back-end/controller"
	"back-end/controller/test"
	"github.com/gin-gonic/gin"
)

type Student struct{}

func (e *Student) Api_Student(router *gin.RouterGroup) {
	Student := router.Group("/Std")
	// studInfoRouterApi := controller.StudInfo_api
	studInfoRouterApi := test.StudInfo_api

	{
		Student.GET("/getStudent/:Page/:PageSize", studInfoRouterApi.QueryStudInfoApi)
		Student.DELETE("/deleteById", studInfoRouterApi.DeleteStudInfoApi)
		Student.POST("/createStudent", studInfoRouterApi.CreateStudInfoApi)
		Student.PUT("/putStudent", studInfoRouterApi.UpdateStudInfoApi)
	}
}
