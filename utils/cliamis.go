package utils

import (
	"back-end/common/request"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var day int = 1000 * 60 * 60 * 24
var issuer string = "xiaozou"

// 初始声明信息
func CreateClaims(BaseClaims request.BaseClaims) request.CustomClaims {
	claims := request.CustomClaims{
		BaseClaims: BaseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),              // 签名生效时间),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour*2)), //生效时长
			Issuer:    issuer,                                                 // 签名的发行者                                           //
		},
	}
	return claims
}
