package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	db "server-go/src/configs"
	"server-go/src/models/entities"
)

type IBookRepository interface {
	CreateBook(chain *entities.Book) error
	FindAll() ([]*entities.Book, error)
	GetBookByID(id string) (*entities.Book, error)
}

type BookRepository struct {
	bookCollection *mongo.Collection
	ctx            context.Context
}

func NewBookRepository() IBookRepository {
	return &BookRepository{
		bookCollection: db.BookCollection,
		ctx:            db.Ctx,
	}
}

func (c *BookRepository) CreateBook(book *entities.Book) error {
	_, err := c.bookCollection.InsertOne(c.ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (c *BookRepository) FindAll() ([]*entities.Book, error) {
	var books []*entities.Book
	cursor, err := c.bookCollection.Find(c.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(c.ctx, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (c *BookRepository) GetBookByID(id string) (*entities.Book, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var book entities.Book
	err = c.bookCollection.FindOne(c.ctx, bson.M{"_id": objID}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
