package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Ctx            = context.TODO()
	BookCollection *mongo.Collection
	UserCollection *mongo.Collection
)

func Setup(uri, database string) {
	option := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(Ctx, option)
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(Ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	db := client.Database(database)
	BookCollection = db.Collection("books")
}
