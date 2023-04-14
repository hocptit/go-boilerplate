package basemodel

import "time"

// Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
// It may be embedded into your model, or you may build your own model without it
//
//	type User struct {
//	  basemodel.Model
//	}
type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
