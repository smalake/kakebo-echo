package event

import "github.com/labstack/echo/v4"

type EventHandler interface {
	Create(echo.Context) error
	GetAll(echo.Context) error
	GetOne(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}
