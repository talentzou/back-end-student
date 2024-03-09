package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 200
	ERROR   = 404
)

// 响应数据
func ResponseHTTP(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

//操作状态成功
func Ok(msg string,c *gin.Context) {
	ResponseHTTP(SUCCESS, map[string]interface{}{},msg, c)
} 

// 操作状态失败
func Fail(msg string,c *gin.Context) {
	ResponseHTTP(ERROR, map[string]interface{}{}, msg, c)
}
