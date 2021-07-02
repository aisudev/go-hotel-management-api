package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UUID        string         `gorm:"primaryKey;not null;unique" json:"uuid"`
	Name        string         `gorm:"not null" json:"name"`
	Exp         float32        `gorm:"not null" json:"exp"`
	Balance     float32        `gorm:"not null" json:"balance"`
	DefaultPoke string         `gorm:"" json:"default_poke"`
	CreateAt    *time.Time     `gorm:"autoCreateTime" json:"-"`
	DeleteAt    gorm.DeletedAt `json:"-"`

	Pokes []Poke `gorm:"foreignKey:UUID" json:""`
}
