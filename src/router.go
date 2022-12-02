package src

import (
	"github.com/gin-gonic/gin"
	bookRouters "go-boilerplate/src/modules/books"
	"go-boilerplate/src/shared/exception"
	getLogger "go-boilerplate/src/shared/logger"
)

func Router(router *gin.Engine) {
	router.Use(func(context *gin.Context) {
		context.Set("A", "BC")
	})
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		//fmt.Println(param.Keys["A"])
		logger := getLogger.GetLogger().Logging
		logger.Infof("%s %s %s %s %d %s %s %s",
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
	router.Use(exception.RecoveryError())
	v1route := router.Group("/api/v1")

	{
		/* A */

		/* A */

		/* B */
		bookRouters.BookController(v1route)
		/* B */

		/* C */

		/* C */

		/* D */

		/* D */

		/* E */

		/* E */

		/* F */

		/* F */

	}

}
