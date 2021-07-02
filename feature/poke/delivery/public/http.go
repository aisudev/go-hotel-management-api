package delivery

import (
	"net/http"
	"poke/domain"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.PokeUsecase
}

func NewPokePublicHandler(e *echo.Group, usecase domain.PokeUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.GET("/poke/public", func(c echo.Context) error { return c.String(http.StatusOK, "POKE PUBLIC") })

	return &h
}
