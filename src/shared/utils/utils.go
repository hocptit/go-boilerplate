package utils

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/src/constants"
)

func TID(c *gin.Context) any {
	return "[" + c.Keys[constants.TRACE_ID].(string) + "] "
}
