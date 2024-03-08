package expense
import "github.com/gin-gonic/gin"

type ExpenseGroup struct{
	DormExpense Expense 
}

func (e *ExpenseGroup) UseExpense(router *gin.RouterGroup) {
	router.Group("/expense")
}
