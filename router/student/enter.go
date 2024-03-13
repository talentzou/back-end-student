package student

import (
	"github.com/gin-gonic/gin"
)

type StudentGroup struct {
	StudentInfo Student
}

func (e *StudentGroup) UseStudent(router *gin.RouterGroup) {
	routers := router.Group("/student")
	{
		e.StudentInfo.Api_Student(routers)
	}
}
