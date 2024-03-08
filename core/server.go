package core

import (
	"back-end/router"
	"github.com/gin-gonic/gin"
)

func RunWindowServer() *gin.Engine {
	Server := gin.New()
	Server.Use(gin.Recovery())
	Server.Group("/")
	return Server
}
