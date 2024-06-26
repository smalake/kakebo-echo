package handler

import (
	"kakebo-echo/internal/appmodel"
	"kakebo-echo/internal/repository"
	"kakebo-echo/internal/service"
	"kakebo-echo/internal/service/auth"
	"kakebo-echo/pkg/structs"
	"kakebo-echo/pkg/user"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(am *appmodel.AppModel) *AuthHandler {
	repo := repository.NewAuthRepository(am)
	service := service.NewAuthService(repo)
	return &AuthHandler{service: service}
}

// Auth関連
func (h *AuthHandler) LoginHandler(ctx echo.Context) error {
	// res := service.Login(ctx)
	// POSTからログイン情報を取得
	u := new(user.LoginRequest)
	if err := ctx.Bind(u); err != nil {
		return ResponseHandler(ctx, structs.HttpResponse{Code: 400, Error: err})
	}

	return ResponseHandler(ctx, res)
}

func (h *AuthHandler) LoginCheckHandler(ctx echo.Context) error {
	service := auth.New(&s.appModel)
	res := service.LoginCheck(ctx)

	return ResponseHandler(ctx, res)
}
