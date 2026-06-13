package router

import (
	"gin-mini-agent/api"

	"github.com/gin-gonic/gin"
)

// InitRoutes 初始化所有路由
func InitRoutes(r *gin.Engine) {
	// 健康检查
	r.GET("/health", api.Health)
}
