package test

import (
	"back-end/global"
	"back-end/model/common/request"
	"back-end/model/common/response"
	"back-end/model/test/dorm"
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
		fmt.Println("添加参数错误",err.Error())
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	for _, v := range violateList {
		var Dorm dorm.Dorm
		dormErr := global.Global_Db.Where("id=? ", v.DormId).First(&Dorm)
		if dormErr.Error != nil {
			response.FailWithMessage("该宿舍不存在无法添加", c)
			return
		}
	}

	// 添加数据
	result := global.Global_Db.Create(&violateList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加学生违纪信息失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *violate_info_api) DeleteVioApi(c *gin.Context) {
	var violateList []student.StudentViolate
	
	err := c.ShouldBindJSON(&violateList)
	if err != nil {
		
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range violateList {
		var student student.StudentViolate
		err2 := global.Global_Db.Model(&student).Where("id=?", value.Id).First(&student)
		if err2.Error != nil {
			response.FailWithMessage("删除违纪为:"+value.Violate+",违纪数据不存在", c)
			return
		}
	}
	for _, del := range violateList {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("删除违纪为:"+del.Violate+"失败:", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *violate_info_api) UpdateVioApi(c *gin.Context) {
	var vio student.StudentViolate
	err := c.ShouldBindJSON(&vio)
	if err != nil {
		fmt.Println("参数错误",err.Error())
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	fmt.Println("参数为",vio)
	// fmt.Println("更新参数为",vio.DormId)
	// 判断数据是否存在
	var tempStudent student.StudentViolate
	err2 := global.Global_Db.Where("id=? ", vio.Id).First(&tempStudent)
	if err2.Error != nil {
		response.FailWithMessage("更新的学号为:"+vio. Violate+"数据不存在", c)
		return
	}
	fmt.Println("宿舍id参数为",vio.DormId)
	var Dorm dorm.Dorm
	dormErr := global.Global_Db.Where("id=? ", vio.DormId).First(&Dorm)
	if dormErr.Error != nil {
		response.FailWithMessage("该宿舍不存在无法更新", c)
		return
	}
	//

	result := global.Global_Db.Model(&vio).Where("id = ?", vio.Id).Updates(vio)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新学生:"+vio.Violate+"失败", c)
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

	arrSlice :=make([]string,3)
	mapLength := len(condition)
	fmt.Println("condition9999", mapLength,condition)
	if mapLength == 0 {
		fmt.Println("进来为空pppp")
		arrSlice = nil
	} else {
		fmt.Println("进来不为空++++++++++")
		if floorDorm, ok := condition["floor_dorm"].([]string); ok {
			words := strings.Split(floorDorm[0], "-")
			arrSlice[0]=words[0]
			arrSlice[1]=words[1]
		}
		if studentName,ok:= condition["student_name"].([]string);ok{
			arrSlice[2]=studentName[0]
		}
		
	}
	// 获取学生用户所属宿舍
	var dormId uint
	if utils.GetUserRoleId(c) == 3 {
		dormId = utils.GetUserDormId(c)
	}

	violateList,total,err:=studentService.QueryStudentViolateList(limit,offset,arrSlice,dormId)
	if err != nil {
		response.FailWithMessage("查询学生违纪信息失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     violateList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var Vio_api = new(violate_info_api)
