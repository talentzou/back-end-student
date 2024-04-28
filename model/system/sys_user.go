package system

import (
	"back-end/model/test/dorm"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UserName  string `json:"userName" gorm:"size:256;comment:用户登录名"`
	Password  string `json:"-"  gorm:"size:256;default:123456;comment:用户登录密码"`
	Sex       string `json:"sex"  gorm:"size:256;comment:性别"`
	Nickname  string `json:"nickname" gorm:"default:无;size:256;comment:用户昵称"`
	Telephone string `json:"telephone" gorm:"size:256;default:18100000000;comment:手机号码"`
	Avatar    string `json:"avatar" gorm:"size:256;default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:头像"`

	Remark string    `json:"remark" gorm:"size:256;comment:备注"`
	RoleId uint      `json:"roleId" gorm:"default:2;comment:用户角色ID"`
	DormId uint      `json:"dormId" gorm:"size:256;comment:宿舍id"`
	Dorm   dorm.Dorm `json:"dorm"`

	// Dorm   string `json:"dorm" gorm:"size:256;comment:宿舍"`
	// Authority uint   `json:"authorityId" gorm:"default:2;comment:用户角色ID"`
	// Role      Role   `json:"role"`

}

// SysAuthorityBtns []SysAuthorityBtn `json:"sysAuthorityBtns" gorm:"foreignKey:Authority;references:Authority"` //用户菜单按钮表
type LoginResponse struct {
	User      SysUser `json:"user"`
	Token     string  `json:"token"`
	ExpiresAt int64   `json:"expiresAt"`
}

type ChangeUserInfo struct {
	Sex       string `json:"sex"  gorm:"size:256;comment:性别"`
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Telephone string `json:"telephone"  gorm:"comment:用户手机号"`           // 用户手机号
	Avatar    string `json:"avatar" gorm:"size:256;default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:头像"`
	// Dorm      string `json:"dorm" gorm:"size:256;comment:宿舍"`
}
