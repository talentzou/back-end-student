package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysBtn struct{}

func (s *SysBtn) BtnRoute(router *gin.RouterGroup) {
	btn := router.Group("btn")
	{
		btn.GET("getBtnTree", system.BtnApi.GetBtnTree)                            //获取按钮树
		btn.GET("getSelfBtn", system.BtnApi.GetSelfBtns)                           // 分页获取用户菜单
		btn.POST("addRelateRoleAndBtn", system.BtnApi.RelateRoleAndBtn)            //添加角色按钮关联
		btn.DELETE("deleteRelateRoleAndBtn", system.BtnApi.DeleteRelateRoleAndBtn) //删除角色按钮关联
	}

}
