package controller

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	// "back-end/model/apidorm"
	"back-end/model/apistudent"
	"back-end/utils"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type violate_info_api struct{}

// 插入
func (d *violate_info_api) CreateVioApi(c *gin.Context) {
	var violateList []apistudent.Violate_model
	err := c.ShouldBindJSON(&violateList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 给数据添加id
	// for _, v := range violateList {
		
	// 	//查询存在数据
	// 	var tempArr apistudent.Violate_model
	// 	query := global.Global_Db.Where("student_number=?", v.StudentNumber).First(&tempArr)
	// 	if query.Error != nil {
	// 		continue
	// 	}
	// 	if tempArr.StudentNumber == v.StudentNumber {
	// 		response.FailWithMessage("该学号:"+v.StudentNumber+"学生已存在", c)
	// 		return
	// 	}

	// }
	for i, _ := range violateList {
		uid := uuid.NewString()
		violateList[i].Id = uid
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
	var violateList []apistudent.Violate_model
	err := c.ShouldBindJSON(&violateList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range violateList {
		var student apistudent.Violate_model
		err2 := global.Global_Db.Model(&student).Where("id=?", value.Id).First(&student)
		if err2.Error != nil {
			response.FailWithMessage("删除学号为:"+value.StudentNumber+"违纪数据不存在", c)
			return
		}
	}
	for _, del := range violateList {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("删除学号为:"+del.StudentNumber+"违纪欣喜失败:", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *violate_info_api) UpdateVioApi(c *gin.Context) {
	var vio apistudent.Violate_model
	err := c.ShouldBindJSON(&vio)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	} 
	var tempStudent apistudent.Violate_model
	err2 := global.Global_Db.Where("id=? ", vio.Id).First(&tempStudent)
	if err2.Error != nil {
		response.FailWithMessage("更新的学号为:"+vio.StudentNumber+"数据不存在", c)
		return
	}
	// // 判断宿舍是否存在
	// var dorm apidorm.Dorm_api
	// dormErr := global.Global_Db.Where("dorm_number", vio.DormNumber).First(&dorm)
	// if dormErr.Error != nil {
	// 	response.FailWithMessage("该宿舍:"+vio.DormNumber+",不存在无法添加", c)
	// 	return
	// }
	result := global.Global_Db.Model(&vio).Where("id = ?", vio.Id).Updates(vio)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新学生:"+vio.StudentName+"失败", c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *violate_info_api) QueryVioApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var vio apistudent.Violate_model
	var violateList []apistudent.Violate_model
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
	count := global.Global_Db.Model(&vio).Where(condition).Count(&total).Error
	if count != nil {
		response.FailWithMessage("系统查寻数量错误", c)
		return
	}
	// 查寻数据
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&violateList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("系统查寻数据失败", c)
		return
	}
	for i, v := range violateList {
		time, err := time.Parse(time.RFC3339, v.RecordDate)
		if err != nil {
			response.FailWithMessage("系统解析时间错误错误", c)
			return
		}
		tt := time.Format("2006-01-02")
		violateList[i].RecordDate = tt
	}
	// fmt.Println("total",total,"数量为",len(violateList),violateList)
	response.OkWithDetailed(request.PageInfo{
		List:     violateList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)
	
}

var Vio_api = new(violate_info_api)
