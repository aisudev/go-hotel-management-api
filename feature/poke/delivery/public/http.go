package delivery

import (
	"net/http"
	"poke/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.PokeUsecase
}

func NewPokePublicHandler(e *echo.Group, usecase domain.PokeUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.GET("/api/poke/:offset/:limit", h.GetMorePokeAPIHandler)
	e.GET("/api/poke/img/:name", h.GetPokeImageAPIHandler)
	e.GET("/api/poke/:name", h.GetPokeAPIHandler)

	return &h
}

// GET MORE POKE API
func (h *Handler) GetMorePokeAPIHandler(c echo.Context) error {
	limitString := c.Param("limit")
	limit, _ := strconv.ParseInt(limitString, 10, 64)

	offsetString := c.Param("offset")
	offset, _ := strconv.ParseInt(offsetString, 10, 64)

	pokemons, err := h.usecase.GetMorePokeAPI(int(limit), int(offset))

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, pokemons)
}

// SEARCH POKE API
func (h *Handler) GetPokeAPIHandler(c echo.Context) error {
	name := c.Param("name")

	pokemons, err := h.usecase.GetPokeAPI(name)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, pokemons)
}

// GET POKE IMG
func (h *Handler) GetPokeImageAPIHandler(c echo.Context) error {
	name := c.Param("name")

	images, err := h.usecase.GetPokeImageAPI(name)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, images)
}
