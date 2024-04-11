package system

import (
	sysRes "back-end/model/common/response"

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
	// for i := range roleList {
	// 	// IdMenuList, err2 := menuService.GETMenuIdList(roleList[i].Model.ID)
	// 	// if err2 != nil {
	// 	// 	continue
	// 	// }
	// 	// roleList[i].Menu_Btn_list.MenuListId = IdMenuList

	// 	// IdBtnList, err2 := btnService.GETBtnIdList(roleList[i].Model.ID)
	// 	// if err2 != nil {
	// 	// 	continue
	// 	// }
	// 	// roleList[i].Btn_List = IdBtnList

	// }
	sysRes.Ok("获取角色列表成功", roleList, c)
}
