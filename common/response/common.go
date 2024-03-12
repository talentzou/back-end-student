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

// 响应数据
func ResponseHTTP(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// // 操作状态成功
// func Ok(msg string,data interface{}, c *gin.Context) {
// 	ResponseHTTP(http.StatusNoContent, data, msg, c)
// }

// // 操作状态失败
// func Fail(msg string, c *gin.Context) {
// 	ResponseHTTP(http.StatusNotFound, map[string]interface{}{}, msg, c)
// }

// 系统响应失败返回信息
func FailWithMessage(message string, c *gin.Context) {
	ResponseHTTP(http.StatusNotFound, map[string]interface{}{}, message, c)
}

// 系统响应成功返回信息
func OkWithMessage(msg string, c *gin.Context) {
	ResponseHTTP(http.StatusOK, map[string]any{}, msg, c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	ResponseHTTP(200, data, message, c)
}
