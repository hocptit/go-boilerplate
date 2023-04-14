package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	db "server-go/src/configs"
	"server-go/src/models/entities"
)

type IUserRepository interface {
	CreateUser(chain *entities.Book) error
}

type UserRepository struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserRepository() IUserRepository {
	return &UserRepository{
		userCollection: db.UserCollection,
		ctx:            db.Ctx,
	}
}

func (c *UserRepository) CreateUser(user *entities.Book) error {
	_, err := c.userCollection.InsertOne(c.ctx, user)
	if err != nil {
		return err
	}
	return nil
}
