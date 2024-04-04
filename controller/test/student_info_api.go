package test

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/test/dorm"
	"back-end/utils"
	"fmt"
	"net/url"
	"strconv"

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
	// 给数据添加id
	for _, v := range studInfoList {
		// 查询宿舍存在
		var Dorm dorm.Dorm
		dormErr := global.Global_Db.Where("dorm_number=?", v.DormNumber).First(&Dorm)
		if dormErr.Error != nil {
			response.FailWithMessage("该宿舍:"+v.DormNumber+"不存在无法添加", c)
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
	err2 := global.Global_Db.Where("id=?", stud.Id).First(&tempStudent)
	if err2.Error != nil {
		response.FailWithMessage("更新的学生:"+stud.StudentName+"数据不存在", c)
		return
	}
	// 判断宿舍是否存在
	var Dorm dorm.Dorm
	dormErr := global.Global_Db.Where("dorm_number", stud.DormNumber).First(&Dorm)
	if dormErr.Error != nil {
		response.FailWithMessage("该宿舍:"+stud.DormNumber+",不存在无法添加", c)
		return
	}
	result := global.Global_Db.Model(&stud).Where("id= ?", stud.Id).Updates(stud)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新学生:"+stud.StudentName+"失败", c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *student_info_api) QueryStudInfoApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var stud dorm.StudInfo
	var studInfoList []dorm.StudInfo
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
	fmt.Println(offset, limit)
	// 查寻数量
	count := global.Global_Db.Model(&stud).Where(condition).Count(&total).Error
	if count != nil {
		response.FailWithMessage("系统查寻数量错误", c)
		return
	}
	// 查寻数据
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&studInfoList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("系统查寻数据失败", c)
		return
	}
	// fmt.Println("total",total,"数量为",len(studInfoList),studInfoList)
	response.OkWithDetailed(request.PageInfo{
		List:     studInfoList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var StudInfo_api = new(student_info_api)
