package request

type ReqRelateRoleAndBtn struct {
	RoleId        int   `json:"roleId"`
	RoleBtnIdList []int `json:"roleBtnIdList"`
}
type ReqRelateRoleAndMenu struct {
	RoleId         int   `json:"roleId"`
	RoleMenuIdList []int `json:"roleMenuIdList"`
}
