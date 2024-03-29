package system

type SysAuthorityBtn struct {
	Id        uint   `json:"-" gorm:"primarykey"`
	Authority uint   `json:"authorityId" gorm:"comment:用户角色ID"`
	BtnName   string `json:"btn_name"  gorm:"size:32;comment:拥有按钮"`
}
