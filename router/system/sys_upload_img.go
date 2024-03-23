package system

import (
	"back-end/controller/system"
	"github.com/gin-gonic/gin"
)

type SysUpload struct{}

func (S *SysUpload) SystemUploadImg(router *gin.RouterGroup) {
	uploadRouter := router.Group("upload")
	{
		uploadRouter.POST("imageUpload", system.SystemApi.UploadHandle)
		uploadRouter.GET("imageDownload", func(c *gin.Context) {

		})
	}

}
