package controller

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/apiexpense"
	"back-end/utils"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)



type expense_dorm_api struct{}

// 插入
func (d *expense_dorm_api) CreateExpenseApi(c *gin.Context) {
	var expenseList []apiexpense.Expense_dorm
	err := c.ShouldBindJSON(&expenseList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 给数据添加id
	for i, _ := range expenseList {
		words := strings.Split(expenseList[i].DormNumber, "-")
		if words[0] !=expenseList[i].FloorsName {
			fmt.Println(words[0] !=expenseList[i].FloorsName, "分割的单词",expenseList[i].FloorsName, words)
			response.FailWithMessage("宿舍"+expenseList[i].DormNumber+"开头前缀与宿舍楼不一致", c)
			return
		}
		uid := uuid.NewString()
		expenseList[i].Id = uid
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
	var expenseList []apiexpense.Expense_dorm
	err := c.ShouldBindJSON(&expenseList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range expenseList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage("宿舍:"+value.DormNumber+","+value.PaymentTime+"费用数据不存在", c)
			return
		}
	}
	for _, del := range expenseList {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("宿舍:"+del.DormNumber+","+del.PaymentTime+"费用数据删除失败", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *expense_dorm_api) UpdateExpenseApi(c *gin.Context) {
	var dormExpense apiexpense.Expense_dorm
	err := c.ShouldBindJSON(&dormExpense)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	var temp apiexpense.Expense_dorm
	err2 := global.Global_Db.Model(&temp).Where("id=?", dormExpense.Id).First(&temp)
	if err2.Error != nil {
		response.FailWithMessage("宿舍:"+dormExpense.DormNumber+","+dormExpense.PaymentTime+"费用数据不存在", c)
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
	var total int64
	var expenseList []apiexpense.Expense_dorm
	var expense apiexpense.Expense_dorm
	// 获取query
	rawUrl := c.Request.URL.String()
	u, er := url.Parse(rawUrl)
	if er != nil {
		fmt.Println("解析url错误")
	}
	queryParams := u.Query()
	// fmt.Println("查寻字符串参数", queryParams)
	// 获取请求体数据
	condition := make(map[string]interface{})
	for index, value := range queryParams {
		key := utils.ToCamelCase(index)
		condition[key] = value
	}
	fmt.Println("condition", condition)
	err := c.ShouldBindJSON(&pages)
	if err != nil {
		fmt.Println("错误为", err)
		response.FailWithMessage("系统错误", c)
		return
	}
	// 分页数据
	offset = pages.PageSize * (pages.Page - 1)
	limit = pages.PageSize
	fmt.Println(offset, limit)
	// 查寻数量
	count := global.Global_Db.Model(&expense).Where(condition).Count(&total).Error
	if count != nil {
		fmt.Println("计算楼层数量错误")
		response.FailWithMessage("系统查询错误", c)
		return
	}
	// 查寻数据
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&expenseList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("系统查寻失败", c)
		return
	}
	for key, v := range expenseList {
		time, err1 := time.Parse(time.RFC3339, v.PaymentTime)
		if err1 != nil {
			response.FailWithMessage("系统解析时间错误错误", c)
			return
		}
		date := time.Format("2006-01-02")
		expenseList[key].PaymentTime = date

	}
	response.OkWithDetailed(request.PageInfo{
		List:     expenseList,
		Total:    total,
		PageSize: pages.PageSize,
		Page:     pages.Page,
	}, "成功", c)

}

var Expense_api = new(expense_dorm_api)
