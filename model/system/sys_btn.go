package system

type SysBtn struct {
	Id       uint       `json:"-" gorm:"primarykey"`
	BtnKey   string     `json:"btn_key"  gorm:"size:32;comment:按钮key"`
	Title    string     `json:"title" gorm:"size:255;comment:按钮名称"`
	Children []MenuTree `json:"children" gorm:"-"`
	Roles    []Role     `json:"roles" gorm:"many2many:role_btns"`
	MenuId   uint
}
type RoleBtns struct {
	SysBtnId int `gorm:"primarykey"` //按钮id
	RoleId   int `gorm:"primarykey"` //角色id
}

// // 角色按钮信息表;joinForeignKey:BtnId;joinReferences:RoleId"
// type RoleBtns struct {
// 	Id     uint `gorm:"primarykey"` //表主键
// 	RoleId uint `json:"role_id"`    //角色id
// 	BtnId  uint `json:"menu_id"`    //菜单id
// }

// type SysAuthorityBtn struct {
// 	Id        uint   `json:"-" gorm:"primarykey"`
// 	MenuID    uint   `gorm:"comment:菜单ID"`
// 	Authority uint   `json:"authorityId" gorm:"comment:用户角色ID"`
// 	BtnKey    string `json:"btn_key"  gorm:"size:32;comment:按钮key"`
// 	Title     string `json:"title" gorm:"size:255;comment:按钮名称"`

// }
