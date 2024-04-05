package test

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/test/dorm"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"

)

type dorm_stay_api struct{}

// 插入
func (d *dorm_stay_api) CreateFloorApi(c *gin.Context) {
	var stayList []dorm.Stay
	fmt.Println("我是留宿.......")
	err := c.ShouldBindJSON(&stayList)
	if err != nil {
		fmt.Println("错误为", err.Error())
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	fmt.Println("插入的数据", stayList)
	//查寻数据存在
	for _, v := range stayList {
		// 查询宿舍存在数据
		var tempDorm dorm.Dorm
		queryDorm := global.Global_Db.Where("dorm_number=? AND floors_name=?", v.DormNumber, v.FloorsName).First(&tempDorm)
		if queryDorm.Error != nil {
			response.FailWithMessage("该宿舍"+v.FloorsName+"-"+v.DormNumber+"不存在,无法添加", c)
			return
		}
		var tempStay []dorm.Stay
		query := global.Global_Db.Where("dorm_number=? AND floors_name=?", v.DormNumber, v.FloorsName).Find(&tempStay)
		if query.Error != nil {
			response.FailWithMessage("该宿舍"+v.FloorsName+"-"+v.DormNumber+"不存在", c)
			return
		}
		for _, t := range tempStay {
			// 判断是宿舍与学生
			if v.StudentName == t.StudentName {
				//   判断日期
				if t.StayTime.StartTime == v.StayTime.StartTime && t.StayTime.EndTime == v.StayTime.EndTime {
					response.FailWithMessage("学生:"+v.StudentName+"留宿申请已存在", c)
					return
				}
			}

		}
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
func (d *dorm_stay_api) DeleteFloorApi(c *gin.Context) {
	var stayList []dorm.Stay
	err := c.ShouldBindJSON(&stayList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range stayList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage("删除为:"+value.DormNumber+",日期为:"+value.StayTime.StartTime.Format("2006-01-02")+"至"+value.StayTime.EndTime.Format("2006-01-02")+"的数据不存在:", c)
			return
		}
	}
	for _, del := range stayList {
		result := global.Global_Db.Where("id=?", del.Id).Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("删除为:"+del.DormNumber+",日期为:"+del.StayTime.StartTime.Format("2006-01-02")+"至"+del.StayTime.EndTime.Format("2006-01-02")+"无法删除:", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *dorm_stay_api) UpdateFloorApi(c *gin.Context) {
	var stay dorm.Stay
	err := c.ShouldBindJSON(&stay)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	// 查寻数据是否存在
	var tempRate dorm.Stay
	err2 := global.Global_Db.Where("id=?", stay.Id).First(&tempRate)
	if err2.Error != nil {
		response.FailWithMessage(stay.DormNumber+":数据不存在:无法更新", c)
		return
	}
	// 查寻宿舍
	var tempDorm dorm.Dorm
	queryDorm := global.Global_Db.Where("dorm_number=? AND floors_name=?", stay.DormNumber, stay.FloorsName).First(&tempDorm)
	if queryDorm.Error != nil {
		response.FailWithMessage("该宿舍"+stay.FloorsName+"-"+stay.DormNumber+"不存在,无法更新", c)
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

func (d *dorm_stay_api) QueryFloorApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var stay dorm.Stay
	var stayList []dorm.Stay

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
	// 分页数据
	offset = PageSize * (Page - 1)
	limit = PageSize
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
	// // 处理时间格式
	// for i, v := range stayList {
	// 	time1, err1 := time.Parse(time.RFC3339, v.StayTime.StartTime)
	// 	if err1 != nil {
	// 		response.FailWithMessage("系统解析时间错误错误", c)
	// 		return
	// 	}
	// 	start := time1.Format("2006-01-02")
	// 	stayList[i].StayTime.StartTime = start
	// 	time2, err2 := time.Parse(time.RFC3339, v.StayTime.EndTime)
	// 	if err2 != nil {
	// 		response.FailWithMessage("系统解析时间错误错误", c)
	// 		return
	// 	}
	// 	end := time2.Format("2006-01-02")
	// 	stayList[i].StayTime.EndTime = end
	// }
	response.OkWithDetailed(request.PageInfo{
		List:     stayList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var Stay_api = new(dorm_stay_api)
