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
	"strings"

	"github.com/gin-gonic/gin"
)

type dorm_rate_api struct{}

// 插入
func (d *dorm_rate_api) CreateRateApi(c *gin.Context) {
	var rateList []dorm.Rate
	err := c.ShouldBindJSON(&rateList)
	if err != nil {
		fmt.Println("结构体错误", err.Error())
		response.FailWithMessage("参数错误", c)
		return
	}

	
	err = rateService.CreateRate(&rateList)
	if err != nil {
		// 处理错误
		response.FailWithMessage(err.Error(), c)
		return

	}
	response.OkWithMessage("添加评分成功", c)
}

// 删除
func (d *dorm_rate_api) DeleteRateApi(c *gin.Context) {
	var rateList []dorm.Rate
	err := c.ShouldBindJSON(&rateList)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}

	err = rateService.DeleteRate(&rateList)
	if err != nil {
		// 处理错误
		response.FailWithMessage(err.Error(), c)
		return

	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *dorm_rate_api) UpdateRateApi(c *gin.Context) {
	var rate dorm.Rate
	err := c.ShouldBindJSON(&rate)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	
	err = rateService.UpdateRate(rate)
	if err != nil {
		// 处理错误
		response.FailWithMessage(err.Error(), c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *dorm_rate_api) QueryRateApi(c *gin.Context) {
	var limit, offset int
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	if er1 != nil && er2 != nil {
		response.FailWithMessage("分页参数错误", c)
		return
	}
	offset = PageSize * (Page - 1)
	limit = PageSize
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
	// 变量
	var arrSlice interface{}

	mapLength := len(condition)
	fmt.Println("condition9999", mapLength, condition)
	if mapLength == 0 {
		arrSlice = nil
	} else {
		floor_dorm := condition["floor_dorm"].([]string)
		// fmt.Println("进来后++++",floor_dorm)
		words := strings.Split(floor_dorm[0], "-")
		arrSlice = words
	}
	// 分页数据

	// fmt.Println(offset, limit)

	// 获取学生用户所属宿舍
	var dormId uint
	if utils.GetUserRoleId(c) == 3 {
		dormId = utils.GetUserDormId(c)
	}

	rateList, total, err := rateService.QueryRate(limit, offset, arrSlice, dormId)
	if err != nil {
		response.FailWithMessage("查寻宿舍评分信息失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     rateList,
		Total:    total,
		PageSize: pages.PageSize,
		Page:     pages.Page,
	}, "成功", c)
}

var Rate_api = new(dorm_rate_api)
