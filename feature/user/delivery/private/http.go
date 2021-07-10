package delivery

import (
	"net/http"
	"poke/domain"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.UserUsecase
}

func NewUserPrivateHandler(e *echo.Group, usecase domain.UserUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.GET("/user", h.GetUserHandler)
	e.PUT("/user", h.UpdateUserHandler)
	e.DELETE("/user", h.DeleteUserHandler)

	return &h
}

// GET USER HANDLER
func (h *Handler) GetUserHandler(c echo.Context) error {
	uuid := c.Get("uuid")

	user, err := h.usecase.GetUser(uuid.(string))

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// UPDATE USER HANDLER
func (h *Handler) UpdateUserHandler(c echo.Context) error {
	type Reg struct {
		Name        string `json:"name"`
		DefaultPoke string `json:"default_poke"`
	}

	reqStruct := Reg{}

	if err := c.Bind(&reqStruct); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	reqMap := map[string]interface{}{}
	for k, v := range map[string]string{"name": reqStruct.Name, "default_poke": reqStruct.DefaultPoke} {
		if v != "" {
			reqMap[k] = v
		}
	}

	uuid := c.Get("uuid")

	if err := h.usecase.UpdateUser(uuid.(string), reqMap); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "user updated.")
}

// DELETE USER HANDLER
func (h *Handler) DeleteUserHandler(c echo.Context) error {
	uuid := c.Get("uuid")

	if err := h.usecase.DeleteUser(uuid.(string)); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "user deleted")
}
