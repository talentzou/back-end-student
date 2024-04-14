package initialize

import (
	"back-end/global"
	"back-end/model/system"
	"back-end/model/test/dorm"
	"back-end/model/test/notice"
	"back-end/model/test/repair"
	"back-end/model/test/student"
	"fmt"
	"os"

	"gorm.io/gorm"
	// "time"
	// "gorm.io/gorm"
)

func RegisterTableTest() {
	db := global.Global_Db
	tables := []interface{}{
		&dorm.Bed{},
		&dorm.Expense{},
		&dorm.Rate{},
		&student.StudentViolate{},
		&dorm.Stay{},
		&dorm.StudInfo{},
		&repair.Repair{},
		&dorm.Dorm{},
		&dorm.Floor{},
		&notice.SysNotice{},
		&system.SysUser{},
		&system.Role{},
		&system.RoleMenus{},
	}
	for _, v := range tables {
		result := db.Migrator().HasTable(v)
		if !result {
			err := db.AutoMigrate(v)
			if err != nil {
				fmt.Println("迁移失败")
				os.Exit(0)
			}
			fmt.Println("执行多少次")

		}
	}

	// role := InitData()
	// err := db.Model(&system.Role{}).Create(&role).Error
	// if err != nil {
	// 	fmt.Println("-----------------初始化菜单失败")
	// }
	// User := mockUser()
	// err1 := db.Model(&system.SysUser{}).Create(&User).Error
	// if err1 != nil {
	// 	fmt.Println("-----------------添加用户失败")
	// }
	// floor := dorm.Floor{
	// 	Id:         1,
	// 	DormAmount: 10,
	// 	FloorsName: "A",
	// 	FloorsType: "男生宿舍",
	// }
	// dorm1 := dorm.Dorm{
	// 	DormNumber: "101",
	// 	Img:        "http://localhost:8080/public/static/微信图片_20220304221251.jpg",
	// 	Capacity:   6,
	// 	Floor:      floor,
	// }
	// err1 := db.Model(&dorm.Dorm{}).Create(&dorm1).Error
	// if err1 != nil {
	// 	fmt.Println("-----------------添加楼宇失败")
	// }
}

func InitData() []system.Role {

	menu := []system.MenuTree{
		{Id: 1, Name: "sys", ParentId: 0, Title: "系统管理"},
		{Id: 2, Name: "dormitory-management", ParentId: 0, Title: "宿舍管理"},
		{Id: 3, Name: "Maintenance", ParentId: 0, Title: "维修管理"},
		{Id: 4, Name: "studentInfo", ParentId: 0, Title: "学生管理"},
		{Id: 5, Name: "notice", ParentId: 0, Title: "通告管理"},
		{Id: 6, Name: "user", ParentId: 1, Title: "用户管理"},
		{Id: 7, Name: "menu", ParentId: 1, Title: "菜单管理"},
		{Id: 8, Name: "role", ParentId: 1, Title: "角色管理"},
		{Id: 10, Name: "floors", ParentId: 2, Title: "宿舍楼信息"},
		{Id: 11, Name: "dorm", ParentId: 2, Title: "宿舍信息"},
		{Id: 12, Name: "rate", ParentId: 2, Title: "宿舍评分"},
		{Id: 13, Name: "stay", ParentId: 2, Title: "留宿申请"},
		{Id: 14, Name: "bed", ParentId: 2, Title: "床位信息"},
		{Id: 15, Name: "expense", ParentId: 2, Title: "水电费信息"},
		{Id: 16, Name: "equipment", ParentId: 3, Title: "维修列表"},
		{Id: 17, Name: "student", ParentId: 4, Title: "学生信息"},
		{Id: 18, Name: "violate", ParentId: 4, Title: "学生违纪"},
		{Id: 19, Name: "message", ParentId: 5, Title: "宿舍通告"},
		{Id: 20, Name: "person", ParentId: 0, Title: "个人信息管理"},
		{Id: 21, Name: "userSelfInfo", ParentId: 20, Title: "个人信息"},
	}
	studentMenu := []system.MenuTree{
		{Id: 2, Name: "dormitory-management", ParentId: 0, Title: "宿舍管理"},
		{Id: 3, Name: "Maintenance", ParentId: 0, Title: "维修管理"},
		{Id: 4, Name: "studentInfo", ParentId: 0, Title: "学生管理"},
		{Id: 12, Name: "rate", ParentId: 2, Title: "宿舍评分"},
		{Id: 13, Name: "stay", ParentId: 2, Title: "留宿申请"},
		{Id: 15, Name: "expense", ParentId: 2, Title: "水电费信息"},
		{Id: 16, Name: "equipment", ParentId: 3, Title: "维修列表"},
		{Id: 18, Name: "violate", ParentId: 4, Title: "学生违纪"},
		{Id: 20, Name: "person", ParentId: 0, Title: "个人信息管理"},
		{Id: 21, Name: "userSelfInfo", ParentId: 20, Title: "个人信息"},
	}
	dorm := []system.MenuTree{
		{Id: 2, Name: "dormitory-management", ParentId: 0, Title: "宿舍管理"},
		{Id: 3, Name: "Maintenance", ParentId: 0, Title: "维修管理"},
		{Id: 4, Name: "studentInfo", ParentId: 0, Title: "学生管理"},
		{Id: 10, Name: "floors", ParentId: 2, Title: "宿舍楼信息"},
		{Id: 11, Name: "dorm", ParentId: 2, Title: "宿舍信息"},
		{Id: 12, Name: "rate", ParentId: 2, Title: "宿舍评分"},
		{Id: 13, Name: "stay", ParentId: 2, Title: "留宿申请"},
		{Id: 14, Name: "bed", ParentId: 2, Title: "床位信息"},
		{Id: 15, Name: "expense", ParentId: 2, Title: "水电费信息"},
		{Id: 16, Name: "equipment", ParentId: 3, Title: "维修列表"},
		{Id: 17, Name: "student", ParentId: 4, Title: "学生信息"},
		{Id: 18, Name: "violate", ParentId: 4, Title: "学生违纪"},
		{Id: 20, Name: "person", ParentId: 0, Title: "个人信息管理"},
		{Id: 21, Name: "userSelfInfo", ParentId: 20, Title: "个人信息"},
	}
	role := []system.Role{
		{
			Model: gorm.Model{
				ID: 1,
			},
			SysUserId: 1001,
			RoleName:  "admin",
			MenuTrees: menu,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			SysUserId: 2001,
			RoleName:  "宿管",
			MenuTrees: dorm,
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			SysUserId: 30001,
			RoleName:  "学生",
			MenuTrees: studentMenu,
		}}
	return role
}

func mockUser() system.SysUser {
	user := system.SysUser{
		Model: gorm.Model{
			ID: 4444,
		},
		UserName:  "test1",
		Password:  "$2a$10$8HFIPEmF/bTcloYrjeiRtuhERB05/WF7TtQOGxveB9c/cEu.VRmie",
		Nickname:  "test1",
		Telephone: "18100000000",
		Avatar:    "https://qmplusimg.henrongyi.top/gva_header.jpg",
		Authority: 1,
		Sex:       "男",
		Role: system.Role{
			RoleName: "测试",
		},
	}
	return user
}
