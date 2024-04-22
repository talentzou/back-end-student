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

	for _, v := range studInfoList {
		var dorm_message dorm.Dorm
		var total int64
		err := global.Global_Db.Model(&dorm.Dorm{}).Preload("StudInfos").Where("id=?", v.DormId).First(&dorm_message).Count(&total).Error
		if err != nil {
			continue
		}
		fmt.Println("宿舍容量", dorm_message.Capacity)
		fmt.Println("宿舍容量目前为止", len(dorm_message.StudInfos))
		if len(dorm_message.StudInfos) >= dorm_message.Capacity {
			response.FailWithMessage("该宿舍:"+dorm_message.DormNumber+"容量已达到最大值", c)
			return
		}

		//查询存在数据
		var tempArr dorm.StudInfo
		query := global.Global_Db.Where("student_number=?", v.StudentNumber).First(&tempArr)
		if query.Error != nil {
			continue
		}
		if tempArr.StudentNumber == v.StudentNumber {
			response.FailWithMessage("该学号:"+v.StudentNumber+"学生已存在", c)
			return
		}

	}

	// 添加数据
	result := global.Global_Db.Create(&studInfoList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加学生信息失败", c)
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
	// 遍历查寻数据是否存在
	for _, value := range studInfoList {
		var student dorm.StudInfo
		err2 := global.Global_Db.Model(&student).Where("Id=?", value.Id).First(&student)
		if err2.Error != nil {
			response.FailWithMessage("删除学号为:"+value.StudentNumber+"数据不存在", c)
			return
		}
	}
	for _, del := range studInfoList {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("删除学生:"+del.StudentName+"失败:", c)
			return
		}
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
	var tempStudent dorm.StudInfo
	err = global.Global_Db.Where("id=?", stud.Id).First(&tempStudent).Error
	if err != nil {
		response.FailWithMessage("更新的学生:"+stud.StudentName+"数据不存在", c)
		return
	}
	err2 := global.Global_Db.Model(&dorm.StudInfo{}).Where("id= ?", stud.Id).Updates(stud)
	if err2.Error != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
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
