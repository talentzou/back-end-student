package utils

import (
	"back-end/global"
	"back-end/model/common/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

// 初始声明信息
func CreateClaims(BaseClaims request.BaseClaims) request.CustomClaims {
	expiresDuration, _ := strconv.Atoi(global.Global_Config.JWT.ExpiresTime)
	expires := time.Duration(expiresDuration)*time.Second
    // fmt.Println("转换后的时间段为：", expiresDuration)
	// fmt.Println("时间为",expires)
	claims := request.CustomClaims{
		BaseClaims: BaseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),                                        // 签名生效时间),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)), //生效时长
			Issuer:    global.Global_Config.JWT.Issuer,                                                  // 签名的发行者                                           //
		},
	}

	return claims
}

// 获取claim用户id
func GetUserID(c *gin.Context) uint {
	claims, exist := c.Get("claims")
	if exist {
		user, _ := claims.(*request.CustomClaims)
		// fmt.Println("找到claims的用户id", user.Id)
		return user.Id
	} else {
		cl, err := GetClaims(c)
		if err != nil {
			return 4040404040
		} else {
			fmt.Println("找到claims的用户id", cl.Id)
			return cl.Id
		}

	}

}

// 获取用户宿舍id
func GetUserDormId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.DormId
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.DormId
	}
}

// 获取角色Id
func GetUserRoleId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.RoleId
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.RoleId
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
