package config

import "github.com/golang-jwt/jwt/v5"

// 鉴权配置
type JWT struct {
	SigningKey string ` json:"signing-key" `
	jwt.Claims
}
