package utils

import (
	"back-end/common/request"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)
var (
	SigningKey = []byte("hellozhj")
	token      *jwt.Token
	s          string
)

// 获取token
func GetToken(c *gin.Context) string {
	token, err := c.Cookie("x-token")
	if err != nil {
		token = c.Request.Header.Get("x-token")
	}
	return token
	// Authorization
}

// 创建jwt,token
func CreateToken(claims jwt.Claims) (string, error) {
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// 解析jwt,token
func ParseToken(tokenString string) (*request.CustomClaims, error) {
	// 解析成功

	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	// 判断错误类型
	if err != nil {
		// 判断错误类型..........
		err = errors.New("invalid token")
	}
	// 判断声明类型符合自定义类型
	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
		// 返回该声明
		return claims, nil
	} else {
		return nil, errors.New("Couldn't handle this token:")
	}

}
