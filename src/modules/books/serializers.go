package books

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/src/models/entities"
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

func (self *BookSerializer) Response() BookResponse {
	res := BookResponse{
		Title:       self.Title,
		Description: self.Description,
	}
	return res
}

type BooksSerializer struct {
	C     *gin.Context
	Books []entities.BookEntity
}

func (self *BooksSerializer) Response() []BookResponse {
	var response []BookResponse
	for _, book := range self.Books {
		serializer := BookSerializer{self.C, book}
		response = append(response, serializer.Response())
	}
	return response
}
