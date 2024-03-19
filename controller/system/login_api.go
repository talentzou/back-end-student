package system

import (
	"back-end/common/request"
	sysReq "back-end/common/request"
	"back-end/common/response"
	sysRes "back-end/common/response"
	"back-end/global"
	"back-end/model/system"
	"back-end/utils"
	"time"

	"github.com/gin-gonic/gin"
)
var SystemApi BaseApi
type BaseApi struct{}


func (b *BaseApi) Login(c *gin.Context) {
	var user sysReq.Login
	err := c.ShouldBindJSON(&user)
	if err != nil {
		sysRes.FailWithMessage("参数合并错误", c)
		return
	}
	var sysUser system.SysUser
	err2 := global.Global_Db.Where("userName=?", user.Username).First(&sysUser)
	if err2 == nil {
		// 判断密码是否正确
		ok := utils.BcryptCheck(user.Password, sysUser.Password)
		if !ok {
			response.FailWithMessage("用户或密码错误", c)
			return
		}
		// 派发token
		b.TokenNext(c, sysUser)
		return
	}

}

func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	// 初始jwt声明信息
	claims := utils.CreateClaims(request.BaseClaims{
		Id:       user.ID,
		UUId:     user.UUID,
		Username: user.UserName,
		NickName: user.Nickname,
	})
	// 创建token
	token, err3 := utils.CreateToken(claims)
	if err3 != nil {
		response.FailWithMessage("获取token失败", c)
		return
	}
	// 设置响应头cookie,
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	response.OkWithDetailed(system.LoginResponse{
		User:  user,
		Token: token,
	}, "登录成功", c)

}
