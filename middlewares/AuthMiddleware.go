package middlewares

import (
	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Mock context user UUID
			c.Set("uuid", "x404")

			return next(c)
		}
	}
}
