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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type  dorm_api interface {
// 	UpdateApi(c *gin.Context,class interface{})
// }

var rateList []apidorm.Rate_api
var ratePage request.PageInfo

type Dorm_rate_api struct{}

// 插入
func (d *Dorm_rate_api) CreateRateApi(c *gin.Context) {
	fmt.Println("我是评分.......")
	var tempArr []apidorm.Rate_api
	err := c.ShouldBindJSON(&rateList)
	if err != nil {
		fmt.Println("结构体错误", err.Error())
		response.FailWithMessage("系统合并错误", c)
		return
	}
	// //查寻存在数据
	query := global.Global_Db.Find(&tempArr)
	if query.Error != nil {
		response.FailWithMessage("系统查寻错误", c)
		return
	}
	// 给数据添加id
	for i, v := range rateList {
		words := strings.Split(v.DormNumber, "-")
		if v.FloorsName != words[0] {
			response.FailWithMessage("宿舍:"+v.DormNumber+"与宿舍楼:"+v.FloorsName+"前缀不一致", c)
			return
		}

		for t := range tempArr {
			// fmt.Println("日期：", rateList[i].RateDate, tempArr[t].RateDate, "布尔")
			// fmt.Println("宿舍：", rateList[i].DormNumber, tempArr[t].DormNumber)
			tt, err := time.Parse(time.RFC3339,tempArr[t].RateDate)
			if err != nil {
				response.FailWithMessage("系统解析事件错误", c)
				return
			}
			dd := tt.Format("2006-01-02")
			if rateList[i].RateDate == dd && rateList[i].DormNumber == tempArr[t].DormNumber {
				date := rateList[i].RateDate
				response.FailWithMessage(rateList[i].DormNumber+"宿舍的"+date+"的日期评分已存在", c)
				return
			}
		}
		uid := uuid.NewString()
		rateList[i].Id = uid
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
func (d *Dorm_rate_api) DeleteRateApi(c *gin.Context) {
	err := c.ShouldBindJSON(&rateList)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range rateList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage("删除时间为:"+value.RateDate+value.FloorsName+":"+value.DormNumber+"宿舍数据不存在", c)
			return
		}
	}
	for _, del := range rateList {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("删除时间为:"+del.RateDate+del.FloorsName+":"+del.DormNumber+"宿舍数据删除失败", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *Dorm_rate_api) UpdateRateApi(c *gin.Context) {
	var rate apidorm.Rate_api
	err := c.ShouldBindJSON(&rate)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	var tempRate apidorm.Rate_api 
	err2 := global.Global_Db.Where("id=?", rate.Id).First(&tempRate)
	if err2.Error != nil {
		response.FailWithMessage(rate.RateDate+":"+rate.DormNumber+":数据不存在:无法更新", c)
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

func (d *Dorm_rate_api) QueryRateApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var rate apidorm.Rate_api
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
	count := global.Global_Db.Model(&rate).Where(condition).Count(&total).Error
	if count != nil {
		fmt.Println("计算楼层数量错误")
		response.FailWithMessage("系统查询错误", c)
		return
	}
	// 查寻数据
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&rateList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("系统查寻失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     rateList,
		Total:    total,
		PageSize: pages.PageSize,
		Page:     pages.Page,
	}, "成功", c)

}

var Rate_api = new(Dorm_rate_api)
