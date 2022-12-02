package repositories

import (
	"go-boilerplate/src/models/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserById(id int) (entities.UserEntity, error)
	ListUserByCond(conds ...interface{}) ([]entities.UserEntity, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (userRepository *UserRepository) GetUserById(id int) (entities.UserEntity, error) {
	var user entities.UserEntity
	err := userRepository.db.First(&user, id).Error
	return user, err
}
func (userRepository *UserRepository) ListUserByCond(conds ...interface{}) ([]entities.UserEntity, error) {
	var users []entities.UserEntity
	err := userRepository.db.Find(&users, conds).Error
	return users, err
}
