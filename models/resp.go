package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Resp http请求响应体
type Resp[T any] struct {
	RequestId string `json:"request_id"`
	Success   bool   `json:"success"`
	Data      T      `json:"data"`
	Msg       string `json:"msg"`
	Total     int64  `json:"total"`
}

// Ok 成功响应
func Ok[T any](c *gin.Context, data T) {
	requestId, _ := c.Get("RequestId")
	c.JSON(http.StatusOK, Resp[T]{
		RequestId: cast.ToString(requestId),
		Success:   true,
		Data:      data,
		Msg:       "操作成功",
	})
}

// OkWithMsg 成功响应（带消息）
func OkWithMsg[T any](c *gin.Context, data T, msg string) {
	requestId, _ := c.Get("RequestId")
	c.JSON(http.StatusOK, Resp[T]{
		RequestId: cast.ToString(requestId),
		Success:   true,
		Data:      data,
		Msg:       msg,
	})
}

// OkWithList 成功响应（带列表和总数）
func OkWithList[T any](c *gin.Context, data T, total int64) {
	requestId, _ := c.Get("RequestId")
	c.JSON(http.StatusOK, Resp[T]{
		RequestId: cast.ToString(requestId),
		Success:   true,
		Data:      data,
		Msg:       "操作成功",
		Total:     total,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, msg string) {
	requestId, _ := c.Get("RequestId")
	c.JSON(http.StatusOK, Resp[any]{
		RequestId: cast.ToString(requestId),
		Success:   false,
		Msg:       msg,
	})
}

// FailWithData 失败响应（带数据）
func FailWithData[T any](c *gin.Context, data T, msg string) {
	requestId, _ := c.Get("RequestId")
	c.JSON(http.StatusOK, Resp[T]{
		RequestId: cast.ToString(requestId),
		Success:   false,
		Data:      data,
		Msg:       msg,
	})
}
