package initialize

import (
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

	// 注册路由
	router.InitRoutes(r)
	global.Log.Info("路由初始化完成")

	return r
}
