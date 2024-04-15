package test

import (
	"back-end/global"
	"back-end/model/common/request"
	"back-end/model/common/response"
	"back-end/model/test/dorm"
	"back-end/model/test/repair"
	"back-end/utils"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type repair_api_ struct{}

// 插入
func (d *repair_api_) CreateRepairApi(c *gin.Context) {
	var repairList []repair.Repair
	err := c.ShouldBindJSON(&repairList)
	fmt.Println("参数为", repairList)
	if err != nil {
		fmt.Println("参数", err.Error())
		response.FailWithMessage("添加参数错误", c)
		return
	}
	// 给数据添加id
	for _, v := range repairList {
		var tempDorm dorm.Dorm
		query := global.Global_Db.Where("id=? ", v.DormId).First(&tempDorm)
		if query.Error != nil {
			response.FailWithMessage("该宿舍不存在，无法添加", c)
			return
		}
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
	var repairList []repair.Repair
	err := c.ShouldBindJSON(&repairList)
	if err != nil {
		fmt.Println("维修参数错误", err.Error())
		response.FailWithMessage("参数错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, v := range repairList {
		err2 := global.Global_Db.Where("id=?", v.Id).First(&v)
		if err2.Error != nil {
			response.FailWithMessage("数据不存在:"+v.SubmitDate.Format("2006-01-02")+"至"+v.FinishDate.Format("2006-01-02"), c)
			return
		}
	}
	for _, del := range repairList {
		result := global.Global_Db.Where("id=?", del.Id).Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("数据删除"+del.SubmitDate.Format("2006-01-02")+"至"+del.FinishDate.Format("2006-01-02")+"失败", c)
			return
		}
	}
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *repair_api_) UpdateRepairApi(c *gin.Context) {
	var Repair repair.Repair
	err := c.ShouldBindJSON(&Repair)
	if err != nil {
		response.FailWithMessage("更新参数错误", c)
		return
	}
	var tempRepair repair.Repair
	err2 := global.Global_Db.Where("id=?", Repair.Id).First(&tempRepair)
	if err2.Error != nil {
		response.FailWithMessage("数据不存在:无法更新", c)
		return
	}
	if Repair.RepairStatus == "已完成" {
		Repair.FinishDate = time.Now()
	}
	fmt.Println("完成修改时间", Repair.RepairStatus)
	result := global.Global_Db.Model(&Repair).Where("id = ?", Repair.Id).Updates(Repair)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新"+Repair.SubmitDate.Format("2006-01-02")+"失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *repair_api_) QueryRepairApi(c *gin.Context) {
	var limit, offset int
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)
	if er1 != nil && er2 != nil {
		response.FailWithMessage("分页参数错误", c)
		return
	}
	// 分页数据
	offset = PageSize * (Page - 1)
	limit = PageSize
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

	arrSlice :=make([]string,3)
	mapLength := len(condition)
	fmt.Println("condition9999", mapLength,condition)
	if mapLength == 0 {
		fmt.Println("进来为空yyy")
		arrSlice = nil
	} else {
		fmt.Println("进来不为空+++yyy+++++++")
		if floorDorm, ok := condition["floor_dorm"].([]string); ok {
			words := strings.Split(floorDorm[0], "-")
			arrSlice[0]=words[0]
			arrSlice[1]=words[1]
		}
		if studentName,ok:= condition["repair_status"].([]string);ok{
			arrSlice[2]=studentName[0]
		}
		fmt.Println("kk",arrSlice[0],"1",arrSlice[1],"2",arrSlice[2])
		
	}

	repairList, total, err := repairService.QueryRepair(limit, offset, arrSlice)
	if err != nil {
		response.FailWithMessage("查询维修信息失败", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     repairList,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

var Repair_api = new(repair_api_)
