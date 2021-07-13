package usecase

import (
	"errors"
	"poke/domain"
)

type userUsecase struct {
	repo domain.UserRepository
}

var UserUsecaseInstance *userUsecase

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	UserUsecaseInstance = &userUsecase{
		repo: repo,
	}

	return UserUsecaseInstance
}

// CREATE USER USECASE
func (u *userUsecase) CreateUser(uuid string, name string) error {
	user := domain.User{
		UUID:        uuid,
		Name:        name,
		Balance:     0,
		Exp:         0,
		DefaultPoke: "",
	}

	return u.repo.CreateUser(&user)
}

// GET USER USECASE
func (u *userUsecase) GetUser(uuid string) (*domain.User, error) {
	return u.repo.GetUser(uuid)
}

// UPDATE USER USECASE
func (u *userUsecase) UpdateUser(uuid string, newUser map[string]interface{}) error {
	return u.repo.UpdateUser(uuid, newUser)
}

// DELETE USER USECASE
func (u *userUsecase) DeleteUser(uuid string) error {
	return u.repo.DeleteUser(uuid)
}

// EXIST USER USECASE
func (u *userUsecase) ExistUser(uuid string) (bool, error) {
	var user *domain.User
	var err error

	if user, err = u.GetUser(uuid); err != nil {
		return false, nil
	}

	if user.UUID == "" {
		return false, nil
	}

	return true, nil
}

func (u *userUsecase) Withdraw(uuid string, amount float32) error {
	user, err := u.repo.GetUser(uuid)
	if err != nil {
		return err
	}

	balance := user.Balance - amount
	if balance <= 0 {
		return errors.New("not enough balance.")
	}

	newData := map[string]interface{}{
		"balance": balance,
	}

	return u.UpdateUser(uuid, newData)
}

func (u *userUsecase) Deposit(uuid string, amount float32) error {
	user, err := u.repo.GetUser(uuid)
	if err != nil {
		return err
	}

	if amount <= 0 {
		return errors.New("not valid amount.")
	}

	newData := map[string]interface{}{
		"balance": user.Balance + amount,
	}

	return u.UpdateUser(uuid, newData)
}
