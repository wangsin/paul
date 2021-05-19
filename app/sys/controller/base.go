package sysController

import (
	"github.com/gin-gonic/gin"
	sysLogic "github.com/wangsin/paul/app/sys/logic"
	"net/http"
)

type SysBaseController struct {
}

func (c *SysBaseController) ServerTime(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, sysLogic.GetTime())
}

func (c *SysBaseController) Panicked(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, sysLogic.GetTime())
	panic("guess")
}
