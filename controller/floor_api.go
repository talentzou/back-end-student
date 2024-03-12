package controller

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/apidorm"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/url"
)

// type  dorm_api interface {
// 	UpdateApi(c *gin.Context,class interface{})
// }

var floors []apidorm.Floors_api
var pages request.PageInfo

type Dorm_floor_api struct{}

// 插入
func (d *Dorm_floor_api) CreateFloorApi(c *gin.Context) {
	fmt.Println("我是楼.......")
	var tempArr []apidorm.Floors_api
	err := c.ShouldBindJSON(&floors)
	if err != nil {
		response.FailWithMessage("系统错误", c)
		return
	}
	//
	query := global.Global_Db.Find(&tempArr)
	if query.Error != nil {
		response.FailWithMessage("系统错误", c)
		return
	}
	// 给数据添加id
	for i, _ := range floors {
		uid := uuid.NewString()
		floors[i].Id = uid
	}
	for i := range tempArr {
		if floors[0].FloorsName == tempArr[i].FloorsName {
			response.FailWithMessage("该楼已存在", c)
			return
		}
	}
	// 添加数据
	result := global.Global_Db.Create(&floors)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *Dorm_floor_api) DeleteFloorApi(c *gin.Context) {
	err := c.ShouldBindJSON(&floors)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range floors {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage("删除的数据不存在:"+err2.Error.Error(), c)
			return
		}
	}
	for _, del := range floors {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("该数据不存在,无法删除:"+result.Error.Error(), c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *Dorm_floor_api) UpdateFloorApi(c *gin.Context) {
	var floor apidorm.Floors_api
	err := c.ShouldBindJSON(&floor)
	if err != nil {
		response.FailWithMessage("系统错误"+err.Error(), c)
		return
	}
	result := global.Global_Db.Model(&floor).Where("id = ?", floor.Id).Updates(floor)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *Dorm_floor_api) QueryFloorApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var floor apidorm.Floors_api
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
	count := global.Global_Db.Model(&floor).Where(condition).Count(&total).Error
	if count != nil {
		fmt.Println("计算楼层数量错误")
		response.FailWithMessage("系统查询错误", c)
		return
	}
	// 查寻数据
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&floors)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("系统查寻失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     floors,
		Total:    total,
		PageSize: pages.PageSize,
		Page:     pages.Page,
	}, "成功", c)

}

var Floor_api = new(Dorm_floor_api)
