package entities

import "go-boilerplate/src/shared/base/base_model"

type BookEntity struct {
	base_model.Model
	Title       string     `json:"title"`
	AuthorId    int        `json:"author"`
	Description string     `json:"description"`
	Author      UserEntity `gorm:"references:ID"`
}
