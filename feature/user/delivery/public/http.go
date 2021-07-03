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

	e.POST("/user", h.CreateUserHandler)

	return &h
}

// CREATE USER HANDLER
func (h *Handler) CreateUserHandler(c echo.Context) error {
	reqMap := map[string]interface{}{}

	if err := c.Bind(&reqMap); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.CreateUser(reqMap["uuid"].(string), reqMap["name"].(string)); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "user created.")
}
