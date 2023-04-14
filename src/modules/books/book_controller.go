package books

import (
	"github.com/gin-gonic/gin"
	"server-go/src/models/repositories"
	_type "server-go/src/share/type"
)

func Controller(v1route _type.V1Route, b repositories.IBookRepository) {
	bookRoute := v1route.V1ProtectedRoute.Group("books")
	bookRoute.GET("/", func(context *gin.Context) {
		GetBooks(context, b)
	})
	bookRoute.POST("/", func(context *gin.Context) {
		CreateBooks(context, b)
	})
	bookRoute.GET("/:bookId", func(context *gin.Context) {
		GetBookID(context, b)
	})
}
