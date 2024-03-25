package expand

import (
	"back-end/router/expand/dorm"
	"back-end/router/expand/expense"
	"back-end/router/expand/notice"
	"back-end/router/expand/repair"
	"back-end/router/expand/student"
	"github.com/gin-gonic/gin"
)

type AppExpandRouterGroup struct {
	Dorm    dorm.DormGroup
	Repair  repair.Repair
	Expense expense.ExpenseGroup
	Student student.StudentGroup
	Notice  notice.Notice
}

func (routers *AppExpandRouterGroup) InitializeExpandRouter(R *gin.RouterGroup) {
	dormRouter := routers.Dorm
	repairRouter := routers.Repair
	expenseRouter := routers.Expense
	studentRouter := routers.Student
	noticeRouter := routers.Notice
	// jwt路由
	{
		dormRouter.UseDormRouter(R) //宿舍管理
		repairRouter.UseRepair(R)   //维修管理
		studentRouter.UseStudent(R) //学生管理
		expenseRouter.UseExpense(R) //费用管理
		noticeRouter.UseNotice(R)   //首页
	}
}
