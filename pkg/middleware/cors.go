package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CORSの設定
func NewCors() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			c.Response().Writer.Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
			c.Response().Header().Set("Access-Control-Max-Age", "12h0m0s")
			c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Authorization")
			c.Response().Header().Set("Access-Control-Expose-Headers", "Content-Length")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request().Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	}
}
