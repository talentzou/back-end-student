package expense

import "github.com/gin-gonic/gin"

type ExpenseGroup struct {
	DormExpense Expense
}

func (e *ExpenseGroup) UseExpense(router *gin.RouterGroup) {
	routers := router.Group("/expense")
	{
		e.DormExpense.Api_Expense(routers)
	}
}
