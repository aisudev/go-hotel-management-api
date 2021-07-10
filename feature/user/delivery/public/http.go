package delivery

import (
	"errors"
	"net/http"
	"poke/domain"
	"poke/middlewares"
	"poke/utils"

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

/*
	@api: CreateUser
	@method: POST
	@body: {uuid, name}
*/
func (h *Handler) CreateUserHandler(c echo.Context) error {
	type Req struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}

	reqStruct := Req{}

	if err := c.Bind(&reqStruct); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	if err := h.usecase.CreateUser(reqStruct.UUID, reqStruct.Name); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "user created.", nil, nil))
}

/*
	@api: Access User
	@method: POST
	@body: {uuid}
*/
func (h *Handler) AuthAccessHandler(c echo.Context) error {
	type Req struct {
		UUID string `json:"uuid"`
	}

	reqStruct := Req{}

	if err := c.Bind(&reqStruct); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	if isExist, err := h.usecase.ExistUser(reqStruct.UUID); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))

	} else if !isExist {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "user not existed.", nil, errors.New("user not existed.")))
	}

	jwtKey := viper.GetString("jwt.access_secret")
	accessToken, err := middlewares.GenerateToken("access", reqStruct.UUID, 3, []byte(jwtKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	jwtKey = viper.GetString("jwt.refresh_secret")
	refreshToken, err := middlewares.GenerateToken("refresh", reqStruct.UUID, 10, []byte(jwtKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}, nil))
}

/*
	@api: Refresh User
	@method: POST
	@body: {refreshToken}
*/
func (h *Handler) AuthRefreshHandler(c echo.Context) error {

	type Req struct {
		RefreshToken string `json:"refreshToken"`
	}

	reqStruct := Req{}

	if err := c.Bind(&reqStruct); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	jwtKey := viper.GetString("jwt.refresh_secret")
	uuid, err := middlewares.VerifyToken("refresh", reqStruct.RefreshToken, []byte(jwtKey))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.Response(false, "", nil, err))
	}

	jwtKey = viper.GetString("jwt.access_secret")
	accessToken, err := middlewares.GenerateToken("access", uuid, 3, []byte(jwtKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	jwtKey = viper.GetString("jwt.refresh_secret")
	refreshToken, err := middlewares.GenerateToken("refresh", uuid, 10, []byte(jwtKey))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response(false, "", nil, err))
	}

	return c.JSON(http.StatusOK, utils.Response(true, "", map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}, nil))
}
