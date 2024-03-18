package request

import (
	"github.com/golang-jwt/jwt/v5"
)

type PageInfo struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page" form:"page"`         // 页码
	PageSize int         `json:"pageSize" form:"pageSize"` // 每页大小
}

// 声明信息
type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims //预定义的声明
}
type BaseClaims struct {	
	ID          uint
	Username    string
}
