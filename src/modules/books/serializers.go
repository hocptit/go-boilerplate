package books

import (
	"github.com/gin-gonic/gin"
)

type BookSerializer struct {
	C *gin.Context
}

type BookResponse struct {
	Title       string `json:"title"`
	Author      int    `json:"author"`
	Description string `json:"description"`
}

func (bookSerializer *BookSerializer) Response() BookResponse {
	res := BookResponse{}
	return res
}

type Serializer struct {
	C *gin.Context
}

func (bookSerializer *Serializer) Response() []BookResponse {
	var response []BookResponse
	// nolint
	return response
}
