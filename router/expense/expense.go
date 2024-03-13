package expense

import "github.com/gin-gonic/gin"

type Expense struct {
}

func (e *Expense) Api_Expense(router *gin.RouterGroup) {
	Expense := router.Group("/Exp")
	{
		Expense.GET("getExpense", func(c *gin.Context) {})
		Expense.DELETE("deleteById", func(c *gin.Context) {})
		Expense.POST("/createExpense", func(c *gin.Context) {})
		Expense.PUT("/putExpense", func(c *gin.Context) {})
	}
}
