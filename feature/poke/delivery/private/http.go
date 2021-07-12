package delivery

import (
	"net/http"
	"poke/domain"
	"poke/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.PokeUsecase
}

func NewPokePrivateHandler(e *echo.Group, usecase domain.PokeUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.GET("/poke/private", func(c echo.Context) error { return c.String(http.StatusOK, "POKE PRIVATE") })

	e.POST("/poke", h.CreatePokeHandler)
	e.GET("/poke/:poke_id", h.GetPokeHandler)
	e.GET("/poke", h.GetAllPokeHandler)

	return &h
}

func (h *Handler) CreatePokeHandler(c echo.Context) error {
	type Poke struct {
		Specie_id uint   `json:"specie_id"`
		Name      string `json:"name"`
	}

	poke := Poke{}
	if err := c.Bind(&poke); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	uuid := c.Get("uuid")
	if err := h.usecase.CreatePoke(uuid.(string), poke.Specie_id, poke.Name); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "created.", nil, nil))
}

func (h *Handler) GetAllPokeHandler(c echo.Context) error {
	uuid := c.Get("uuid")
	var pokes []map[string]interface{}
	var err error

	if pokes, err = h.usecase.GetAllPoke(uuid.(string)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", pokes, nil))
}

func (h *Handler) GetPokeHandler(c echo.Context) error {
	poke_id := c.Param("poke_id")

	poke, err := h.usecase.GetPoke(poke_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", poke, nil))
}

func (h *Handler) ChangeNamePokeHandler(c echo.Context) error {
	type Poke struct {
		Poke_id string `json:"poke_id"`
		Name    string `json:"name"`
	}

	poke := Poke{}
	if err := c.Bind(&poke); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	if err := h.usecase.ChangeNamePoke(poke.Poke_id, poke.Name); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "updated.", nil, nil))
}

func (h *Handler) TreatPokeHandler(c echo.Context) error {
	poke_id := c.Param("poke_id")

	if err := h.usecase.TreatPoke(poke_id); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "treated.", nil, nil))
}

func (h *Handler) DeletePokeHandler(c echo.Context) error {
	poke_id := c.Param("poke_id")

	if err := h.usecase.DeletePoke(poke_id); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "deleted.", nil, nil))
}
