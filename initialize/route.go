package initialize

import (
	"back-end/middleware"
	"back-end/router"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(s *gin.Engine) {
	root := s.Group("")
	s.Static("/public", "./public") //静态资源
	systemRouter := router.AppRoute.System
	expandRouter := router.AppRoute.Expand

	// 系统路由
	{
		sysBase := root.Group("base")
		systemRouter.InitializeSys(sysBase)
	}
	// jwt路由
	{
		expandJwt := root.Group("/jwt")
		expandJwt.Use(middleware.JwtAuth())
		expandRouter.InitializeExpandRouter(expandJwt)
	}

}
