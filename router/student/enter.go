package student

import "github.com/gin-gonic/gin"

type StudentGroup struct{
	StudentInfo Student
}

func (e *StudentGroup) UseStudent(router *gin.RouterGroup) {
	router.Group("/student")
}
