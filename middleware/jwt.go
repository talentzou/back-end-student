package middleware

import (
	"back-end/common/response"
	"back-end/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头token
		token := utils.GetToken(c)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		// 返回声明信息
		claims, err := utils.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.ErrTokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		// 用于将控制权传递给下一个中间件或处理函数
		c.Next()
		c.Set("claims", claims)
	}
}
