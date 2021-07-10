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
