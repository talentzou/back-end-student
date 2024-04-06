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

type dorm_bed_api struct{}

// 插入
func (d *dorm_bed_api) CreateBedApi(c *gin.Context) {
	var bedList []dorm.Bed
	fmt.Println("我是床位.......")
	err := c.ShouldBindJSON(&bedList)
	if err != nil {
		response.FailWithMessage("系统合并参数错误", c)
		return
	}
	// 遍历body数据
	for i, v := range bedList {

		var tempArr []dorm.Bed
		er := global.Global_Db.Where(&dorm.Bed{DormNumber: v.DormNumber}).Find(&tempArr).Error
		if er != nil {
			fmt.Println("该宿舍没有数据")
			// response.FailWithMessage("该宿舍"+v.DormNumber+"不存在", c)
			// return
		}
		if len(tempArr) == 0 {
			continue
		}
		// 找到宿舍已存在容量
		length := len(tempArr)

		//  找到宿舍拥有容量
		var Dorm dorm.Dorm
		words := strings.Split(v.DormNumber, "-")
		dormErr := global.Global_Db.Model(&dorm.Dorm{}).Where(&dorm.Dorm{FloorsName: words[0], DormNumber: words[1]}).First(&Dorm).Error
		if dormErr != nil {
			response.FailWithMessage("宿舍"+v.DormNumber+"不存在,先添加该宿舍", c)
			return
		}

		// 判断宿舍容量是否超出
		if length >= Dorm.DormCapacity {
			response.FailWithMessage("超出"+bedList[i].DormNumber+"宿舍最大容量MAX", c)
			return
		}
		// 判断床位
		fmt.Println("床位数据为")
		for _, v := range tempArr {
			// 判断床位存在
			// fmt.Println("我进入床位数据：", b)
			if bedList[i].BedNumber == v.BedNumber {
				response.FailWithMessage("宿舍"+bedList[i].DormNumber+":"+strconv.Itoa(v.BedNumber)+"号床位，已有人", c)
				return
			}
		}

	}
	// for i2, val := range bedList {
	// 	// 给数据添加id
	// 	var Dorm dorm.Dorm
	// 	// 前面宿舍数据没有，再次判断宿舍是否存在
	// 	dormErr := global.Global_Db.Where(&dorm.Dorm{DormNumber: bedList[i2].DormNumber}).First(&Dorm).Error
	// 	if dormErr != nil {
	// 		response.FailWithMessage("宿舍"+val.DormNumber+"不存在,先添加该宿舍", c)
	// 		return
	// 	}
	// 	if bedList[i2].BedNumber > Dorm.DormCapacity {
	// 		response.FailWithMessage("床位编号："+strconv.Itoa(bedList[i2].BedNumber)+"不能大于宿舍容量", c)
	// 		return
	// 	}

	// }
	// 添加数据
	result := global.Global_Db.Create(&bedList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("添加床位失败", c)
		return
	}
	response.OkWithMessage("添加床位成功", c)
}

// 删除
func (d *dorm_bed_api) DeleteBedApi(c *gin.Context) {
	var bedList []dorm.Bed
	err := c.ShouldBindJSON(&bedList)
	if err != nil {
		response.FailWithMessage("参数合并错误", c)
		return
	}
	// 遍历查寻数据是否存在
	for _, value := range bedList {
		err2 := global.Global_Db.Where("id=?", value.Id).First(&value)
		if err2.Error != nil {
			response.FailWithMessage("删除的数据:"+value.DormNumber+":"+strconv.Itoa(value.BedNumber)+"号床不存在:", c)
			return
		}
	}
	for _, del := range bedList {
		result := global.Global_Db.Delete(&del)
		if result.Error != nil {
			// 处理错误
			response.FailWithMessage("数据"+del.DormNumber+":"+strconv.Itoa(del.BedNumber)+"删除失败:", c)
			return
		}
	}
	response.OkWithMessage("删除床位成功", c)
}

// 更新
func (d *dorm_bed_api) UpdateBedApi(c *gin.Context) {
	var bed dorm.Bed
	err := c.ShouldBindJSON(&bed)
	if err != nil {
		response.FailWithMessage("系统合并错误", c)
		return
	}
	// 判断数据是否存在
	var tempBed dorm.Bed
	err2 := global.Global_Db.Where("id=?", bed.Id).First(&tempBed)
	if err2.Error != nil {
		response.FailWithMessage(bed.DormNumber+":"+strconv.Itoa(bed.BedNumber)+"号床数据不存在:无法更新", c)
		return
	}
	// 判断更新的宿舍是否存在
	var Dorm dorm.Dorm
	words := strings.Split(bed.DormNumber, "-")
	dormErr := global.Global_Db.Model(&dorm.Dorm{}).Where(&dorm.Dorm{FloorsName: words[0], DormNumber: words[1]}).First(&Dorm).Error
	if dormErr != nil {
		response.FailWithMessage("宿舍"+bed.DormNumber+"不存在,先添加该宿舍", c)
		return
	}
	// 查找属于该宿舍数据
	var tempArr []dorm.Bed
	er := global.Global_Db.Where(&dorm.Bed{DormNumber: bed.DormNumber}).Find(&tempArr).Error
	if er != nil {
		response.FailWithMessage("系统查询错误", c)
		return
	}
	for _, v := range tempArr {
		if bed.BedStatus == "没人" {
			bed.StudentName = "无"
		}
		// 判断是否是原来的数据
		if v.Id == bed.Id {
			continue
		}
		if bed.BedNumber == v.BedNumber {
			response.FailWithMessage("该床位已有人", c)
			return
		}
	}
	// 更新
	result := global.Global_Db.Model(&bed).Where("id = ?", bed.Id).Updates(bed)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("更新失败", c)
		return

	}
	response.OkWithMessage("更新成功", c)
}

// 查寻

func (d *dorm_bed_api) QueryBedApi(c *gin.Context) {
	var bedList []dorm.Bed
	// 获取query
	rawUrl := c.Request.URL.String()
	u, er := url.Parse(rawUrl)
	if er != nil {
		fmt.Println("解析url错误")
	}
	queryParams := u.Query()
	// fmt.Println("查寻字符串参数", queryParams)
	condition := make(map[string]interface{})
	for index, value := range queryParams {
		key := utils.ToCamelCase(index)
		condition[key] = value
	}
	fmt.Println("condition", condition)
	// 查寻
	result := global.Global_Db.Where(condition).Find(&bedList)
	if result.Error != nil {
		// 处理错误
		response.FailWithMessage("查寻数据失败", c)
		return
	}
	total := len(bedList)
	fmt.Println("HFDHDHDHHD")
	response.OkWithDetailed(request.PageInfo{
		List:  bedList,
		Total: int64(total),
	}, "成功", c)
}

var Bed_api = new(dorm_bed_api)
