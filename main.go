package main

import (
	"back-end/core"
	"back-end/global"
	"back-end/initialize"
	"fmt"
)

func main() {
	core.Viper()//读取配置文件
	global.Global_Db = initialize.GormMysql() //初始化数据库
	if global.Global_Db != nil {
		// initialize.RegisterTable() // 初始化表
		initialize.RegisterTableTest()
	}
	r := core.RunWindowServer()//创建应用实例
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("测试成功")
}
