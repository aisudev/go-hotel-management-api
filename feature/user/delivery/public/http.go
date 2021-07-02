package delivery

import (
	"net/http"
	"poke/domain"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.UserUsecase
}

func NewUserPublicHandler(e *echo.Group, usecase domain.UserUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.GET("/auth", func(c echo.Context) error { return c.String(http.StatusOK, "AUTH") })

	return &h
}
