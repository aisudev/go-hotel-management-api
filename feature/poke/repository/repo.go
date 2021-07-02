package repository

import (
	"poke/domain"

	"gorm.io/gorm"
)

type pokeRepository struct {
	db *gorm.DB
}

func NewPokeRepository(db *gorm.DB) domain.PokeRepository {
	return &pokeRepository{
		db: db,
	}
}

func (r *pokeRepository) GetPoke(poke_id string) (*domain.Poke, error) {
	return nil, nil
}

func (r *pokeRepository) GetAllPoke(uuid string) ([]domain.Poke, error) {
	return nil, nil
}

func (r *pokeRepository) CreatePoke(poke *domain.Poke) error {
	return nil
}

func (r *pokeRepository) UpdatePoke(newPoke map[string]interface{}) error {
	return nil
}

func (r *pokeRepository) DeletePoke(poke_id string) error {
	return nil
}
