package test

import (
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

type student_info_api struct{}

// 插入
func (d *student_info_api) CreateStudInfoApi(c *gin.Context) {
	var studInfoList []dorm.StudInfo
	err := c.ShouldBindJSON(&studInfoList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	err = studentService.CreateStudentInfo(&studInfoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *student_info_api) DeleteStudInfoApi(c *gin.Context) {
	var studInfoList []dorm.StudInfo
	err := c.ShouldBindJSON(&studInfoList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	err = studentService.DeleteStudentInfo(&studInfoList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *student_info_api) UpdateStudInfoApi(c *gin.Context) {
	var stud dorm.StudInfo
	err := c.ShouldBindJSON(&stud)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	err = studentService.UpdateStudentInfo(stud)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *student_info_api) QueryStudInfoApi(c *gin.Context) {
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
		response.FailWithMessage("系统解析url错误", c)
		return
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

	studInfoList, total, err := studentService.QueryStudentInfoList(limit, offset, condition)
	if err != nil {
		response.FailWithMessage("查询学生信息失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     studInfoList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var StudInfo_api = new(student_info_api)
