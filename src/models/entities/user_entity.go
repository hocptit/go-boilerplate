package entities

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type User struct {
	ID            string    `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName     string    `json:"firstName" bson:"firstName"`
	LastName      string    `json:"lastName" bson:"lastName"`
	Email         string    `json:"email" bson:"email"`
	EmailVerified bool      `json:"emailVerified" bson:"emailVerified"`
	UserUID       string    `json:"userUID" bson:"userUID"`
	IsLock        bool      `json:"isLock" bson:"isLock"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
}

func (c *User) MarshalBSON() ([]byte, error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	c.UpdatedAt = time.Now()

	type my User
	return bson.Marshal((*my)(c))
}
