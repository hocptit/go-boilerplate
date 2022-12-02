package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	bookRouters "go-boilerplate/src/modules/books"
	"go-boilerplate/src/shared/exception"
	"io"
	"log"
	"os"
	"time"
)

func Router(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(exception.RecoveryError())
	// Force log's color
	gin.ForceConsoleColor()
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		log.Println("OKOK")
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
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
