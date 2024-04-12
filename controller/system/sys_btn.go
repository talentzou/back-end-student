package system

import (
	// "back-end/global"
	sysReq "back-end/model/common/request"
	sysRes "back-end/model/common/response"
	"back-end/model/system"
	"back-end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type btnApi struct{}

var BtnApi btnApi

// 获取角色按钮
func (b *btnApi) GetSelfBtns(c *gin.Context) {
	fmt.Println("按钮角色id为", utils.GetUserRoleId(c))
	btns, err := btnService.GetBtnRoleTree(utils.GetUserRoleId(c))
	if err != nil {
		fmt.Println("出错了")
		sysRes.FailWithMessage("获取角色按钮失败", c)
		return
	}
	if btns == nil {
		btns = []system.SysBtn{}
	}
	sysRes.OkWithDetailed(sysRes.SysBtnsResponse{Btns: btns}, "获取按钮成功", c)
}

// 获取按钮菜单树
func (b *btnApi) GetBtnTree(c *gin.Context) {
	btnTree, err := btnService.GetBtnTreeMap()
	if err != nil {
		sysRes.FailWithMessage("获取按钮树失败", c)
	}
	sysRes.OkWithDetailed(sysRes.SysBtnsTree{BtnTree: btnTree}, "获取按钮树成功", c)
}

// 添加角色按钮关联
func (b *btnApi)RelateRoleAndBtn(c *gin.Context) {
	var role_btn sysReq.ReqRelateRoleAndBtn
	err := c.ShouldBindJSON(&role_btn)
	if err != nil {
		sysRes.FailWithMessage("参数错误", c)
		return
	}
	err = btnService.AddRelateBtn(role_btn.RoleId, role_btn.RoleBtnIdList)
	if err != nil {
		sysRes.FailWithMessage("添加失败", c)
	}
	sysRes.OkWithMessage("添加角色按钮关联成功", c)

}
// 删除角色按钮关联
func (b *btnApi)DeleteRelateRoleAndBtn(c *gin.Context) {
	var role_btn sysReq.ReqRelateRoleAndBtn
	err := c.ShouldBindJSON(&role_btn)
	if err != nil {
		sysRes.FailWithMessage("参数错误", c)
		return
	}
	err=btnService.DeleteRelateBtn( role_btn.RoleId,role_btn.RoleBtnIdList)
	if err != nil {
		sysRes.FailWithMessage("删除失败", c)
	}
	sysRes.OkWithMessage("删除角色按钮关联成功", c)
}
