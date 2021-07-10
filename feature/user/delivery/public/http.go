package delivery

import (
	"net/http"
	"poke/domain"
	"poke/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Handler struct {
	usecase domain.UserUsecase
}

func NewUserPublicHandler(e *echo.Group, usecase domain.UserUsecase) *Handler {
	h := Handler{usecase: usecase}

	e.POST("/user", h.CreateUserHandler)

	e.POST("/auth/access", h.AuthAccessHandler)
	e.POST("/auth/refresh", h.AuthRefreshHandler)

	return &h
}

// CREATE USER HANDLER
func (h *Handler) CreateUserHandler(c echo.Context) error {
	type Req struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}

	reqStruct := Req{}

	if err := c.Bind(&reqStruct); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.CreateUser(reqStruct.UUID, reqStruct.Name); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "user created.")
}

// ACCESS PERMISSION HANDLER
func (h *Handler) AuthAccessHandler(c echo.Context) error {
	type Req struct {
		UUID string `json:"uuid"`
	}

	reqStruct := Req{}

	if err := c.Bind(&reqStruct); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if isExist, err := h.usecase.ExistUser(reqStruct.UUID); err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	} else if !isExist {
		return c.String(http.StatusBadRequest, "not exist")
	}

	jwtKey := viper.GetString("jwt.access_secret")
	accessToken, err := middlewares.GenerateToken("access", reqStruct.UUID, 3, []byte(jwtKey))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	jwtKey = viper.GetString("jwt.refresh_secret")
	refreshToken, err := middlewares.GenerateToken("refresh", reqStruct.UUID, 10, []byte(jwtKey))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

// REFRESH PERMISSION HANDLER
func (h *Handler) AuthRefreshHandler(c echo.Context) error {

	reqMap := map[string]interface{}{}

	if err := c.Bind(&reqMap); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	jwtKey := viper.GetString("jwt.refresh_secret")
	uuid, err := middlewares.VerifyToken("refresh", reqMap["refreshToken"].(string), []byte(jwtKey))

	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	jwtKey = viper.GetString("jwt.access_secret")
	accessToken, err := middlewares.GenerateToken("access", uuid, 3, []byte(jwtKey))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	jwtKey = viper.GetString("jwt.refresh_secret")
	refreshToken, err := middlewares.GenerateToken("refresh", uuid, 10, []byte(jwtKey))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
