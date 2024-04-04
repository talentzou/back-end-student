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
	"strings"

	"github.com/gin-gonic/gin"
)



type dorm_api struct{}

// 插入
func (d *dorm_api) CreateDormApi(c *gin.Context) {
	var dormList []dorm.Dorm
	fmt.Println("我进来了..........")
	err := c.ShouldBindJSON(&dormList)
	if err != nil {
		fmt.Println("参数错误为",err.Error())
		response.FailWithMessage("系统合并错误", c)
		return
	}
	// 遍历添加的数据
	for key, v := range dormList {
		words := strings.Split(dormList[key].DormNumber, "-")
		if words[0] != dormList[key].FloorsName {
			fmt.Println(words[0] != dormList[key].FloorsName , "分割的单词", dormList[key].FloorsName , words)
			response.FailWithMessage("宿舍"+dormList[key].DormNumber+"开头前缀与宿舍楼不一致", c)
			return
		}
		// 查询宿舍楼存在数据
		var tempFloor dorm.Floor
		query1 := global.Global_Db.Where("floors_name=?", v.FloorsName ).First(&tempFloor).Error 
		if query1!= nil {
			response.FailWithMessage("该宿舍楼"+v.FloorsName +"不存在,无法添加", c)
			return
		}
		// 查询宿舍存在数据
		var tempArr dorm.Dorm
		query := global.Global_Db.Where("dorm_number=?", v.DormNumber).First(&tempArr)
		if query.Error != nil {
			continue
		}
		if tempArr.DormNumber == v.DormNumber {
			response.FailWithMessage("该宿舍:"+v.DormNumber+"已存在", c)
			return
		}
		

	}
	// 添加数据
	result := global.Global_Db.Model(&dorm.Dorm{}).Create(&dormList).Error
	if result != nil {
		fmt.Println("宿舍错误",result)
		// 处理错误
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *dorm_api) DeleteDormApi(c *gin.Context) {
	var dormList []dorm.Dorm
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
		result := global.Global_Db.Where("id=?", del.Id).Delete(&del)
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
	var Dorm dorm.Dorm
	err := c.ShouldBindJSON(&Dorm)
	fmt.Println("宿舍更新信息为",Dorm)
	if err != nil {
		response.FailWithMessage("系统错误"+err.Error(), c)
		return
	}
	words := strings.Split(Dorm.DormNumber, "-")
	if words[0] != Dorm.FloorsName {
		response.FailWithMessage("宿舍与宿舍楼前缀不一致", c)
		return
	}
	result := global.Global_Db.Model(&Dorm).Where("id = ?", Dorm.Id).Updates(Dorm)
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
	var Dorm dorm.Dorm
	var dormList []dorm.Dorm
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
	// 分页
	offset = PageSize * (Page - 1)
	limit = PageSize
	fmt.Println(offset, limit)
	// 查寻数量
	count := global.Global_Db.Model(&Dorm).Where(condition).Count(&total).Error
	if count != nil {
		response.FailWithMessage("系统查寻数量错误", c)
		return
	}
	result := global.Global_Db.Model(&Dorm).Where(condition).Limit(limit).Offset(offset).Find(&dormList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("查寻数据错误", c)
		return
	}
	fmt.Println("数据为", len(dormList))
	response.OkWithDetailed(request.PageInfo{
		List:     dormList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var Dorm_api = new(dorm_api)
