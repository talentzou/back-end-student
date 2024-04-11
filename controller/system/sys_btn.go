package system

import (
	// "back-end/global"
	"back-end/model/system"
	"back-end/utils"
	"fmt"
	// sysReq "back-end/model/common/request"
	sysRes "back-end/model/common/response"
	"github.com/gin-gonic/gin"
)

type btnApi struct{}

var BtnApi btnApi

// 获取角色按钮
func (b *btnApi) GetSelfBtns(c *gin.Context) {
	fmt.Println("按钮角色id为",utils.GetUserRoleId(c))
	btns, err := btnService.GetBtnTreeMap(utils.GetUserRoleId(c))
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
