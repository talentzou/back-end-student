package core

import (
	"back-end/initialize"
	"back-end/middleware"
	"github.com/gin-gonic/gin"
)

func RunWindowServer() *gin.Engine {
	Server := gin.New()                 //创建工程实例
	Server.Use(gin.Recovery())          //系统自动修复
	Server.Use(middleware.Cors())       //允许跨域
	initialize.InitializeRouter(Server) //初始路由
	return Server
}
