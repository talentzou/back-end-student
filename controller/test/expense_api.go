package test

import (
	"back-end/global"
	"back-end/model/common/request"
	"back-end/model/common/response"
	"back-end/model/test/dorm"
	"back-end/utils"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type expense_dorm_api struct{}

// 插入
func (d *expense_dorm_api) CreateExpenseApi(c *gin.Context) {
	var expenseList []dorm.Expense
	err := c.ShouldBindJSON(&expenseList)
	if err != nil {
		fmt.Println("错误", err.Error())
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 给数据添加id
	for _, v := range expenseList {
		var tempDorm dorm.Dorm
		query := global.Global_Db.Where("id=? ", v.DormId).First(&tempDorm)
		if query.Error != nil {
			response.FailWithMessage("宿舍不存在,无法添加", c)
			return
		}
	}
	// 添加数据
	result := global.Global_Db.Create(&expenseList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *expense_dorm_api) DeleteExpenseApi(c *gin.Context) {
	var expenseList []dorm.Expense
	err := c.ShouldBindJSON(&expenseList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range expenseList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage(value.PaymentTime.Format("2006-01-02")+"费用数据不存在", c)
			return
		}
	}
	for _, del := range expenseList {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage(del.PaymentTime.Format("2006-01-02")+"费用数据删除失败", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *expense_dorm_api) UpdateExpenseApi(c *gin.Context) {
	var dormExpense dorm.Expense
	err := c.ShouldBindJSON(&dormExpense)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	var temp dorm.Expense
	err2 := global.Global_Db.Model(&temp).Where("id=?", dormExpense.Id).First(&temp)
	if err2.Error != nil {
		response.FailWithMessage(dormExpense.PaymentTime.Format("2006-01-02")+"费用数据不存在", c)
		return
	}
	result := global.Global_Db.Model(&dormExpense).Where("id = ?", dormExpense.Id).Updates(dormExpense)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *expense_dorm_api) QueryExpenseApi(c *gin.Context) {
	var limit, offset int

	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	if er1 != nil && er2 != nil {
		response.FailWithMessage("分页参数错误", c)
		return
	}
	// 分页数据
	offset = PageSize * (Page - 1)
	limit = PageSize
	// 获取query
	rawUrl := c.Request.URL.String()
	u, er := url.Parse(rawUrl)
	if er != nil {
		fmt.Println("解析url错误")
	}
	queryParams := u.Query()
	// 获取请求体数据
	condition := make(map[string]interface{})
	for index, value := range queryParams {
		key := utils.ToCamelCase(index)
		condition[key] = value
	}
	fmt.Println("condition", condition)
	var arrSlice interface{}

	mapLength := len(condition)
	// fmt.Println("condition9999", mapLength,condition)
	if mapLength == 0 {
		arrSlice = nil
	} else {
		floor_dorm := condition["floor_dorm"].([]string)
		// fmt.Println("进来后++++",floor_dorm)
		words := strings.Split(floor_dorm[0], "-")
		arrSlice = words
	}

	// 获取学生用户所属宿舍
	var dormId uint
	if utils.GetUserRoleId(c) == 3 {
		dormId = utils.GetUserDormId(c)
	}

	expenseList, total, err := expenseService.QueryExpense(limit, offset, arrSlice, dormId)
	if err != nil {
		response.FailWithMessage("查询水电费信息失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     expenseList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var Expense_api = new(expense_dorm_api)
