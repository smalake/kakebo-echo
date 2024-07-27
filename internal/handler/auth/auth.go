package auth

import (
	"kakebo-echo/internal/model"
	authRepo "kakebo-echo/internal/repository/auth"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/internal/service/auth"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/errors"
	"kakebo-echo/pkg/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	service auth.AuthService
}

func New(cl postgresql.ClientInterface) AuthHandler {
	repo := authRepo.New(cl)
	db := cl.GetDB()
	transRepo := transaction.New(db)
	service := auth.New(repo, transRepo)
	return &authHandler{service: service}
}

// Auth関連
func (h *authHandler) Login(ctx echo.Context) error {
	// POSTからログイン情報を取得
	u := new(model.LoginRequest)
	if err := ctx.Bind(u); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get Login Request: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	err := h.service.Login(u.Uid)
	if err != nil {
		if err == errors.ErrUserNotFound {
			ctx.Logger().Errorf("[FATAL] User not found in DB: %+v", err)
			return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusUnauthorized, Error: errors.ErrUserNotFound})
		}
		ctx.Logger().Errorf("[FATAL] failed to find user: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *authHandler) Register(ctx echo.Context) error {
	// POSTから登録情報を取得
	r := new(model.RegisterRequest)
	if err := ctx.Bind(r); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get Register User Request: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	err := h.service.Register(r)
	if err != nil {
		if err == errors.ErrUserAlreadyExist {
			return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusConflict, Error: err})
		}
		ctx.Logger().Errorf("[FATAL] failed to create user: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *authHandler) LoginCheck(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return errors.InternalServerError
	}
	admin, err := h.service.LoginCheck(uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to check login: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}

	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"admin": admin}})
}

func (h *authHandler) Logout(ctx echo.Context) error {
	return nil
}
