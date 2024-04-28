package utils

import (
	"back-end/model/common/request"
	"errors"
	// "fmt"
	"net"
	"back-end/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

// error strings should not be capitalized (ST1005)
var (
	// SigningKey = []byte("hellozhj")
	SigningKey = []byte(global.Global_Config.JWT.SigningKey)
	token      *jwt.Token
)

// 获取token
func GetToken(c *gin.Context) string {
	token, err := c.Cookie("x-token")
	if err != nil {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

// 设置响应头cookie
func SetToken(c *gin.Context, token string, maxAge int) {
	// fmt.Println("maxAge", maxAge)
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

// token过期，清除token
func ClearToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

// 创建jwt,token
func CreateToken(claims jwt.Claims) (string, error) {
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(SigningKey)
	return t, err
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
		if errors.Is(err, jwt.ErrTokenExpired) {
			// token过期
			return nil, ErrTokenExpired
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			// token被串改
			return nil, ErrTokenMalformed
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			// token还未生效
			return nil, ErrTokenNotValidYet
		} else {
			//默认无法处理
			return nil, ErrTokenInvalid
		}
	}

	// 判断声明类型符合自定义类型
	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
		// 返回该声明
		return claims, nil
	} else {
		return nil, errors.New("token has invalid claims")
	}

}
