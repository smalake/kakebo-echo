package auth

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Login(echo.Context) error
	Register(echo.Context) error
	LoginCheck(echo.Context) error
	Logout(echo.Context) error
}
