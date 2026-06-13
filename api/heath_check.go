package api

import (
	"github.com/gin-gonic/gin"
)

// HeathCheck 健康检查
func HeathCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
