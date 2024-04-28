package middleware

import (
	"back-end/utils"

	"github.com/gin-gonic/gin"
)

func IsStudent() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		utils.GetUserRoleId(ctx)
	}
}