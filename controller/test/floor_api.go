package test

import (
	// "back-end/global"
	"back-end/model/common/request"
	"back-end/model/common/response"
	"back-end/model/test/dorm"
	// "back-end/utils"
	"fmt"
	// "net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pages request.PageInfo

type dorm_floor_api struct{}

var Floor_api = new(dorm_floor_api)

// 插入
func (d *dorm_floor_api) CreateFloorApi(c *gin.Context) {
	var floors []dorm.Floor
	// fmt.Println("我是楼.......")
	err := c.ShouldBindJSON(&floors)
	if err != nil {
		fmt.Println("参数错误为", err)
		response.FailWithMessage("添加的参数错误", c)
		return
	}
	fmt.Println("++++++执行到这")
	err = floorService.CreateFloor(&floors)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// //查询存在数据
	// var tempArr []dorm.Floor
	// query := global.Global_Db.Find(&tempArr)
	// if query.Error != nil {
	// 	response.FailWithMessage("系统查寻错误", c)
	// 	return
	// }

	// for i := range tempArr {
	// 	if floors[0].FloorsName == tempArr[i].FloorsName {
	// 		response.FailWithMessage("该楼已存在", c)
	// 		return
	// 	}
	// }
	// // 添加数据
	// result := global.Global_Db.Create(&floors)
	// if result.Error != nil {
	// 	// 处理错误
	// 	response.FailWithMessage("添加失败", c)
	// 	return
	// }
	response.OkWithMessage("添加成功", c)
}

// 删除
func (d *dorm_floor_api) DeleteFloorApi(c *gin.Context) {
	fmt.Println()
	var floors []dorm.Floor
	err := c.ShouldBindJSON(&floors)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("进入的数据为", floors)
	err = floorService.DeleteFloor(&floors)
	if err != nil {
		// fmt.Println("错误为", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	//遍历查寻数据是否存在
	// var floor dorm.Floor
	// for _, value := range floors {
	// 	fmt.Println("刪除的数据为", value.Id)
	// 	err2 := global.Global_Db.Where("id= ?", value.Id).First(&floor).Error
	// 	if err2 != nil {
	// 		fmt.Println("错误为", err)
	// 		response.FailWithMessage("删除的数据不存在:"+value.FloorsName, c)
	// 		return
	// 	}
	// }
	// for _, del := range floors {
	// 	result := global.Global_Db.Where("id= ?", del.Id).Delete(&del)
	// 	if result.Error != nil {
	// 		// 处理错误
	// 		response.FailWithMessage("该数据不存在,无法删除:"+del.FloorsName, c)
	// 		return
	// 	}
	// }
	response.OkWithMessage("删除成功", c)
}

// 更新
func (d *dorm_floor_api) UpdateFloorApi(c *gin.Context) {
	var floor dorm.Floor
	err := c.ShouldBindJSON(&floor)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	fmt.Println("参数为", floor)
	// var floorsList dorm.Floor
	// err = global.Global_Db.Where("floors_name = ?", floor.FloorsName).First(&floorsList).Error
	// // fmt.Println(err, "+++++++宿舍楼是否存在", floor.FloorsName)
	// if err == nil {
	// 	if floorsList.Id == floor.Id {
	// 	} else {
	// 		response.FailWithMessage("宿舍楼已存在", c)
	// 		return
	// 	}
	// }

	// result := global.Global_Db.Model(&dorm.Floor{}).Where("id= ?", floor.Id).Updates(floor)
	err = floorService.UpdateFloor(floor)
	if err != nil {
		// 处理错误
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// 查寻带参数

func (d *dorm_floor_api) QueryFloorApi(c *gin.Context) {
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

	var searchStr request.SearchParams
	err := c.ShouldBindJSON(&searchStr)
	if err != nil {
		response.FailWithMessage("搜索参数错误", c)
		return
	}
	fmt.Println("搜索参数为++++++", searchStr.QueryStr)

	floors, total, err := floorService.QueryFloor(limit, offset, searchStr.QueryStr)
	if err != nil {
		response.FailWithMessage("搜索数据不存在", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     floors,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)

}

// 不带参数查寻
func (d *dorm_floor_api) GetFloor(c *gin.Context) {
	var limit, offset int
	P, _ := c.Params.Get("Page")
	Size, _ := c.Params.Get("PageSize")
	PageSize, er1 := strconv.Atoi(Size)
	Page, er2 := strconv.Atoi(P)

	if er1 != nil && er2 != nil {
		response.FailWithMessage("分页参数错误", c)
		return
	}
	floors, total, err := floorService.QueryFloor(limit, offset, nil)
	if err != nil {
		response.FailWithMessage("搜索数据不存在", c)
		return
	}
	response.OkWithDetailed(request.PageInfo{
		List:     floors,
		Total:    total,
		PageSize: PageSize,
		Page:     Page,
	}, "成功", c)
}

// 获取宿舍楼与宿舍
func (d *dorm_floor_api) GetFloorWithDorm(c *gin.Context) {
	list, err := floorService.GetFloorDorm()
	if err != nil {
		response.FailWithMessage("宿舍楼详细数据不存在", c)
		return
	}
	response.OkWithDetailed(&response.FloorWithDormList{
		List: list,
	}, "获取宿舍楼详细信息成功", c)

}
