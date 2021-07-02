package delivery

import (
	"poke/domain"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.UserUsecase
}

func NewUserPrivateHandler(e *echo.Group, usecase domain.UserUsecase) *Handler {
	h := Handler{usecase: usecase}

	return &h
}
