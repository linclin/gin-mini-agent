package router

import (
	"gin-mini-agent/api/v1/file"
	"gin-mini-agent/pkg/global"

	"github.com/gin-gonic/gin"
)

func InitFileRouter(r *gin.RouterGroup) (R gin.IRoutes) {
	router := r.Group("file", gin.BasicAuth(gin.Accounts{
		global.Conf.Auth.User: global.Conf.Auth.Password,
	}))
	{
		router.POST("/cp", file.FileCp)
	}
	return router
}
