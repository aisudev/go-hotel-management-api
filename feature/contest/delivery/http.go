package delivery

import (
	"net/http"
	"poke/domain"
	"poke/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.ContestUsecase
}

func NewContestHandler(e *echo.Group, usecase domain.ContestUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.GET("/contest/:contest_id", h.GetContestHandler)
	e.GET("/contest", h.GetAllContestHandler)

	e.POST("/contest", h.ContestHandler)

	return &h
}

func (h *Handler) GetAllContestHandler(c echo.Context) error {
	contest, err := h.usecase.GetAllContest()
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", contest, nil))
}

func (h *Handler) GetContestHandler(c echo.Context) error {
	contest_id := c.Param("contest_id")

	contest, err := h.usecase.GetContest(contest_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", contest, nil))
}

func (h *Handler) ContestHandler(c echo.Context) error {
	type Challenger struct {
		Red  string `json:"red"`
		Blue string `json:"blue"`
	}

	challenger := Challenger{}
	if err := c.Bind(&challenger); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	result, err := h.usecase.Contest(challenger.Red, challenger.Blue)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", result, nil))
}
