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

// Auth関連
func (a AppModel) LoginGoogle(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) LoginMail(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "ok")
}

func (a AppModel) Logout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) RegisterUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) Join(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) Revision(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) GetParentName(ctx echo.Context, group string) error {
	return nil
}

func (a AppModel) LoginCheck(ctx echo.Context) error {

	return nil
}

func (a AppModel) AuthCode(ctx echo.Context) error {

	return nil
}

func (a AppModel) ResendCode(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// Event関連
func (a AppModel) CreateEvent(ctx echo.Context) error {

	return nil
}

func (a AppModel) GetAllEvent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) GetOneEvent(ctx echo.Context, id int) error {
	return nil
}

func (a AppModel) UpdateEvent(ctx echo.Context, id int) error {
	return nil
}

func (a AppModel) DeleteEvent(ctx echo.Context, id int) error {
	// ------------- Path parameter "id" -------------

	return nil
}

// Private関連
func (a AppModel) GetAllPrivate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) GetOnePrivate(ctx echo.Context, id int) error {
	return nil
}

func (a AppModel) UpdatePrivate(ctx echo.Context, id int) error {
	return nil
}

func (a AppModel) DeletePrivate(ctx echo.Context, id int) error {
	return nil
}

// Pattern関連
func (a AppModel) GetPattern(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) RegisterPattern(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) UpdatePattern(ctx echo.Context, id int) error {
	return nil
}

func (a AppModel) DeletePattern(ctx echo.Context, id int) error {
	return nil
}

// Setting関連
func (a AppModel) GetDisplayName(ctx echo.Context) error {

	return nil
}

func (a AppModel) UpdateDisplayName(ctx echo.Context) error {

	return nil
}

func (a AppModel) Invite(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) IsParent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (a AppModel) SendMail(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// HealthCheck
func (a AppModel) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}
