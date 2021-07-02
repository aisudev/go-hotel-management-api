package delivery

import (
	"poke/domain"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.PokeUsecase
}

func NewPokePrivateHandler(e *echo.Group, usecase domain.PokeUsecase) *Handler {
	h := Handler{usecase: usecase}

	return &h
}
