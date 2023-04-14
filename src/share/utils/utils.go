package utils

import (
	"server-go/src/share/constant"

	"github.com/gin-gonic/gin"
)

func TID(c *gin.Context) string {
	return "[" + c.Keys[constant.TraceID].(string) + "] "
}

func CoalesceInt(args ...int) int {
	for _, arg := range args {
		if arg != 0 {
			return arg
		}
	}
	return 1
}
