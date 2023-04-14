package books

import (
	"log"
	"net/http"
	"server-go/src/models/entities"
	"server-go/src/models/repositories"
	"server-go/src/modules/books/dto"
	basevalidator "server-go/src/share/base/base_validator"
	getLogger "server-go/src/share/logger"
	response "server-go/src/share/response"
	"server-go/src/share/utils"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context, b repositories.IBookRepository) {
	logger := getLogger.GetLogger().Logging
	logger.Info(utils.TID(c), "My LOGGER")
	TestPanic()
	// Validator
	var listBookDto dto.ListBookDto
	basevalidator.ValidatorQuery(c, &listBookDto)
	// Get data
	books, err := b.FindAll()
	if err != nil {
		panic(err)
	}
	response.ReturnData(c, http.StatusOK, books)
}

func GetBookID(c *gin.Context, b repositories.IBookRepository) {
	logger := getLogger.GetLogger().Logging
	type IDDto struct {
		BookID string `uri:"bookId" binding:"required"`
	}
	var id IDDto

	basevalidator.ValidatorParams(c, &id)
	logger.Info(utils.TID(c), id.BookID)

	book, err := b.GetBookByID(id.BookID)
	if err != nil {
		panic(err)
	}
	response.ReturnData(c, http.StatusOK, book)
}

func CreateBooks(c *gin.Context, b repositories.IBookRepository) {
	var createBookDto dto.CreateBookDto
	basevalidator.ValidatorsBody(c, &createBookDto)
	log.Println(createBookDto)
	err := b.CreateBook(&entities.Book{
		Title:       createBookDto.Title,
		Author:      createBookDto.Author,
		Description: createBookDto.Description,
	})
	if err != nil {
		panic(err)
	}
	response.ReturnData(c, http.StatusCreated, true)
}
