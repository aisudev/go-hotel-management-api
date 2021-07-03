package repository

import (
	"poke/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *domain.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetUser(uuid string) (*domain.User, error) {
	var user domain.User

	if err := r.db.Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(uuid string, newUser map[string]interface{}) error {
	var user domain.User

	if err := r.db.Model(&user).Where("uuid = ?", uuid).Updates(newUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(uuid string) error {
	var user domain.User

	if err := r.db.Unscoped().Where("uuid = ?", uuid).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
