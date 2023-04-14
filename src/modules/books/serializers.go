package books

import (
	"go-server/src/models/entities"

	"github.com/gin-gonic/gin"
)

type BookSerializer struct {
	C *gin.Context
	entities.BookEntity
}

type BookResponse struct {
	Title       string `json:"title"`
	Author      int    `json:"author"`
	Description string `json:"description"`
}

func (bookSerializer *BookSerializer) Response() BookResponse {
	res := BookResponse{
		Title:       bookSerializer.Title,
		Description: bookSerializer.Description,
	}
	return res
}

type Serializer struct {
	C     *gin.Context
	Books []entities.BookEntity
}

func (bookSerializer *Serializer) Response() []BookResponse {
	var response []BookResponse
	// nolint
	for _, book := range bookSerializer.Books {
		serializer := BookSerializer{bookSerializer.C, book}
		response = append(response, serializer.Response())
	}
	return response
}
