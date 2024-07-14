package private

import "github.com/labstack/echo/v4"

type PrivateHandler interface {
	Create(echo.Context) error
	GetAll(echo.Context) error
	GetOne(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}
