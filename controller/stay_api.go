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

var stayList []apidorm.Stay_api

type Dorm_stay_api struct{}

// 插入
func (d *Dorm_stay_api) CreateFloorApi(c *gin.Context) {
	fmt.Println("我是留宿.......")
	err := c.ShouldBindJSON(&stayList)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	fmt.Println("插入的数据", stayList)
	//
	for _, v := range stayList {
		var tempArr []apidorm.Stay_api
		query := global.Global_Db.Where("dorm_number=?", v.DormNumber).Find(&tempArr)
		if query.Error != nil {
			response.FailWithMessage("系统查询错误", c)
			return
		}
		for _, t := range tempArr {
			layout := "2006-01-02"
			time1, err1 := time.Parse(layout, v.StayTime.StartTime)
			if err1 != nil {
				fmt.Println("err11", err1.Error())
				response.FailWithMessage("系统解析时间错误错误", c)
				return
			}
			start := time1.Format("2006-01-02")
			time2, err2 := time.Parse(layout, v.StayTime.EndTime)
			if err2 != nil {
				fmt.Println("err22", err1.Error())
				response.FailWithMessage("系统解析时间错误错误", c)
				return
			}
			end := time2.Format("2006-01-02")
			// fmt.Println("宿舍", v.DormNumber, t.DormNumber, "bool", v.DormNumber == t.DormNumber)
			// fmt.Println("学生名", v.StudentName, t.StudentName, "bool", v.StudentName == t.StudentName)
			if v.DormNumber == t.DormNumber && v.StudentName == t.StudentName {

				s, err3 := time.Parse(time.RFC3339, t.StayTime.StartTime)
				if err3 != nil {
					fmt.Println("err33", err3.Error())
					response.FailWithMessage("系统解析时间错误错误", c)
					return
				}

				ss := s.Format("2006-01-02")
				e, err4 := time.Parse(time.RFC3339, t.StayTime.EndTime)
				if err4 != nil {
					fmt.Println("err44", err4.Error())
					response.FailWithMessage("系统解析时间错误错误", c)
					return
				}
				ee := e.Format("2006-01-02")
				// fmt.Println("数据库开始数据", t.StayTime.StartTime, "开始时间", start, ss, "bool", ss == start)
				// fmt.Println("数据库结束数据", t.StayTime.EndTime, "结束时间", end, ee, "bool", ee == end)
				if ss == start && ee == end {
					response.FailWithMessage("学生:"+v.StudentName+"留宿申请已存在", c)
					return
				}
			}

		}
	}

	// 给数据添加id
	for i, _ := range stayList {
		words := strings.Split(stayList[i].DormNumber, "-")
		if words[0] != stayList[i].FloorsName {
			fmt.Println(words[0] != stayList[i].FloorsName, "分割的单词", stayList[i].FloorsName, words)
			response.FailWithMessage("宿舍:"+stayList[i].DormNumber+"开头前缀与宿舍楼:"+stayList[i].FloorsName+"不一致", c)
			return
		}
		uid := uuid.NewString()
		stayList[i].Id = uid

	}
	// 添加数据
	result := global.Global_Db.Create(&stayList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *Dorm_stay_api) DeleteFloorApi(c *gin.Context) {
	err := c.ShouldBindJSON(&stayList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range stayList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage("删除为:"+value.DormNumber+",日期为:"+value.StayTime.StartTime+"至"+value.StayTime.EndTime+"的数据不存在:", c)
			return
		}
	}
	for _, del := range stayList {
		result := global.Global_Db.Where("id=?", del.Id).Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("删除为:"+del.DormNumber+",日期为:"+del.StayTime.StartTime+"至"+del.StayTime.EndTime+"无法删除:", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *Dorm_stay_api) UpdateFloorApi(c *gin.Context) {
	var stay apidorm.Stay_api
	err := c.ShouldBindJSON(&stay)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	var tempRate apidorm.Stay_api
	err2 := global.Global_Db.Where("id=?", stay.Id).First(&tempRate)
	if err2.Error != nil {
		response.FailWithMessage(stay.DormNumber+":数据不存在:无法更新", c)
		return
	}

	words := strings.Split(stay.DormNumber, "-")
	if words[0] != stay.FloorsName {
		response.FailWithMessage("宿舍与宿舍楼前缀不一致", c)
		return
	}
	result := global.Global_Db.Model(&stay).Where("id = ?", stay.Id).Updates(stay)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *Dorm_stay_api) QueryFloorApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var stay apidorm.Stay_api
	// 获取query
	rawUrl := c.Request.URL.String()
	u, er := url.Parse(rawUrl)
	if er != nil {
		fmt.Println("解析url错误")
		response.FailWithMessage("系统解析url错误", c)
		return
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
	count := global.Global_Db.Model(&stay).Where(condition).Count(&total).Error
	if count != nil {
		fmt.Println("查询留宿申请数量错误")
		response.FailWithMessage("系统查询错误", c)
		return
	}
	// 查寻数据
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&stayList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("系统查寻数据失败", c)
		return
	}
	// 处理时间格式
	for i, v := range stayList {
		time1, err1 := time.Parse(time.RFC3339, v.StayTime.StartTime)
		if err1 != nil {
			response.FailWithMessage("系统解析时间错误错误", c)
			return
		}
		start := time1.Format("2006-01-02")
		stayList[i].StayTime.StartTime = start
		time2, err2 := time.Parse(time.RFC3339, v.StayTime.EndTime)
		if err2 != nil {
			response.FailWithMessage("系统解析时间错误错误", c)
			return
		}
		end := time2.Format("2006-01-02")
		stayList[i].StayTime.StartTime = end
	}
	response.OkWithDetailed(request.PageInfo{
		List:     stayList,
		Total:    total,
		PageSize: pages.PageSize,
		Page:     pages.Page,
	}, "成功", c)

}

var Stay_api = new(Dorm_stay_api)
