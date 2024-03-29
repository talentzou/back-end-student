package system

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UUID      string `json:"UUID" gorm:"size:256;comment:用户uuid"`
	UserName  string `json:"userName" gorm:"size:256;comment:用户登录名"`
	Password  string `json:"-"  gorm:"comment:用户登录密码"`
	Sex       string `json:"sex"  gorm:"size:256;comment:性别"`
	Nickname  string `json:"nickname" gorm:"size:256;comment:用户昵称"`
	Telephone string `json:"telephone" gorm:"size:256;default:18100000000;comment:手机号码"`
	Email     string `json:"email" gorm:"size:256;default:123456@qq.com;comment:邮箱"`
	Avatar    string `json:"avatar" gorm:"size:256;default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:头像"`
	Authority uint   `json:"authorityId" gorm:"default:2;comment:用户角色ID"`
	Remark    string `json:"remark" gorm:"size:256;comment:备注"`
	Dorm      string `json:"dorm" gorm:"size:256;comment:宿舍"`
}

type LoginResponse struct {
	User      SysUser `json:"user"`
	Token     string  `json:"token"`
	ExpiresAt int64   `json:"expiresAt"`
}

type ChangeUserInfo struct {
	Sex       string `json:"sex"  gorm:"size:256;comment:性别"`
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Telephone string `json:"telephone"  gorm:"comment:用户手机号"`           // 用户手机号
	Email     string `json:"email"  gorm:"comment:用户邮箱"`                // 用户邮箱
	Avatar    string `json:"avatar" gorm:"size:256;default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:头像"`
	Dorm      string `json:"dorm" gorm:"size:256;comment:宿舍"`
}
