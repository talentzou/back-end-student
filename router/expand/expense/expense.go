package expense

import (
	// "back-end/controller"
	"back-end/controller/test"
	"github.com/gin-gonic/gin"
)

type Expense struct {
}

func (e *Expense) Api_Expense(router *gin.RouterGroup) {
	Expense := router.Group("/Exp")
	// expenseRouterApi := controller.Expense_api
	expenseRouterApi := test.Expense_api
	{
		Expense.GET("getExpense/:Page/:PageSize", expenseRouterApi.QueryExpenseApi)
		Expense.DELETE("deleteById", expenseRouterApi.DeleteExpenseApi)
		Expense.POST("/createExpense", expenseRouterApi.CreateExpenseApi)
		Expense.PUT("/putExpense", expenseRouterApi.UpdateExpenseApi)
	}
}
