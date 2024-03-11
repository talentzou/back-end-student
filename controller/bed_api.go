package controller

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/apidorm"
	"fmt"
	"net/url"
     "back-end/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type  dorm_api interface {
// 	UpdateApi(c *gin.Context,class interface{})
// }

var beds []apidorm.Bed_api
var bedPages request.PageInfo

type Dorm_bed_api struct{}

// 插入
func (d *Dorm_bed_api) CreateBedApi(c *gin.Context) {
	fmt.Println("我是床位.......")
	var tempArr []apidorm.Bed_api
	err := c.ShouldBindJSON(&beds)
	if err != nil {
		response.FailWithMessage("系统合并参数错误"+err.Error(), c)
		return
	}
	//
	query := global.Global_Db.Find(&tempArr)
	if query.Error != nil {
		response.FailWithMessage("系统查寻错误", c)
		return
	}
	// 给数据添加id
	for i, _ := range beds {
		uid := uuid.NewString()
		beds[i].Id = uid
	}
	for i := range tempArr {
		if beds[0].DormNumber == tempArr[i].DormNumber {
			response.FailWithMessage("该宿舍已存在", c)
			return
		}
	}
	// 添加数据
	result := global.Global_Db.Create(&beds)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加失败", c)
		return
	}
	response.Ok("添加成功", c)
}

// 删除
func (d *Dorm_bed_api) DeleteBedApi(c *gin.Context) {
	err := c.ShouldBindJSON(&beds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range beds {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage("删除的数据不存在:"+err2.Error.Error(), c)
			return
		}
	}
	for _, del := range beds {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("该数据不存在,无法删除:"+result.Error.Error(), c)
			return
		}
	}
	response.Ok("删除成功", c)
}

// 更新
func (d *Dorm_bed_api) UpdateBedApi(c *gin.Context) {
	var bed apidorm.Bed_api
	err := c.ShouldBindJSON(&bed)
	if err != nil {
		response.FailWithMessage("系统错误"+err.Error(), c)
		return
	}
	result := global.Global_Db.Model(&bed).Where("id = ?", bed.Id).Updates(bed)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
		return

	}
	response.Ok("更新成功", c)
}

// 查寻

func (d *Dorm_bed_api) QueryBedApi(c *gin.Context) {
	var limit, offset int
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
	offset = pages.PageSize * pages.Page
	limit = pages.PageSize
	fmt.Println(offset, limit)
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&beds)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("查寻失败", c)
		return
	}
	fmt.Println("数据为", beds)
	response.ResponseHTTP(200, beds, "成功", c)

}


var Bed_api = new(Dorm_bed_api)
