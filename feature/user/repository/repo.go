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
	return nil
}

func (r *userRepository) GetUser(uuid string) (*domain.User, error) {
	return nil, nil
}

func (r *userRepository) UpdateUser(uuid string, newUser map[string]interface{}) error {
	return nil
}

func (r *userRepository) DeleteUser(uuid string) error {
	return nil
}
