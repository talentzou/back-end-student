package initialize

import (
	"back-end/global"
	"back-end/model/test/dorm"
	"back-end/model/test/expense"
	"back-end/model/test/notice"
	"back-end/model/test/repair"
	"back-end/model/test/student"
	"back-end/model/system"
	"fmt"
	"os"
)

func RegisterTableTest() {
	db := global.Global_Db
	tables := []interface{}{
		&dorm.Bed{},
		&expense.Expense{},
		&dorm.Rate{},
		&student.StudentViolate{},
		&dorm.Stay{},
		&dorm.StudInfo{},
		&repair.Repair{},
		&dorm.Dorm{},
		&dorm.Floor{},
		&notice.SysNotice{},
		&system.SysUser{},
		&system.SysAuthorityBtn{},
	}
	err := db.AutoMigrate(&system.SysUser{},)
			if err != nil {
				fmt.Println("迁移失败")
				os.Exit(0)
			}
	for _, v := range tables {
		
		result := db.Migrator().HasTable(v)
		if !result {
			fmt.Println("执行多少次")
			err := db.AutoMigrate(v)
			if err != nil {
				fmt.Println("迁移失败")
				os.Exit(0)
			}
		}
	}
}
