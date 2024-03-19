package system

import (
	"back-end/common/response"
	"github.com/gin-gonic/gin"
)

// 上传图片
func (b *BaseApi) UploadHandle(c *gin.Context) {
	uploadFile, err := c.FormFile("image")
	if err != nil {
		response.FailWithMessage("上传文件失败", c)
		return
	}
	filepath := "/public/images/" + uploadFile.Filename
	err2 := c.SaveUploadedFile(uploadFile, filepath)
	if err2 != nil {
		response.Fail("文件服务器保存失败", c)
		return
	}
	imageUrl := c.Request.Host + filepath
	response.Ok("上传成功", gin.H{"url": imageUrl}, c)
}
