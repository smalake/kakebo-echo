package setting

import "github.com/labstack/echo/v4"

type SettingHandler interface {
	AdminCheck(ctx echo.Context) error
}