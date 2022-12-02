package books

import (
	"github.com/gin-gonic/gin"
)

func BookController(v1route *gin.RouterGroup) {
	bookRoute := v1route.Group("/books")
	{
		bookRoute.GET("/", func(context *gin.Context) {
			GetBooks(context)
		})
		bookRoute.POST("/", func(context *gin.Context) {
			CreateBooks(context)
		})
		bookRoute.GET("/:bookId", func(context *gin.Context) {
			GetBookId(context)
		})
	}
}
