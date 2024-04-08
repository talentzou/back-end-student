package system

// 菜单树表
type MenuTree struct {
	Id        uint       `json:"id" gorm:"primarykey"`
	ParentId  uint       `json:"parent_id"`
	Name      string     `json:"name" gorm:"size:256"`
	Title     string     `json:"title" gorm:"size:256"`
	Children  []MenuTree `json:"children" gorm:"-"`
	RoleMenus []RoleMenu `json:"-" gorm:"many2many:role_authority_menus;foreignKey:Id;associationForeignKey:MenuId"`
}

//foreignKey:MenuId;references:Id
//foreignKey:Id;associationForeignKey:MenuId

// 角色菜单树信息表
type RoleMenu struct {
	Id          uint       `gorm:"primarykey"` //表主键
	AuthorityId uint       `json:"authorityId" gorm:"comment:用户角色Id"`
	MenuId      uint       `json:"menu_id"`
	RoleName    string     `json:"roleName" gorm:"comment:角色名称"`
	MenuTrees   []MenuTree `json:"Menu" gorm:"many2many:role_authority_menus;foreignKey:Id;associationForeignKey:Id"`
}

//foreignKey:MenuId;associationForeignKey:Id
// 菜单树表ForeignKey:Id;references:MenuId
