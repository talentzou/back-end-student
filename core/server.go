package core

import (
	"github.com/gin-gonic/gin"
	"back-end/router"
)

func RunWindowServer() *gin.Engine {
	Server := gin.New()
	Server.Use(gin.Recovery())
    AppRouter:=router.AppRouter
	AppRouter.InitializeRouter(Server)
	return Server
}
