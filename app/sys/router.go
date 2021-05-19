package appSys

import (
	"github.com/gin-gonic/gin"
	sysController "github.com/wangsin/paul/app/sys/controller"
)

func InitSysRouter(e *gin.Engine) {
	sysGroup := e.Group("/paul/system")
	{
		sysGroup.GET("/time", new(sysController.SysBaseController).ServerTime)
		sysGroup.GET("/panicked", new(sysController.SysBaseController).Panicked)
	}
}
