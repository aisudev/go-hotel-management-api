package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{"msg": "Authorization..."})
		}
	}
}
