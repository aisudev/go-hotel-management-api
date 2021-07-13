package repository

import (
	"poke/domain"

	"gorm.io/gorm"
)

type pokeRepository struct {
	db *gorm.DB
}

var PokeRepoInstance *pokeRepository

func NewPokeRepository(db *gorm.DB) domain.PokeRepository {
	PokeRepoInstance = &pokeRepository{
		db: db,
	}

	return PokeRepoInstance
}

func (r *pokeRepository) GetPoke(poke_id string) (*domain.Poke, error) {
	var poke domain.Poke

	if err := r.db.Where("poke_id = ?", poke_id).Find(&poke).Error; err != nil {
		return nil, err
	}

	return &poke, nil
}

func (r *pokeRepository) GetAllPoke(uuid string) ([]domain.Poke, error) {
	var pokes []domain.Poke

	if err := r.db.Where("uuid = ?", uuid).Find(&pokes).Error; err != nil {
		return nil, err
	}

	return pokes, nil
}

func (r *pokeRepository) CreatePoke(poke *domain.Poke) error {
	return r.db.Create(poke).Error
}

func (r *pokeRepository) UpdatePoke(poke_id string, newPoke map[string]interface{}) error {
	var poke domain.Poke
	return r.db.Model(&poke).Where("poke_id = ?", poke_id).Updates(newPoke).Error
}

func (r *pokeRepository) DeletePoke(poke_id string) error {
	return nil
}
