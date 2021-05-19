package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGlobalRecoverMiddleware() gin.RecoveryFunc {
	return func(context *gin.Context, err interface{}) {
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "panicked",
				"code": http.StatusInternalServerError,
				"err":  err,
			})
		}
	}
}
