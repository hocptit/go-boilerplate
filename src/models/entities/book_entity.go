package entities

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Book struct {
	ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Author      string    `json:"author" bson:"author"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (c *Book) MarshalBSON() ([]byte, error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	c.UpdatedAt = time.Now()

	type my Book
	return bson.Marshal((*my)(c))
}
