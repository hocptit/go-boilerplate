package repositories

import (
	"go-server/src/models/entities"
	"go-server/src/modules/books/dto"

	"gorm.io/gorm"
)

type IBookRepository interface {
	ListBooks(conds ...interface{}) ([]entities.BookEntity, error)
	CreateBooks(bookData dto.CreateBookDto) entities.BookEntity
	GetBookByID(id int) (entities.BookEntity, error)
}
type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (bookRepository *BookRepository) ListBooks(conds ...interface{}) ([]entities.BookEntity, error) {
	var books []entities.BookEntity
	err := bookRepository.db.Find(&books, conds).Error
	return books, err
}

func (bookRepository *BookRepository) CreateBooks(bookData dto.CreateBookDto) entities.BookEntity {
	book := entities.BookEntity{
		Title:       bookData.Title,
		AuthorID:    bookData.Author,
		Description: bookData.Description,
	}
	bookRepository.db.Create(&book)
	return book
}

func (bookRepository *BookRepository) GetBookByID(id int) (entities.BookEntity, error) {
	var book entities.BookEntity
	err := bookRepository.db.First(&book, id).Error
	return book, err
}
