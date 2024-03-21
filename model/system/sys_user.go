package system

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UUID      string `json:"UUID" gorm:"size:256;comment:用户uuid"`
	UserName  string `json:"userName" gorm:"size:256;comment:用户登录名"`
	Password  string `json:"-"  gorm:"comment:用户登录密码"`
	Nickname  string `json:"nickname" gorm:"size:256;comment:用户昵称"`
	Avatar    string `json:"avatar" gorm:"size:256;default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:头像"`
	Authority uint   `json:"authorityId" gorm:"default:2;comment:用户角色ID"`
}

type LoginResponse struct {
	User      SysUser `json:"user"`
	Token     string  `json:"token"`
	ExpiresAt int64   `json:"expiresAt"`
}
