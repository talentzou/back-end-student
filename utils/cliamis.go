package utils

import (
	"back-end/common/request"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var day int = 1000 * 60 * 60 * 24
var issuer string = "xiaozou"

// 初始声明信息
func CreateClaims(BaseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: BaseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),         // 签名生效时间),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), //生效时长
			Issuer:    issuer,                                            // 签名的发行者                                           //
		},
	}
	return claims
}

// 获取claim用户uuid
func GetUserUUID(c *gin.Context) string {
	claims, exist := c.Get("claims")
	if exist {
		fmt.Println("找到claims")
		user, _ := claims.(*request.CustomClaims)
		return user.UUId
	} else {
		cl, err := GetClaims(c)
		fmt.Println("找到claims")
		if err != nil {
			return uuid.New().String()
		} else {
			return cl.UUId
		}
		
	}

}

// 获取token声明信息
func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := GetToken(c)
	claims, err := ParseToken(token)
	if err != nil {
		fmt.Println("解析token失败")
	}
	return claims, err
}
