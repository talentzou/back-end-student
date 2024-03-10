package controller

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/apidorm"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var floors []apidorm.Floors_api
var pages request.PageInfo

type Dorm_db_api struct{}

// 插入
func (d *Dorm_db_api) CreateApi(c *gin.Context) {
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
			response.Fail("该楼已存在", c)
			return
		}
	}
	// 添加数据
	result := global.Global_Db.Create(&floors)
	if result.Error != nil {
		// 处理错误
		response.Fail("添加失败", c)
		return
	}
	response.Ok("添加成功", c)
}

// 删除
func (d *Dorm_db_api) DeleteApi(c *gin.Context) {
	err := c.ShouldBindJSON(&floors)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range floors {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.Fail("删除的数据不存在:"+err2.Error.Error(), c)
			return
		}
	}
	for _, del := range floors {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.Fail("该数据不存在,无法删除:"+result.Error.Error(), c)
			return
		}
	}
	response.Ok("删除成功", c)
}

// 更新
func (d *Dorm_db_api) UpdateApi(c *gin.Context) {
	var floor apidorm.Floors_api
	err := c.ShouldBindJSON(&floor)
	if err != nil {
		response.FailWithMessage("系统错误"+err.Error(), c)
		return
	}
	
		result := global.Global_Db.Model(&floor).Where("id = ?",floor.Id).Updates(floor)
		if result.Error != nil {
			// 处理错误
			response.Fail("更新失败", c)
			return
	
	}
	response.Ok("更新成功", c)
}

// 查寻

func (d *Dorm_db_api) QueryApi(c *gin.Context) {
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
		key := ToCamelCase(index)
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
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&floors)
	if result.Error != nil {
		// 处理错误
		response.Fail("查寻失败", c)
		return
	}
	fmt.Println("数据为", floors)
	response.ResponseHTTP(200, floors, "成功", c)

}

// 将字符串命名为驼峰
func ToCamelCase(str string) string {
	words := make([]string, 10)
	re := regexp.MustCompile(`[a-z]+|[A-Z][^A-Z]*`)
	words = re.FindAllString(str, -1)
	// fmt.Println("单词为：", words)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	result := strings.Join(words, "_")
	// fmt.Println("单词最后为",result)
	return result
}

var Dorm_api = new(Dorm_db_api)
