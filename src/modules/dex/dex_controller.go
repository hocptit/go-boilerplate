package dex

import (
	"github.com/gin-gonic/gin"
)

func Controller(v1route *gin.RouterGroup) {
	dexRoute := v1route.Group("/dex")
	dexRoute.GET("/", func(context *gin.Context) {
		GetDexes(context)
	})
}
