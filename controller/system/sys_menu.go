package system

import (
	"back-end/common/request"
	"back-end/common/response"
	"back-end/config"
	"back-end/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAsyncMenu(c *gin.Context) {
	id, b := c.Params.Get("authorityId")
	if !b {
		response.FailWithMessage("缺少角色参数标识", c)
		return
	}
	authority, err := strconv.Atoi(id)
	if err != nil {
		response.FailWithMessage("参数标识不是Number类型", c)
		return
	}
	fmt.Println("authority", authority)
	var routes []config.Common
	if authority == 1 {
		routes = global.Global_Font_End_Route.Admin
		response.OkWithDetailed(request.SysMenusResponse{
			Authority: authority,
			Menu:      routes,
		}, "获取管理员路由成功", c)
		return
	} else {
		routes = global.Global_Font_End_Route.Student
		response.OkWithDetailed(request.SysMenusResponse{
			Authority: authority,
			Menu:      routes,
		}, "获取学生路由成功", c)
		return
	}

}
