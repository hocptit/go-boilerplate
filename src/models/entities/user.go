package entities

import "go-boilerplate/src/shared/base/base_model"

type Address struct {
	District string `json:"district"`
	Street   string `json:"street"`
}

type UserEntity struct {
	base_model.Model
	Username string  `json:"username"`
	Password string  `json:"password"`
	Address  Address `json:"address" gorm:"embedded"`
}
