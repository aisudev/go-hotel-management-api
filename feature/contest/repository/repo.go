package repository

import (
	"poke/domain"

	"gorm.io/gorm"
)

type contestRepository struct {
	db *gorm.DB
}

func NewContestRepositry(db *gorm.DB) domain.ContestRepository {
	return &contestRepository{
		db: db,
	}
}

func (r *contestRepository) CreateContest(contest *domain.Contest) error {
	return r.db.Create(contest).Error
}

func (r *contestRepository) GetContest(contest_id string) (*domain.Contest, error) {
	var contest domain.Contest

	if err := r.db.Where("contest_id = ?", contest_id).Find(&contest).Error; err != nil {
		return nil, err
	}

	return &contest, nil
}

func (r *contestRepository) GetAllContest() ([]domain.Contest, error) {
	var contest []domain.Contest

	if err := r.db.Find(&contest).Error; err != nil {
		return nil, err
	}

	return contest, nil
}
