// package main

// import (
// 	"back-end/model/system"
// )

// func InitData() []system.Role {

// 	menu := []system.MenuTree{
// 		{Id: 1, Name: "sys", ParentId: 0, Title: "系统管理"},
// 		{Id: 2, Name: "dormitory-management", ParentId: 0, Title: "宿舍管理"},
// 		{Id: 3, Name: "Maintenance", ParentId: 0, Title: "维修管理"},
// 		{Id: 4, Name: "studentInfo", ParentId: 0, Title: "学生管理"},
// 		{Id: 5, Name: "notice", ParentId: 0, Title: "通告管理"},
// 		{Id: 6, Name: "user", ParentId: 1, Title: "用户管理"},
// 		{Id: 7, Name: "menu", ParentId: 1, Title: "菜单管理"},
// 		{Id: 8, Name: "role", ParentId: 1, Title: "角色管理"},
// 		{Id: 10, Name: "floors", ParentId: 2, Title: "宿舍楼信息"},
// 		{Id: 11, Name: "dorm", ParentId: 2, Title: "宿舍信息"},
// 		{Id: 12, Name: "rate", ParentId: 2, Title: "宿舍评分"},
// 		{Id: 13, Name: "stay", ParentId: 2, Title: "留宿申请"},
// 		{Id: 14, Name: "bed", ParentId: 2, Title: "床位信息"},
// 		{Id: 15, Name: "expense", ParentId: 2, Title: "水电费信息"},
// 		{Id: 16, Name: "equipment", ParentId: 3, Title: "维修列表"},
// 		{Id: 17, Name: "student", ParentId: 4, Title: "学生信息"},
// 		{Id: 18, Name: "violate", ParentId: 4, Title: "学生违纪"},
// 		{Id: 19, Name: "message", ParentId: 5, Title: "宿舍通告"},
// 		{Id: 20, Name: "person", ParentId: 0, Title: "个人信息管理"},
// 		{Id: 21, Name: "userSelfInfo", ParentId: 20, Title: "个人信息"},
// 	}
// 	studentMenu := []system.MenuTree{
// 		{Id: 2, Name: "dormitory-management", ParentId: 0, Title: "宿舍管理"},
// 		{Id: 3, Name: "Maintenance", ParentId: 0, Title: "维修管理"},
// 		{Id: 4, Name: "studentInfo", ParentId: 0, Title: "学生管理"},
// 		{Id: 12, Name: "rate", ParentId: 2, Title: "宿舍评分"},
// 		{Id: 13, Name: "stay", ParentId: 2, Title: "留宿申请"},
// 		{Id: 15, Name: "expense", ParentId: 2, Title: "水电费信息"},
// 		{Id: 16, Name: "equipment", ParentId: 3, Title: "维修列表"},
// 		{Id: 18, Name: "violate", ParentId: 4, Title: "学生违纪"},
// 		{Id: 20, Name: "person", ParentId: 0, Title: "个人信息管理"},
// 		{Id: 21, Name: "userSelfInfo", ParentId: 20, Title: "个人信息"},
// 	}
// 	dorm := []system.MenuTree{
// 		{Id: 2, Name: "dormitory-management", ParentId: 0, Title: "宿舍管理"},
// 		{Id: 3, Name: "Maintenance", ParentId: 0, Title: "维修管理"},
// 		{Id: 4, Name: "studentInfo", ParentId: 0, Title: "学生管理"},
// 		{Id: 10, Name: "floors", ParentId: 2, Title: "宿舍楼信息"},
// 		{Id: 11, Name: "dorm", ParentId: 2, Title: "宿舍信息"},
// 		{Id: 12, Name: "rate", ParentId: 2, Title: "宿舍评分"},
// 		{Id: 13, Name: "stay", ParentId: 2, Title: "留宿申请"},
// 		{Id: 14, Name: "bed", ParentId: 2, Title: "床位信息"},
// 		{Id: 15, Name: "expense", ParentId: 2, Title: "水电费信息"},
// 		{Id: 16, Name: "equipment", ParentId: 3, Title: "维修列表"},
// 		{Id: 17, Name: "student", ParentId: 4, Title: "学生信息"},
// 		{Id: 18, Name: "violate", ParentId: 4, Title: "学生违纪"},
// 		{Id: 20, Name: "person", ParentId: 0, Title: "个人信息管理"},
// 		{Id: 21, Name: "userSelfInfo", ParentId: 20, Title: "个人信息"},
// 	}
// 	role := []system.Role{
// 		{
// 			Id:        1,
// 			SysUserId: 1001,
// 			RoleName:  "admin",
// 			MenuTrees: menu,
// 		},
// 		{
// 			Id:        2,
// 			SysUserId: 2001,
// 			RoleName:  "宿管",
// 			MenuTrees: dorm,
// 		},
// 		{
// 			Id:        3,
// 			SysUserId: 30001,
// 			RoleName:  "学生",
// 			MenuTrees: studentMenu,
// 		}}
// 	return role
// }
