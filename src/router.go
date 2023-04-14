package src

import (
	"go-server/src/configs"
	bookRouters "go-server/src/modules/books"
	dexRouters "go-server/src/modules/dex"
	"go-server/src/share/constant"
	"go-server/src/share/exception"
	getLogger "go-server/src/share/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Router(router *gin.Engine, config configs.Config) {
	router.Use(func(context *gin.Context) {
		context.Set(constant.TraceID, uuid.New().String())
	})
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logger := getLogger.GetLogger().Logging
		logger.Infof("[%s] %s %s %s %s %d %s %s %s",
			param.Keys[constant.TraceID],
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage)
		return ""
	}))
	router.Use(exception.RecoveryError(config.AppIsReturnDetailErrors))
	v1route := router.Group("/api/v1")
	/* A */

	/* End A */

	/* B */
	bookRouters.Controller(v1route)
	/* End B */

	/* C */

	/* End C */

	/* D */
	dexRouters.Controller(v1route)
	/* End D */

	/* E */

	/* End E */

	/* F */

	/* End F */

	/* G */

	/* End G */
}
