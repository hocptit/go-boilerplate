package basemodel

import "time"

// Model a basic GoLang type which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
// It may be embedded into your model, or you may build your own model without it
//
//	type User type {
//	  basemodel.Model
//	}
type Model struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
