package controller

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/apidorm"
	"back-end/utils"
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type  dorm_api interface {
// 	UpdateApi(c *gin.Context,class interface{})
// }

var dorms []apidorm.Dorm_api
var dormPages request.PageInfo

type Dorm_dorm_api struct{}

// 插入
func (d *Dorm_dorm_api) CreateDormApi(c *gin.Context) {
	fmt.Println("我进来了..........")
	err := c.ShouldBindJSON(&dorms)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
// 遍历添加的数据
	for key,v := range dorms {
		words := strings.Split(dorms[key].DormNumber, "-")
		if words[0] != dorms[key].FloorsName {
			fmt.Println(words[0] != dorms[key].FloorsName, "分割的单词", dorms[key].FloorsName, words)
			response.FailWithMessage("宿舍"+dorms[key].DormNumber+"开头前缀与宿舍楼不一致", c)
			return
		}
		// 查询存在数据
		var tempArr []apidorm.Dorm_api
		query := global.Global_Db.Where("floors_name=?", v.FloorsName).Find(&tempArr)
		if query.Error != nil {
			response.FailWithMessage("系统查寻错误", c)
			return
		}
		for i := range tempArr {
			if dorms[i].DormNumber == tempArr[i].DormNumber {
				response.FailWithMessage("该宿舍"+dorms[i].DormNumber+"已存在", c)
				return
			}
	
		}

	}

	// 给数据添加id
	for i, _ := range dorms {
		uid := uuid.NewString()
		dorms[i].Id = uid
	}
	// 添加数据
	result := global.Global_Db.Create(&dorms)
	if result.Error != nil {
		fmt.Println("计算宿舍数量错误")
		// 处理错误
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *Dorm_dorm_api) DeleteDormApi(c *gin.Context) {
	err := c.ShouldBindJSON(&dorms)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range dorms {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			fmt.Println(err2.Error.Error())
			response.FailWithMessage("删除的数据不存在:", c)
			return
		}
	}
	for _, del := range dorms {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			fmt.Println(result.Error.Error())
			response.FailWithMessage("数据删除失败:"+result.Error.Error(), c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *Dorm_dorm_api) UpdateDormApi(c *gin.Context) {
	var dorm apidorm.Dorm_api
	err := c.ShouldBindJSON(&dorm)
	if err != nil {
		response.FailWithMessage("系统错误"+err.Error(), c)
		return
	}
	words := strings.Split(dorm.DormNumber, "-")
	if words[0] != dorm.FloorsName {
		response.FailWithMessage("宿舍与宿舍楼前缀不一致", c)
		return
	}
	result := global.Global_Db.Model(&dorm).Where("id = ?", dorm.Id).Updates(dorm)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *Dorm_dorm_api) QueryDormApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var dorm apidorm.Dorm_api
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
	count := global.Global_Db.Where(condition).Count(&total).Error
	if count != nil {
		response.FailWithMessage(count.Error(), c)
		return
	}
	result := global.Global_Db.Model(&dorm).Where(condition).Limit(limit).Offset(offset).Find(&dorms)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("查寻错误", c)
		return
	}
	fmt.Println("数据为", len(dorms))
	response.OkWithDetailed(request.PageInfo{
		List:     dorms,
		Total:    total,
		PageSize: pages.PageSize,
		Page:     pages.Page,
	}, "成功", c)

}

var Dorm_api = new(Dorm_dorm_api)
