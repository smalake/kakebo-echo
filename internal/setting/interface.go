package setting

import "github.com/labstack/echo"

type IsParent struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SendMail struct {
	Parent bool `json:"parent"`
}

type SettingInterface interface {
	// (GET /display-name)
	Get(ctx echo.Context) error

	// (POST /display-name)
	Update(ctx echo.Context) error

	// (GET /invite)
	Invite(ctx echo.Context) error

	// (GET /is-parent)
	IsParent(ctx echo.Context) error

	// (POST /send-mail)
	SendMail(ctx echo.Context) error
}
