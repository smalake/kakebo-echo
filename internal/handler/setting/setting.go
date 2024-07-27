package setting

import (
	"errors"
	settingRepo "kakebo-echo/internal/repository/setting"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/internal/service/setting"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/structs"
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
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}

	isAdmin, err := h.service.AdminCheck(uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get check is admin: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"admin": isAdmin}})

}