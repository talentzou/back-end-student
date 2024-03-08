package student

import "github.com/gin-gonic/gin"

type Student struct{}

 
func (e *Student) Api_Student(router *gin.RouterGroup) {
	Student := router.Group("/Student")
	{
		Student.GET("getStudent", func(c *gin.Context) {})
		Student.DELETE("deleteById", func(c *gin.Context) {})
		Student.POST("/createStudent", func(c *gin.Context) {})
		Student.PUT("/putStudent", func(c *gin.Context) {})
	}
}
