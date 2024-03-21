package system

import (
	"back-end/common/response"
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
)

// 上传图片
func (b *BaseApi) UploadHandle(c *gin.Context) {
	fmt.Println("照片上传")
	uploadFile, err := c.FormFile("image")
	if err != nil {
		fmt.Println("错误1",err.Error())
		response.FailWithMessage("上传文件失败", c)
		return
	}
	filepath := "./public/static/" + uploadFile.Filename
	err2 := c.SaveUploadedFile(uploadFile, filepath)
	if err2 != nil {
		fmt.Println("错误2",err2.Error())
		response.Fail("文件服务器保存失败", c)
		return
	}
	imageUrl :=path.Join(c.Request.Host,filepath)
	fmt.Println("照片地址",imageUrl)
	response.Ok("上传成功", gin.H{"url": "http://"+imageUrl}, c)
}
