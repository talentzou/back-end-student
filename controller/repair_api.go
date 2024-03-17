package controller

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/global"
	"back-end/model/apirepair"
	"back-end/utils"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


var repair apirepair.Repair_dorm

type repair_api_ struct{}

// 插入
func (d *repair_api_) CreateRepairApi(c *gin.Context) {
	var repairList []apirepair.Repair_dorm
	err := c.ShouldBindJSON(&repairList)
	if err != nil {
		fmt.Println("参数",err.Error())
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 给数据添加id
	for i, v := range repairList {
		words := strings.Split(v.DormNumber, "-")
		if words[0] != v.FloorsName {
			response.FailWithMessage("宿舍:"+v.DormNumber+"与宿舍楼:"+v.FloorsName+"前缀不一致", c)
			return
		}
		if len(v.Problems) == 0 {
			response.FailWithMessage("宿舍:"+v.DormNumber+"维修问题不能为空", c)
			return
		}
		// 判断宿舍是否是否存在
		// var dorm apidorm.Dorm_api
		// query := global.Global_Db.Where("dorm_number=?", v.DormNumber).First(&dorm)
		// if query.Error != nil {
		// 	response.FailWithMessage("该宿舍"+v.DormNumber+"不存在", c)
		// 	return
		// }
		uid := uuid.NewString()
		repairList[i].Id = uid
	}

	// 添加数据
	result := global.Global_Db.Create(&repairList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加维修失败", c)
		return
	}
	response.OkWithMessage("添加维修成功", c)
}

// 删除
func (d *repair_api_) DeleteRepairApi(c *gin.Context) {
	var repairList []apirepair.Repair_dorm
	err := c.ShouldBindJSON(&repairList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, v := range repairList {
		err2 := global.Global_Db.Where("id=?", v.Id).First(&v)
		if err2.Error != nil {
			response.FailWithMessage("删除的宿舍:"+v.DormNumber+"数据不存在:", c)
			return
		}
	}
	for _, del := range repairList {
		result := global.Global_Db.Where("id=?", del.Id).Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("宿舍:"+del.DormNumber+"数据删除失败", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *repair_api_) UpdateRepairApi(c *gin.Context) {
	err := c.ShouldBindJSON(&repair)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	var tempRepair apirepair.Repair_dorm
	err2 := global.Global_Db.Where("id=?", repair.Id).First(&tempRepair)
	if err2.Error != nil {
		response.FailWithMessage(repair.DormNumber+":数据不存在:无法更新", c)
		return
	}
	// 更改的宿舍与宿舍楼是否一致
	words := strings.Split(repair.DormNumber, "-")
	if words[0] != repair.FloorsName {
		response.FailWithMessage("宿舍:"+repair.DormNumber+"与宿舍楼:"+repair.FloorsName+"前缀不一致", c)
		return
	}
	result := global.Global_Db.Model(&repair).Where("id = ?", repair.Id).Updates(repair)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("宿舍:"+repair.DormNumber+"更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *repair_api_) QueryRepairApi(c *gin.Context) {
	var limit, offset int
	var total int64
	var repairList []apirepair.Repair_dorm
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
		fmt.Println("系统解析url错误")
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
	count := global.Global_Db.Model(&repair).Where(condition).Count(&total).Error
	if count != nil {
		// fmt.Println("查寻维修数量错误")
		response.FailWithMessage("系统查询维修数量错误", c)
		return
	}
	// 查寻数据
	result := global.Global_Db.Where(condition).Limit(limit).Offset(offset).Find(&repairList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("系统查寻数据失败", c)
		return
	}
	for i, v := range repairList {
		time, err := time.Parse(time.RFC3339, v.SubmitDate)
		if err != nil {
			response.FailWithMessage("系统解析时间错误错误", c)
			return
		}
		tt := time.Format("2006-01-02")
		repairList[i].SubmitDate = tt
	}
	response.OkWithDetailed(request.PageInfo{
		List:     repairList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var Repair_api = new(repair_api_)
