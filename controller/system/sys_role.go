package system

import (
	sysRes "back-end/model/common/response"
	"back-end/model/system"
	"fmt"

	"github.com/gin-gonic/gin"
)

type roleApi struct{}

var RoleApi roleApi

// 获取角色列表
func (r *roleApi) GetRoles(c *gin.Context) {
	roleList, err := roleService.GetRoleList()
	if err != nil {
		sysRes.FailWithMessage("查找角色列表失败", c)
	}
	sysRes.Ok("获取角色列表成功", roleList, c)
}

// 删除
func (r *roleApi) DeleteRoles(c *gin.Context) {
	var roles system.Role
	err := c.ShouldBindJSON(&roles)
	fmt.Println("参数为98989", roles.Model.ID )
	if err != nil {
		sysRes.FailWithMessage("参数错误", c)
		return
	}
	if roles.Model.ID <= 4 {
		sysRes.FailWithMessage("不能删除系统基础角色", c)
		return
	}
	err = roleService.DeleteRolesList(roles)
	if err != nil {
		sysRes.FailWithMessage("删除角色失败", c)
		return
	}
	sysRes.OkWithMessage("删除角色成功", c)
}

// 添加
func (r *roleApi) CreateRoles(c *gin.Context) {
	var roles system.Role
	err := c.ShouldBindJSON(&roles)
	if err != nil {
		sysRes.FailWithMessage("参数错误", c)
		return
	}
	err = roleService.CreateRolesList(roles)
	if err != nil {
		sysRes.FailWithMessage("添加角色失败", c)
		return
	}
	sysRes.OkWithMessage("添加角色成功", c)
}
// 获取角色信息
func (r *roleApi) GetRoleMessage(c *gin.Context) {
	roleList, err := roleService.GetRolesMsg()
	if err != nil {
		sysRes.FailWithMessage("查找角色列表失败", c)
	}
	sysRes.Ok("获取角色列表成功", roleList, c)
}




