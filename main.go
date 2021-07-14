package main

import (
	"fmt"
	"poke/middlewares"
	"poke/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	privateUserDelivery "poke/feature/user/delivery/private"
	publicUserDelivery "poke/feature/user/delivery/public"
	userRepo "poke/feature/user/repository"
	userUsecase "poke/feature/user/usecase"

	privatePokeDelivery "poke/feature/poke/delivery/private"
	publicPokeDelivery "poke/feature/poke/delivery/public"
	pokeRepo "poke/feature/poke/repository"
	pokeUsecase "poke/feature/poke/usecase"

	privateContestDelivery "poke/feature/contest/delivery"
	contestRepo "poke/feature/contest/repository"
	contestUsecase "poke/feature/contest/usecase"
)

func init() {
	// environment
	utils.ViperInit()

	// connection to db
	utils.DBConnection()

	// Auto Migration
	// utils.AutoMigration()
}

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	// Group
	private := e.Group("")
	private.Use(middlewares.AuthMiddleware())

	public := e.Group("")

	// User API
	// Private
	privateUserDelivery.NewUserPrivateHandler(private,
		userUsecase.NewUserUsecase(
			userRepo.NewUserRepository(utils.DB),
		),
	)
	// Public
	publicUserDelivery.NewUserPublicHandler(public,
		userUsecase.NewUserUsecase(
			userRepo.NewUserRepository(utils.DB),
		),
	)

	// Poke API
	// Private
	privatePokeDelivery.NewPokePrivateHandler(private,
		pokeUsecase.NewPokeUsecase(
			pokeRepo.NewPokeRepository(utils.DB),
		),
	)
	// Public
	publicPokeDelivery.NewPokePublicHandler(public,
		pokeUsecase.NewPokeUsecase(
			pokeRepo.NewPokeRepository(utils.DB),
		),
	)

	// Contest API
	// Private
	privateContestDelivery.NewContestHandler(private,
		contestUsecase.NewContestUsecase(
			contestRepo.NewContestRepositry(utils.DB),
		),
	)

	// Listen
	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%s", viper.GetString("app.port")),
		),
	)
}
