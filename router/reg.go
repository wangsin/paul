package router

import (
	"github.com/gin-gonic/gin"
	appSys "github.com/wangsin/paul/app/sys"
	"github.com/wangsin/paul/middleware"
)

func InitRouter() *gin.Engine {
	// 这里需要注册所有系统需要的路由
	baseEngine := gin.New()
	baseEngine.Use(gin.CustomRecovery(middleware.GetGlobalRecoverMiddleware()))

	appSys.InitSysRouter(baseEngine)

	return baseEngine
}
