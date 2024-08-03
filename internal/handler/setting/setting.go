package setting

import (
	"errors"
	settingRepo "kakebo-echo/internal/repository/setting"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/internal/service/setting"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/structs"
	"kakebo-echo/pkg/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type settingHandler struct {
	service setting.SettingService
}

func New(cl postgresql.ClientInterface) SettingHandler {
	repo := settingRepo.New(cl)
	db := cl.GetDB()
	transRepo := transaction.New(db)
	service := setting.New(repo, transRepo)
	return &settingHandler{service: service}
}

func (h settingHandler) AdminCheck(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}

	isAdmin, err := h.service.AdminCheck(uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get check is admin: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"admin": isAdmin}})

}

func (h settingHandler) GetName(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}
	name, err := h.service.GetName(uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get display name: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"name": name}})

}

func (h settingHandler) UpdateName(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}
	// POSTから新規登録するイベントの情報を取得
	type updateName struct {
		Name string `json:"name"`
	}
	e := new(updateName)
	if err := ctx.Bind(e); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get Update Name Request: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	err := h.service.UpdateName(uid, e.Name)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to update display name: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})

}
