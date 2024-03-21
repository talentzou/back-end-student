package core

import (
	"back-end/middleware"
	"back-end/router"
	"github.com/gin-gonic/gin"
)

func RunWindowServer() *gin.Engine {
	Server := gin.New()
	Server.Use(gin.Recovery())
	Server.Use(middleware.Cors())
	AppRouter := router.AppRouter
	AppRouter.InitializeRouter(Server)
	return Server
}
