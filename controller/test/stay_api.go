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

	"github.com/gin-gonic/gin"
)

type dorm_stay_api struct{}

// 插入
func (d *dorm_stay_api) CreateStayApi(c *gin.Context) {
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
		err := global.Global_Db.Where("id=?", v.DormId).First(&tempDorm).Error
		if err != nil {

			response.FailWithMessage("该宿舍不存在,无法添加", c)
			return
		}
		var tempStay []dorm.Stay
		err = global.Global_Db.Where("dorm_id=? ", v.DormId).Find(&tempStay).Error
		if err != nil {
			continue
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
func (d *dorm_stay_api) DeleteStayApi(c *gin.Context) {
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
			response.FailWithMessage("删除日期为:"+value.StayTime.StartTime.Format("2006-01-02")+"至"+value.StayTime.EndTime.Format("2006-01-02")+"的数据不存在:", c)
			return
		}
		if value.Opinions == "不同意" || value.Opinions == "同意" {
			if utils.GetUserRoleId(c) > 2 {
				response.FailWithMessage("状态发生改变，权限不足，无法删除", c)
				return
			}
		}
	}
	for _, del := range stayList {
		result := global.Global_Db.Where("id=?", del.Id).Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("删除日期为:"+del.StayTime.StartTime.Format("2006-01-02")+"至"+del.StayTime.EndTime.Format("2006-01-02")+"无法删除:", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *dorm_stay_api) UpdateStayApi(c *gin.Context) {
	var stay dorm.Stay
	err := c.ShouldBindJSON(&stay)
	if err != nil {
		fmt.Println("残心参数", err.Error())
		response.FailWithMessage("参数错误", c)
		return
	}
	// 查寻数据是否存在
	var tempRate dorm.Stay
	err2 := global.Global_Db.Where("id=?", stay.Id).First(&tempRate)
	if err2.Error != nil {
		response.FailWithMessage(stay.StayCause+":数据不存在:无法更新", c)
		return
	}
	// 查寻宿舍
	var tempDorm dorm.Dorm
	queryDorm := global.Global_Db.Where("id=?", stay.DormId).First(&tempDorm)
	if queryDorm.Error != nil {
		response.FailWithMessage("该宿舍不存在,无法更新", c)
		return
	}

	if stay.Opinions == "不同意" || stay.Opinions == "同意" {
		if utils.GetUserRoleId(c) > 2 {
			response.FailWithMessage("状态发生改变，权限不足，无法更新", c)
			return
		}
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

func (d *dorm_stay_api) QueryStayApi(c *gin.Context) {
	var limit, offset int
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	if er1 != nil && er2 != nil {
		response.FailWithMessage("分页参数错误", c)
		return
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
	fmt.Println("condition9999999", condition)
	// 分页数据
	offset = PageSize * (Page - 1)
	limit = PageSize
	// fmt.Println(offset, limit)
	fmt.Println("ttttttttt+++++++ttttttttt,执行到这里")

	// 获取学生用户所属宿舍
	var dormId uint
	if utils.GetUserRoleId(c) == 3 {
		dormId = utils.GetUserDormId(c)
	}

	stayList, total, err := stayService.QueryStay(limit, offset, condition, dormId)
	if err != nil {
		response.FailWithMessage("查询留宿申请信息失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     stayList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var Stay_api = new(dorm_stay_api)
