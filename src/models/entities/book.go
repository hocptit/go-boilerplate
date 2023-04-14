package entities

import basemodel "go-server/src/share/base/base_model"

type BookEntity struct {
	basemodel.Model
	Title       string     `json:"title"`
	AuthorID    int        `json:"author"`
	Description string     `json:"description"`
	Author      UserEntity `gorm:"references:ID"`
}
