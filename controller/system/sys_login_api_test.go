package system

import (
	"back-end/model/system"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBaseApi_Login(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		b    *BaseApi
		args args
	}{

		// TODO: Add test cases.
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Login(tt.args.c)
		})
	}
}

func TestBaseApi_Logout(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		j    *BaseApi
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.j.Logout(tt.args.c)
		})
	}
}

func TestBaseApi_TokenNext(t *testing.T) {
	type args struct {
		c    *gin.Context
		user system.SysUser
	}
	tests := []struct {
		name string
		b    *BaseApi
		args args
	}{
	  {},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.TokenNext(tt.args.c, tt.args.user)
		})
	}
}
