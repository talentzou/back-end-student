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
	"strings"

	"github.com/gin-gonic/gin"
)

type dorm_rate_api struct{}

// 插入
func (d *dorm_rate_api) CreateRateApi(c *gin.Context) {
	var rateList []dorm.Rate
	err := c.ShouldBindJSON(&rateList)
	if err != nil {
		fmt.Println("结构体错误", err.Error())
		response.FailWithMessage("参数错误", c)
		return
	}

	for i, v := range rateList {
		// //查寻存在数据
		var tempArr []dorm.Rate
		err := global.Global_Db.Where("dorm_id=?", v.DormId).Find(&tempArr).Error
		if err != nil {
			response.FailWithMessage("系统查寻错误", c)
			return
		}

		for t := range tempArr {

			if rateList[i].RateDate == tempArr[t].RateDate && rateList[i].DormId == tempArr[t].DormId {
				response.FailWithMessage(tempArr[t].RateDate.Format("2006-01-02")+"的日期评分已存在", c)
				return
			}
		}

	}
	// 添加数据
	result := global.Global_Db.Create(&rateList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加评分失败", c)
		return
	}
	response.OkWithMessage("添加评分成功", c)
}

// 删除
func (d *dorm_rate_api) DeleteRateApi(c *gin.Context) {
	var rateList []dorm.Rate
	err := c.ShouldBindJSON(&rateList)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range rateList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage("删除时间为:"+value.RateDate.Format("2006-01-02")+"宿舍数据不存在", c)
			return
		}
	}
	for _, del := range rateList {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("删除时间为:"+del.RateDate.Format("2006-01-02")+"宿舍数据删除失败", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *dorm_rate_api) UpdateRateApi(c *gin.Context) {

	var rate dorm.Rate
	err := c.ShouldBindJSON(&rate)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	var tempRate dorm.Rate
	err2 := global.Global_Db.Where("id=?", rate.Id).First(&tempRate)
	if err2.Error != nil {
		response.FailWithMessage(rate.RateDate.Format("2006-01-02")+"数据不存在:无法更新", c)
		return
	}
	result := global.Global_Db.Model(&rate).Where("id = ?", rate.Id).Updates(rate)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新rate失败", c)
		return

	}
	response.OkWithMessage("更新rate成功", c)
}

// 查寻

func (d *dorm_rate_api) QueryRateApi(c *gin.Context) {
	var limit, offset int
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	if er1 != nil && er2 != nil {
		response.FailWithMessage("分页参数错误", c)
		return
	}
	offset = PageSize * (Page - 1)
	limit = PageSize
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
	// 变量
	var arrSlice interface{}

	mapLength := len(condition)
	fmt.Println("condition9999", mapLength,condition)
	if mapLength == 0 {
		arrSlice = nil
	} else {
		floor_dorm := condition["floor_dorm"].([]string)
		// fmt.Println("进来后++++",floor_dorm)
		words := strings.Split(floor_dorm[0], "-")
		arrSlice = words
	}
	// 分页数据

	// fmt.Println(offset, limit)
	rateList, total, err := rateService.QueryRate(limit, offset, arrSlice)
	if err != nil {
		response.FailWithMessage("查寻宿舍评分信息失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     rateList,
		Total:    total,
		PageSize: pages.PageSize,
		Page:     pages.Page,
	}, "成功", c)
}

var Rate_api = new(dorm_rate_api)
