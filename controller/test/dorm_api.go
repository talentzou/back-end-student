package test

import (
	// "back-end/global"
	"back-end/model/common/request"
	"back-end/model/common/response"
	"back-end/model/test/dorm"
	"back-end/utils"
	"fmt"
	"net/url"
	"strconv"
	// "strings"

	"github.com/gin-gonic/gin"
)

type dorm_api struct{}

// 插入
func (d *dorm_api) CreateDormApi(c *gin.Context) {
	var dormList []dorm.Dorm
	err := c.ShouldBindJSON(&dormList)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = dormService.CreateDorm(&dormList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *dorm_api) DeleteDormApi(c *gin.Context) {
	var dormList []dorm.Dorm
	err := c.ShouldBindJSON(&dormList)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = dormService.DeleteDorm(&dormList)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	// 处理错误
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *dorm_api) UpdateDormApi(c *gin.Context) {
	var Dorm dorm.Dorm
	err := c.ShouldBindJSON(&Dorm)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	result := dormService.UpdateDorm(Dorm)
	if result != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 查寻
func (d *dorm_api) QueryDormApi(c *gin.Context) {
	var limit, offset int
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	if er1 != nil && er2 != nil {
		response.FailWithMessage("分页参数错误", c)
		return
	}
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
	// 分页
	offset = PageSize * (Page - 1)
	limit = PageSize
	fmt.Println(offset, limit)

	dormList, total, err := dormService.QueryDorm(limit, offset, condition)
	if err != nil {
		response.FailWithMessage("查寻失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     dormList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

// 查寻宿舍与学生信息
func (d *dorm_api) GetDormWithStudent(c *gin.Context) {

	dormList, _, err := dormService.QueryDorm(0, 0, nil)
	if err != nil {
		response.FailWithMessage("查寻失败", c)
		return
	}
	response.Ok("获取成功", dormList, c)
}

var Dorm_api = new(dorm_api)
