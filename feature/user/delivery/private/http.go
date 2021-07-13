package delivery

import (
	"net/http"
	"poke/domain"
	"poke/utils"

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

	e.PUT("/user/withdraw", h.WithdrawHandler)
	e.PUT("/user/deposit", h.DepositHandler)

	return &h
}

/*
	@api:GetUser
	@method:GET
*/
func (h *Handler) GetUserHandler(c echo.Context) error {
	uuid := c.Get("uuid")

	user, err := h.usecase.GetUser(uuid.(string))

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", user, nil))
}

/*
	@api:UpdateUser
	@method:PUT
	@body: {name, default_poke}
*/
func (h *Handler) UpdateUserHandler(c echo.Context) error {
	type Req struct {
		Name        string `json:"name"`
		DefaultPoke string `json:"default_poke"`
	}

	reqStruct := Req{}

	if err := c.Bind(&reqStruct); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	reqMap := map[string]interface{}{}
	for k, v := range map[string]string{"name": reqStruct.Name, "default_poke": reqStruct.DefaultPoke} {
		if v != "" {
			reqMap[k] = v
		}
	}

	uuid := c.Get("uuid")

	if err := h.usecase.UpdateUser(uuid.(string), reqMap); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "user updated.", nil, nil))
}

/*
	@api: Withdraw
	@method: PUT
*/
func (h *Handler) WithdrawHandler(c echo.Context) error {
	type Req struct {
		Balance float32 `json:"balance"`
	}

	reqStruct := Req{}
	if err := c.Bind(&reqStruct); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	uuid := c.Get("uuid")
	if err := h.usecase.Withdraw(uuid.(string), reqStruct.Balance); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "withdraw success.", nil, nil))
}

/*
	@api: Deposit
	@method: PUT
*/
func (h *Handler) DepositHandler(c echo.Context) error {
	type Req struct {
		Balance float32 `json:"balance"`
	}

	reqStruct := Req{}
	if err := c.Bind(&reqStruct); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	uuid := c.Get("uuid")
	if err := h.usecase.Deposit(uuid.(string), reqStruct.Balance); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "deposit success.", nil, nil))
}

/*
	@api: DeleteUser
	@method: Delete
*/
func (h *Handler) DeleteUserHandler(c echo.Context) error {
	uuid := c.Get("uuid")

	if err := h.usecase.DeleteUser(uuid.(string)); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "user deleted.", nil, nil))
}
