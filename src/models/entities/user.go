package entities

import basemodel "go-server/src/share/base/base_model"

type Address struct {
	District string `json:"district"`
	Street   string `json:"street"`
}

type UserEntity struct {
	basemodel.Model
	Username string  `json:"username"`
	Password string  `json:"password"`
	Address  Address `json:"address" gorm:"embedded"`
}
