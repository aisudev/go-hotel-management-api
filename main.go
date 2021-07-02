package main

import (
	"fmt"
	"net/http"
	"poke/utils"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mtslzr/pokeapi-go"
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

		r, _ := pokeapi.Resource("pokemon", 0, 10)

		results := r.Results
		var pokemons []map[string]interface{}

		for _, poke := range results {
			result, _ := pokeapi.Pokemon(poke.Name)

			pokemons = append(pokemons, map[string]interface{}{
				"id":      result.ID,
				"name":    result.Name,
				"imgUrls": result.Sprites,
			})
		}

		return c.JSON(http.StatusOK, pokemons)
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
