package service

import (
	"kakebo-echo/internal/appmodels"
	"kakebo-echo/internal/service/auth"

	"github.com/labstack/echo/v4"
)

type Service struct {
	appModel appmodels.AppModel
}

func New(am *appmodels.AppModel) *Service {
	return &Service{appModel: *am}
}

// Auth関連
func (s Service) LoginGoogleHandler(ctx echo.Context) error {
	service := auth.New(&s.appModel)
	res := service.LoginGoogle(ctx)

	return ResponseHandler(ctx, res)
}

func (s Service) LoginMailHandler(ctx echo.Context) error {
	service := auth.New(&s.appModel)
	res := service.LoginMail(ctx)

	return ResponseHandler(ctx, res)
}

func (s Service) Logout(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) RegisterUserHandler(ctx echo.Context) error {
	service := auth.New(&s.appModel)
	res := service.RegisterUser(ctx)

	return ResponseHandler(ctx, res)
}

func (s Service) Join(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) Revision(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) GetParentName(ctx echo.Context, group string) error {
	return nil
}

func (s Service) LoginCheck(ctx echo.Context) error {

	return nil
}

func (s Service) AuthCode(ctx echo.Context) error {

	return nil
}

func (s Service) ResendCode(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// Event関連
func (s Service) CreateEvent(ctx echo.Context) error {

	return nil
}

func (s Service) GetAllEvent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) GetOneEvent(ctx echo.Context, id int) error {
	return nil
}

func (s Service) UpdateEvent(ctx echo.Context, id int) error {
	return nil
}

func (s Service) DeleteEvent(ctx echo.Context, id int) error {
	// ------------- Path parameter "id" -------------

	return nil
}

// Private関連
func (s Service) GetAllPrivate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) GetOnePrivate(ctx echo.Context, id int) error {
	return nil
}

func (s Service) UpdatePrivate(ctx echo.Context, id int) error {
	return nil
}

func (s Service) DeletePrivate(ctx echo.Context, id int) error {
	return nil
}

// Pattern関連
func (s Service) GetPattern(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) RegisterPattern(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) UpdatePattern(ctx echo.Context, id int) error {
	return nil
}

func (s Service) DeletePattern(ctx echo.Context, id int) error {
	return nil
}

// Setting関連
func (s Service) GetDisplayName(ctx echo.Context) error {

	return nil
}

func (s Service) UpdateDisplayName(ctx echo.Context) error {

	return nil
}

func (s Service) Invite(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) IsParent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

func (s Service) SendMail(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}

// HealthCheck
func (s Service) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	return err
}
