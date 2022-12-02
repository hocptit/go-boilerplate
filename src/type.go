package src

import "gorm.io/gorm"

type DB struct {
	connection *gorm.DB
}
