package usecase

import "poke/domain"

type pokeUsecase struct {
	repo domain.PokeRepository
}

func NewPokeUsecase(repo domain.PokeRepository) domain.PokeUsecase {
	return &pokeUsecase{
		repo: repo,
	}
}

func (u *pokeUsecase) GetMorePokeAPI(offset int, limit int) ([]map[string]interface{}, error) {
	return nil, nil
}

func (u *pokeUsecase) GetPokeAPI(filter interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func (u *pokeUsecase) GetPokeImageAPI(filter interface{}) ([]string, error) {
	return nil, nil
}

func (u *pokeUsecase) CreatePoke(specie_id uint, name string) error {
	return nil
}

func (u *pokeUsecase) GetPoke(poke_id string) (map[string]interface{}, error) {
	return nil, nil
}

func (u *pokeUsecase) GetAllPoke() ([]map[string]interface{}, error) {
	return nil, nil
}

func (u *pokeUsecase) UpdatePoke(poke_id string, newPoke map[string]interface{}) error {
	return nil
}

func (u *pokeUsecase) DeletePoke(poke_id string) error {
	return nil
}

func (u *pokeUsecase) VerifyPoke(poke_id string) error {
	return nil
}
