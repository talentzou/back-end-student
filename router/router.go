package router

import (
	"back-end/router/dorm"
	"back-end/router/expense"
	"back-end/router/login"
	"back-end/router/repair"
	"back-end/router/student"

	"github.com/gin-gonic/gin"
)

type AppRouterGroup struct {
	Dorm    dorm.DormGroup
	Login   login.Login
	Repair  repair.Repair
	Expense expense.ExpenseGroup
	Student student.StudentGroup
}

func (routers *AppRouterGroup) InitializeRouter(s *gin.Engine) {
	root := s.Group("/")
	dormRouter := routers.Dorm
	loginRouter := routers.Login
	repairRouter := routers.Repair
	expenseRouter := routers.Expense
	studentRouter := routers.Student
	{
		dormRouter.UseDormRouter(root)
		repairRouter.UseRepair(root)
		studentRouter.UseStudent(root)
		loginRouter.UseLogin(root)
		expenseRouter.UseExpense(root)

	}

}

var AppRouter = new(AppRouterGroup)
