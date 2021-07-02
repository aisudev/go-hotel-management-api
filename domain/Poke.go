package domain

import (
	"time"

	"gorm.io/gorm"
)

type Poke struct {
	Poke_id   string         `gorm:"primaryKey;not null;unique" json:"poke_id"`
	UUID      string         `gorm:"not null" json:"-"`
	Specie_id uint           `gorm:"not null" json:"specie_id"`
	Name      string         `gorm:"not null" json:"name"`
	Exp       float32        `gorm:"not null" json:"exp"`
	Health    float32        `gorm:"not null" json:"health"`
	Damage    float32        `gorm:"not null" json:"damage"`
	CreateAt  *time.Time     `gorm:"autoCreateTime" json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}
