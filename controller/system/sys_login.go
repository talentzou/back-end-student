package system

import (
	"back-end/global"
	// "back-end/model/common/request"
	sysReq "back-end/model/common/request"
	// "back-end/model/common/response"
	sysRes "back-end/model/common/response"
	"back-end/model/system"
	"back-end/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var SystemApi BaseApi

type BaseApi struct{}

// 登录
func (b *BaseApi) Login(c *gin.Context) {
	var user sysReq.Login
	err := c.ShouldBindJSON(&user)
	// fmt.Println("用户密码", user)
	if err != nil {
		sysRes.FailWithMessage("参数合并错误", c)
		return
	}
	var sysUser system.SysUser
	// error1 := global.Global_Db.Model(&system.SysUser{}).Association("Role").Error
	// if error1 != nil {
	// 	fmt.Println("关联失败,里11111")
	// }

	// err2 := global.Global_Db.Model(&system.SysUser{}).Preload("Role").Where("user_name=?", user.Username).First(&sysUser).Error
	err2 := global.Global_Db.Model(&system.SysUser{}).Where("user_name=?", user.Username).First(&sysUser).Error

	if err2 != nil {
		sysRes.FailWithMessage("该用户不存在", c)
		return
	} else {
		// fmt.Println("err2", sysUser.Authority == user.Authority)
		if sysUser.RoleId != user.Authority {
			sysRes.FailWithMessage("该用户不存在", c)
			return
		}
		// if sysUser.Authority != user.Authority {
		// 	sysRes.FailWithMessage("该用户不存在", c)
		// 	return
		// }
		ok := utils.BcryptCheck(user.Password, sysUser.Password)
		fmt.Println("密码比对情况", ok)
		if !ok {
			sysRes.FailWithMessage("用户或密码错误", c)
			return
		}

		// 派发token
		fmt.Println("生成token")
		b.TokenNext(c, sysUser)
		return
	}
}

// 退出
func (j *BaseApi) Logout(c *gin.Context) {
	utils.ClearToken(c)
	sysRes.OkWithMessage("jwt设置失效成功", c)
}

// 创建token
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	// 初始jwt声明信息
	claims := utils.CreateClaims(sysReq.BaseClaims{
		Id:       user.ID,
		Username: user.UserName,
		NickName: user.Nickname,
		// AuthorityId: user.Authority,
		// RoleId:      user.Role.ID,
		RoleId: user.RoleId,
		DormId:   user.DormId,
	})
	// 创建token
	token, err3 := utils.CreateToken(claims)
	if err3 != nil {
		sysRes.FailWithMessage("获取token失败", c)
		return
	}
	// 设置响应头cookie,
	// fmt.Println("激活时间为时间为",claims.RegisteredClaims.NotBefore.Format("2006-01-02 15:04:05"))
	// fmt.Println("过期时间为",claims.RegisteredClaims.ExpiresAt.Format("2006-01-02 15:04:05"))
	// fmt.Println("运行时间为",time.Now().Format("2006-01-02 15:04:05"))
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	sysRes.OkWithDetailed(system.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.UnixMilli(),
	}, "登录成功", c)

}
