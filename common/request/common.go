package request

import (
	"back-end/config"
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

// 基本声明信息
type BaseClaims struct {
	Id          uint
	UUId        string
	Username    string
	NickName    string
	AuthorityId uint
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Authority uint   `json:"authority"` //角色
}
type SysMenusResponse struct {
	Authority int             `json:"authority"`
	Menu      []config.Common `json:"menu"`
}
