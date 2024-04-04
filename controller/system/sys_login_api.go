package system

import (
	"back-end/common/request"
	sysReq "back-end/common/request"
	"back-end/common/response"
	sysRes "back-end/common/response"
	"back-end/global"
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
	fmt.Println("用户密码", user)
	if err != nil {
		sysRes.FailWithMessage("参数合并错误", c)
		return
	}
	var sysUser system.SysUser
	err2 := global.Global_Db.Model(&system.SysUser{}).Preload("SysAuthorityBtns").Where("user_name=?", user.Username).First(&sysUser).Error
	if err2 != nil {
		response.FailWithMessage("该用户不存在", c)
		return
	} else {
		fmt.Println("err2", sysUser.Authority == user.Authority)
		if sysUser.Authority != user.Authority {
			response.FailWithMessage("该用户不存在", c)
			return
		}
		ok := utils.BcryptCheck(user.Password, sysUser.Password)
		fmt.Println("密码比对情况", ok)
		if !ok {
			response.FailWithMessage("用户或密码错误", c)
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
	response.OkWithMessage("jwt设置失效成功", c)
}
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	// 初始jwt声明信息
	claims := utils.CreateClaims(request.BaseClaims{
		Id:          user.ID,
		Username:    user.UserName,
		NickName:    user.Nickname,
		AuthorityId: user.Authority,
	})
	// 创建token
	token, err3 := utils.CreateToken(claims)
	if err3 != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 设置响应头cookie,
	// fmt.Println("激活时间为时间为",claims.RegisteredClaims.NotBefore.Format("2006-01-02 15:04:05"))
	// fmt.Println("过期时间为",claims.RegisteredClaims.ExpiresAt.Format("2006-01-02 15:04:05"))
	// fmt.Println("运行时间为",time.Now().Format("2006-01-02 15:04:05"))
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	response.OkWithDetailed(system.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.UnixMilli(),
	}, "登录成功", c)

}
