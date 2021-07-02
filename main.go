package main

import (
	"fmt"
	"net/http"
	"poke/utils"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// environment
	utils.ViperInit()

	// connection to db
	DBConnection()
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running")
	})

	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%s", viper.GetString("app.port")),
		),
	)
}

func DBConnection() {
	var err error

	DB, err = gorm.Open(sqlite.Open("db/poke.db"), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("fatal error db connection: %s \n", err))
	}

	fmt.Println("DB is established...")
}
