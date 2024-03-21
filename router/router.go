package router

import (
	"back-end/middleware"
	"back-end/router/dorm"
	"back-end/router/expense"
	"back-end/router/repair"
	"back-end/router/student"
	"back-end/router/system"

	"github.com/gin-gonic/gin"
)

type AppRouterGroup struct {
	Dorm    dorm.DormGroup
	Login   system.Login
	Repair  repair.Repair
	Expense expense.ExpenseGroup
	Student student.StudentGroup
}

func (routers *AppRouterGroup) InitializeRouter(s *gin.Engine) {
	s.Static("/public", "./public")
	root := s.Group("")
	dormRouter := routers.Dorm
	loginRouter := routers.Login
	repairRouter := routers.Repair
	expenseRouter := routers.Expense
	studentRouter := routers.Student
	// 系统路由
	{
		loginRouter.UseLogin(root)
	}
	ExpandRouterGroup := root.Group("/jwt")
	ExpandRouterGroup.Use(middleware.JwtAuth())
	// jwt路由
	{
		dormRouter.UseDormRouter(ExpandRouterGroup)
		repairRouter.UseRepair(ExpandRouterGroup)
		studentRouter.UseStudent(ExpandRouterGroup)
		expenseRouter.UseExpense(ExpandRouterGroup)
		system.SystemUploadImg(ExpandRouterGroup)

	}

}

var AppRouter = new(AppRouterGroup)
