package controller

import (
	"back-end/global"
	"back-end/model/apidorm"
	"fmt"
	"back-end/common/response"
	"github.com/gin-gonic/gin"
)
var floor []apidorm.Floors_api
type Dorm_db_api struct{}

func (d *Dorm_db_api) CreateApi(c *gin.Context) {
	result := global.Global_Db.Create(&floor) 
	if result.Error != nil {
        // 处理错误
        response.Fail("添加失败",c)
        return
    }
	response.Ok("添加成功",c)
}
func (d *Dorm_db_api) DeleteApi(c *gin.Context) {
	result:=global.Global_Db.Delete(&floor, []string{})
	if result.Error != nil {
        // 处理错误
        response.Fail("删除失败",c)
        return
    }
	response.Ok("删除成功",c)
}
func (d *Dorm_db_api) UpdateApi(c *gin.Context) {

}
func (d *Dorm_db_api) QueryApi(c *gin.Context) {
	
	result := global.Global_Db.Limit(10).Offset(0).Find(&floor)
	if result.Error != nil {
        // 处理错误
        response.Fail("查寻失败",c)
        return
    }
	fmt.Println("数据为", floor)
	response.ResponseHTTP(200,floor,"成功",c)

}

var Dorm_api = new(Dorm_db_api)
