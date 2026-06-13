package initialize

import (
	"gin-mini-agent/api"
	"gin-mini-agent/pkg/global"
	"gin-mini-agent/router"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	if global.Conf.System.RunMode == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()

	r := gin.New()
	// 内置中间件
	r.Use(requestid.New())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/heath_check", api.HeathCheck)
	global.Log.Info("初始化健康检查接口完成")

	// API 路由
	apiGroup := r.Group(global.Conf.System.UrlPathPrefix)
	v1Group := apiGroup.Group("v1")
	router.InitFileRouter(v1Group)
	global.Log.Info("初始化基础路由完成")

	return r
}
