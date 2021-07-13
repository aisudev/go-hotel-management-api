package utils

import (
	"fmt"
	"poke/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error

	DB, err = gorm.Open(sqlite.Open("db/poke.db"), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("fatal error db connection: %s \n", err))
	}

	fmt.Println("DB is established...")
}

func AutoMigration() {
	DB.AutoMigrate(&domain.User{}, &domain.Poke{})
}
