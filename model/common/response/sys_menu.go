package response

import "back-end/model/system"

type SysMenusResponse struct {
	Menus []system.MenuTree `json:"menus"`
	
}

// type SysBaseMenusResponse struct {
// 	Menus []system.SysBaseMenu `json:"menus"`
// }

// type SysBaseMenuResponse struct {
// 	Menu system.SysBaseMenu `json:"menu"`
// }
