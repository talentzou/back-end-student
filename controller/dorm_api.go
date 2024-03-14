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


var dormPages request.PageInfo

type dorm_api struct{}

// 插入
func (d *dorm_api) CreateDormApi(c *gin.Context) {
	var dormList []apidorm.Dorm_api
	fmt.Println("我进来了..........")
	err := c.ShouldBindJSON(&dormList)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	// 遍历添加的数据
	for key, v := range dormList {
		words := strings.Split(dormList[key].DormNumber, "-")
		if words[0] != dormList[key].FloorsName {
			fmt.Println(words[0] != dormList[key].FloorsName, "分割的单词", dormList[key].FloorsName, words)
			response.FailWithMessage("宿舍"+dormList[key].DormNumber+"开头前缀与宿舍楼不一致", c)
			return
		}
		// 查询存在数据
		var tempArr apidorm.Dorm_api
		query := global.Global_Db.Where("dorm_number=?", v.DormNumber).First(&tempArr)
		if query.Error != nil {
			response.FailWithMessage("系统查寻错误", c)
			return
		}
		if tempArr.DormNumber==v.DormNumber{
			response.FailWithMessage("该宿舍:"+v.DormNumber+"已存在", c)
			return
		}

	}
	// 给数据添加id
	for i, _ := range dormList {
		uid := uuid.NewString()
		dormList[i].Id = uid
	}
	// 添加数据
	result := global.Global_Db.Create(&dormList)
	if result.Error != nil {
		fmt.Println("宿舍错误")
		// 处理错误
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *dorm_api) DeleteDormApi(c *gin.Context) {
	var dormList []apidorm.Dorm_api
	err := c.ShouldBindJSON(&dormList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range dormList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			fmt.Println(err2.Error.Error())
			response.FailWithMessage("删除的数据不存在:", c)
			return
		}
	}
	for _, del := range dormList {
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
func (d *dorm_api) UpdateDormApi(c *gin.Context) {
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

func (d *dorm_api) QueryDormApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var dorm apidorm.Dorm_api
	var dormList []apidorm.Dorm_api
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
	result := global.Global_Db.Model(&dorm).Where(condition).Limit(limit).Offset(offset).Find(&dormList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("查寻错误", c)
		return
	}
	fmt.Println("数据为", len(dormList))
	response.OkWithDetailed(request.PageInfo{
		List:     dormList,
		Total:    total,
		PageSize: pages.PageSize,
		Page:     pages.Page,
	}, "成功", c)

}

var Dorm_api = new(dorm_api)
