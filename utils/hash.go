package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// 对密码加密
func BcryptHash(password string) []byte {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return bytes
}

// 比较密码哈希值确认
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err==nil
}
