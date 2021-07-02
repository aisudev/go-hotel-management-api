package usecase

import "poke/domain"

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (r *userUsecase) CreateUser(user *domain.User) error {
	return nil
}

func (r *userUsecase) GetUser(uuid string) (*domain.User, error) {
	return nil, nil
}

func (r *userUsecase) UpdateUser(uuid string, newUser map[string]interface{}) error {
	return nil
}

func (r *userUsecase) DeleteUser(uuid string) error {
	return nil
}
