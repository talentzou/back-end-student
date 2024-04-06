package system

import (
	"back-end/config"
	"back-end/global"
	sysReq "back-end/model/common/request"
	sysRes "back-end/model/common/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAsyncMenu(c *gin.Context) {
	id, b := c.Params.Get("authorityId")
	if !b {
		sysRes.FailWithMessage("缺少角色参数标识", c)
		return
	}
	authority, err := strconv.Atoi(id)
	if err != nil {
		sysRes.FailWithMessage("参数标识不是Number类型", c)
		return
	}
	fmt.Println("authority", authority)
	var routes []config.Route
	if authority == 1 {
		routes = global.Global_Web_Route.Admin
		sysRes.OkWithDetailed(sysReq.SysMenusResponse{
			Authority: authority,
			Menu:      routes,
		}, "获取管理员路由成功", c)
		return
	} else if authority == 2 {
		routes = global.Global_Web_Route.Dorm
			sysRes.OkWithDetailed(sysReq.SysMenusResponse{
			Authority: authority,
			Menu:      routes,
		}, "获取宿舍管路由成功", c)
		return
	} else {
		routes = global.Global_Web_Route.Student
		sysRes.OkWithDetailed(sysReq.SysMenusResponse{
			Authority: authority,
			Menu:      routes,
		}, "获取学生路由成功", c)
		return
	}

}
