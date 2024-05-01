package test

import (
	"back-end/model/common/request"
	"back-end/model/common/response"

	// "back-end/model/test/dorm"
	"back-end/model/test/student"
	"back-end/utils"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	// "strings"

	"github.com/gin-gonic/gin"
)

type violate_info_api struct{}

// 插入
func (d *violate_info_api) CreateVioApi(c *gin.Context) {
	var violateList []student.StudentViolate
	err := c.ShouldBindJSON(&violateList)
	if err != nil {
		fmt.Println("添加参数错误", err.Error())
		response.FailWithMessage("系统合并参数错误", c)
		return
	}

	err = violateService.CreateViolate(&violateList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *violate_info_api) DeleteVioApi(c *gin.Context) {
	var violateList []student.StudentViolate

	err := c.ShouldBindJSON(&violateList)
	if err != nil {
		fmt.Println("参数错误为", err.Error())
		response.FailWithMessage("系统合并参数错误", c)
		return
	}

	err = violateService.DeleteViolate(&violateList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *violate_info_api) UpdateVioApi(c *gin.Context) {
	var violate student.StudentViolate
	err := c.ShouldBindJSON(&violate)
	if err != nil {
		fmt.Println("参数错误", err.Error())
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	err = violateService.UpdateViolate(violate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 查寻
func (d *violate_info_api) QueryVioApi(c *gin.Context) {
	var limit, offset int
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	if er1 != nil && er2 != nil {
		fmt.Println("分页数错误", er1.Error(), er2.Error())
	}
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

	// 分页数据
	offset = PageSize * (Page - 1)
	limit = PageSize

	arrSlice := make([]string, 3)
	mapLength := len(condition)
	fmt.Println("condition9999", mapLength, condition)
	if mapLength == 0 {
		fmt.Println("进来为空pppp")
		arrSlice = nil
	} else {
		fmt.Println("进来不为空++++++++++")
		if floorDorm, ok := condition["floor_dorm"].([]string); ok {
			words := strings.Split(floorDorm[0], "-")
			arrSlice[0] = words[0]
			arrSlice[1] = words[1]
		}
		if studentName, ok := condition["student_name"].([]string); ok {
			arrSlice[2] = studentName[0]
		}

	}
	// 获取学生用户所属宿舍
	var dormId uint
	if utils.GetUserRoleId(c) == 3 {
		dormId = utils.GetUserDormId(c)
	}

	violateList, total, err := violateService.QueryStudentViolateList(limit, offset, arrSlice, dormId)
	if err != nil {
		response.FailWithMessage("查询学生违纪信息失败", c)
		return
	}
	// fmt.Println("+++++++发送数量为", total)
	// tt := total
	// fmt.Println("-----发送数量为", tt)
	response.OkWithDetailed(request.PageInfo{
		List:     violateList,
		PageSize: PageSize,
		Page:     Page,
		Total:    total,
	}, "成功", c)

}

var Vio_api = new(violate_info_api)
