package books

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go-boilerplate/src/configs"
	"go-boilerplate/src/constants/error_code"
	"go-boilerplate/src/models/repositories"
	"go-boilerplate/src/modules/books/dto"
	"go-boilerplate/src/shared/base/base_validator"
	"go-boilerplate/src/shared/exception"
	getLogger "go-boilerplate/src/shared/logger"
	response "go-boilerplate/src/shared/response"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetBooks(c *gin.Context) {
	logger := getLogger.GetLogger().Logging
	logger.Info("Xin chao")
	TestPanic()
	// Validator
	var listBookDto dto.ListBookDto
	base_validator.ValidatorQuery(c, &listBookDto)
	//fmt.Println(listBookDto)
	var bookRepository repositories.IBookRepository = repositories.NewBookRepository(configs.GetDB())

	books, err := bookRepository.ListBooks()
	if err != nil {
		panic(err)
	}
	response.ReturnData(c, http.StatusOK, books)

	//serializer := &BooksSerializer{c, books}
	//response.ReturnData(c, http.StatusOK, serializer.Response())

}

func GetBookId(c *gin.Context) {
	type IDDto struct {
		BookId int `uri:"bookId" binding:"required"`
	}
	var id IDDto

	base_validator.ValidatorParams(c, &id)
	var bookRepository repositories.IBookRepository = repositories.NewBookRepository(configs.GetDB())
	book, err := bookRepository.GetBookById(id.BookId)
	if err != nil {
		panic(exception.BadRequestError(error_code.NOT_FOUND_BOOK, err.Error()))
	}
	response.ReturnData(c, http.StatusOK, book)
}

func CreateBooks(c *gin.Context) {
	var userRepository repositories.IUserRepository = repositories.NewUserRepository(configs.GetDB())
	var bookRepository repositories.IBookRepository = repositories.NewBookRepository(configs.GetDB())

	var createBookDto dto.CreateBookDto
	base_validator.ValidatorsBody(c, &createBookDto)
	log.Println(createBookDto)
	// check author
	_, errGetAuthor := userRepository.GetUserById(createBookDto.Author)

	if errors.Is(errGetAuthor, gorm.ErrRecordNotFound) {
		panic(exception.BadRequestError(error_code.NOT_FOUND_AUTHOR, errGetAuthor.Error()))
		return
	}
	book := bookRepository.CreateBooks(createBookDto)
	response.ReturnData(c, http.StatusCreated, book)
	return
}
