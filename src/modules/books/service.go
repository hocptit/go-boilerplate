package books

import (
	"go-server/src/configs"
	errorCode "go-server/src/constants/error_code"
	"go-server/src/models/repositories"
	"go-server/src/modules/books/dto"
	basevalidator "go-server/src/share/base/base_validator"
	"go-server/src/share/exception"
	getLogger "go-server/src/share/logger"
	response "go-server/src/share/response"
	"go-server/src/share/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
	logger := getLogger.GetLogger().Logging
	logger.Info(utils.TID(c), "My LOGGER")
	TestPanic()
	// Validator
	var listBookDto dto.ListBookDto
	basevalidator.ValidatorQuery(c, &listBookDto)
	var bookRepository repositories.IBookRepository = repositories.NewBookRepository(configs.GetDB())

	books, err := bookRepository.ListBooks()
	if err != nil {
		panic(err)
	}
	response.ReturnData(c, http.StatusOK, books)
}

func GetBookID(c *gin.Context) {
	type IDDto struct {
		BookID int `uri:"bookId" binding:"required"`
	}
	var id IDDto

	basevalidator.ValidatorParams(c, &id)
	var bookRepository repositories.IBookRepository = repositories.NewBookRepository(configs.GetDB())
	book, err := bookRepository.GetBookByID(id.BookID)
	if err != nil {
		panic(exception.BadRequestError(errorCode.NotFoundBook, err.Error()))
	}
	response.ReturnData(c, http.StatusOK, book)
}

func CreateBooks(c *gin.Context) {
	var userRepository repositories.IUserRepository = repositories.NewUserRepository(configs.GetDB())
	var bookRepository repositories.IBookRepository = repositories.NewBookRepository(configs.GetDB())

	var createBookDto dto.CreateBookDto
	basevalidator.ValidatorsBody(c, &createBookDto)
	log.Println(createBookDto)
	// check author
	_, errGetAuthor := userRepository.GetUserByID(createBookDto.Author)

	if errors.Is(errGetAuthor, gorm.ErrRecordNotFound) {
		panic(exception.BadRequestError(errorCode.NotFoundAuthor, errGetAuthor.Error()))
		return
	}
	book := bookRepository.CreateBooks(createBookDto)
	response.ReturnData(c, http.StatusCreated, book)
}
