package main

import (
	"back-end/core"
	"back-end/global"
	"back-end/initialize"
)

func main() {
	// core.ViperRouter()//初始化读取路由配置文件
	core.ViperServer()//读取服务器配置文件
	global.Global_Db = initialize.GormMysql() //初始化数据库
	if global.Global_Db != nil {
		// initialize.RegisterTable() // 初始化表
		initialize.RegisterTableTest()
		db, _ := global.Global_Db.DB()
		defer db.Close()
	}
	r := core.RunWindowServer()//创建应用实例
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
