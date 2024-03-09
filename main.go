package main

import (
	"back-end/core"
	"back-end/global"
	"back-end/initialize"
)

func main() {
	global.Global_Db = initialize.GormMysql()
	if global.Global_Db != nil {
		initialize.RegisterTable() // 初始化表
	}
	r := core.RunWindowServer()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
