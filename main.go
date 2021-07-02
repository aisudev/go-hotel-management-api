package main

import (
	"fmt"
	"net/http"
	"poke/domain"
	"poke/middlewares"
	"poke/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	privateUserDelivery "poke/feature/user/delivery/private"
	publicUserDelivery "poke/feature/user/delivery/public"
	userRepo "poke/feature/user/repository"
	userUsecase "poke/feature/user/usecase"

	privatePokeDelivery "poke/feature/poke/delivery/private"
	publicPokeDelivery "poke/feature/poke/delivery/public"
	pokeRepo "poke/feature/poke/repository"
	pokeUsecase "poke/feature/poke/usecase"
)

var DB *gorm.DB

func init() {
	// environment
	utils.ViperInit()

	// connection to db
	DBConnection()

	// Auto Migration
	AutoMigration()
}

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running")
	})

	// Group
	private := e.Group("")
	private.Use(middlewares.AuthMiddleware())

	public := e.Group("")

	// User API
	// Private
	privateUserDelivery.NewUserPrivateHandler(private,
		userUsecase.NewUserUsecase(
			userRepo.NewUserRepository(DB),
		),
	)
	// Public
	publicUserDelivery.NewUserPublicHandler(public,
		userUsecase.NewUserUsecase(
			userRepo.NewUserRepository(DB),
		),
	)

	// Poke API
	// Private
	privatePokeDelivery.NewPokePrivateHandler(private,
		pokeUsecase.NewPokeUsecase(
			pokeRepo.NewPokeRepository(DB),
		),
	)
	// Public
	publicPokeDelivery.NewPokePublicHandler(public,
		pokeUsecase.NewPokeUsecase(
			pokeRepo.NewPokeRepository(DB),
		),
	)

	// Listen
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

func AutoMigration() {
	DB.AutoMigrate(&domain.User{}, &domain.Poke{})
}
