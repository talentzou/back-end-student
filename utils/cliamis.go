package utils

import (
	"back-end/model/common/request"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

// 获取claim用户id
func GetUserID(c *gin.Context) uint {
	claims, exist := c.Get("claims")
	if exist {
		fmt.Println("找到claims")
		user, _ := claims.(*request.CustomClaims)
		return user.Id
	} else {
		cl, err := GetClaims(c)
		fmt.Println("找到claims")
		if err != nil {
			return 4040
		} else {
			return cl.Id
		}

	}

}
// 获取角色AuthorityId
func GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
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
