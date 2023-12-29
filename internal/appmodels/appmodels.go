package appmodels

import (
	"kakebo-echo/pkg/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AppModel struct {
	mysqlCli *mysql.Client
}

func New(mysqlCli *mysql.Client) *AppModel {
	return &AppModel{mysqlCli: mysqlCli}
}

func (a AppModel) AuthCode(ctx echo.Context) error {

	return nil
}

// GetDisplayName converts echo context to params.
func (a AppModel) GetDisplayName(ctx echo.Context) error {

	return nil
}

// UpdateDisplayName converts echo context to params.
func (a AppModel) UpdateDisplayName(ctx echo.Context) error {

	return nil
}

// GetAllEvent converts echo context to params.
func (a AppModel) GetAllEvent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// CreateEvent converts echo context to params.
func (a AppModel) CreateEvent(ctx echo.Context) error {

	return nil
}

// DeleteEvent converts echo context to params.
func (a AppModel) DeleteEvent(ctx echo.Context, id int) error {
	// ------------- Path parameter "id" -------------

	return nil
}

// GetOneEvent converts echo context to params.
func (a AppModel) GetOneEvent(ctx echo.Context, id int) error {
	return nil
}

// UpdateEvent converts echo context to params.
func (a AppModel) UpdateEvent(ctx echo.Context, id int) error {
	return nil
}

// GetParentName converts echo context to params.
func (a AppModel) GetParentName(ctx echo.Context, group string) error {
	return nil
}

// HealthCheck converts echo context to params.
func (a AppModel) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// Invite converts echo context to params.
func (a AppModel) Invite(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// IsParent converts echo context to params.
func (a AppModel) IsParent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// Join converts echo context to params.
func (a AppModel) Join(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// LoginCheck converts echo context to params.
func (a AppModel) LoginCheck(ctx echo.Context) error {

	return nil
}

// LoginGoogle converts echo context to params.
func (a AppModel) LoginGoogle(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// LoginMail converts echo context to params.
func (a AppModel) LoginMail(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "ok")
}

// Logout converts echo context to params.
func (a AppModel) Logout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// GetPattern converts echo context to params.
func (a AppModel) GetPattern(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// RegisterPattern converts echo context to params.
func (a AppModel) RegisterPattern(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// DeletePattern converts echo context to params.
func (a AppModel) DeletePattern(ctx echo.Context, id int) error {
	return nil
}

// UpdatePattern converts echo context to params.
func (a AppModel) UpdatePattern(ctx echo.Context, id int) error {
	return nil
}

// GetAllPrivate converts echo context to params.
func (a AppModel) GetAllPrivate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// DeletePrivate converts echo context to params.
func (a AppModel) DeletePrivate(ctx echo.Context, id int) error {
	return nil
}

// GetOnePrivate converts echo context to params.
func (a AppModel) GetOnePrivate(ctx echo.Context, id int) error {
	return nil
}

// UpdatePrivate converts echo context to params.
func (a AppModel) UpdatePrivate(ctx echo.Context, id int) error {
	return nil
}

// RegisterUser converts echo context to params.
func (a AppModel) RegisterUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// ResendCode converts echo context to params.
func (a AppModel) ResendCode(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// Revision converts echo context to params.
func (a AppModel) Revision(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// SendMail converts echo context to params.
func (a AppModel) SendMail(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}
