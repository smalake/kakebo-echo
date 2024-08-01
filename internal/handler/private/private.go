package private

import (
	"errors"
	"kakebo-echo/internal/model"
	privateRepo "kakebo-echo/internal/repository/private"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/internal/service/private"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/structs"
	"kakebo-echo/pkg/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type privateHandler struct {
	service private.PrivateService
}

func New(cl postgresql.ClientInterface) PrivateHandler {
	repo := privateRepo.New(cl)
	db := cl.GetDB()
	transRepo := transaction.New(db)
	service := private.New(repo, transRepo)
	return &privateHandler{service: service}
}

func (h *privateHandler) Create(ctx echo.Context) error {
	// POSTから新規登録するイベントの情報を取得
	e := new(model.EventCreate)
	if err := ctx.Bind(e); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get Create Private Request: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}

	ids, revision, err := h.service.Create(*e, uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to create private: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"ids": ids, "revision": revision}})
}

func (h *privateHandler) GetAll(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}

	privates, err := h.service.GetAll(uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get privates: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"privates": privates}})
}

func (h *privateHandler) GetOne(ctx echo.Context) error {
	idString := ctx.Param("id")
	uid := ctx.Get("uid").(string)
	if idString == "" {
		ctx.Logger().Error("[FATAL] faild to get ID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: errors.New("faild to get ID")})
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] ID is bad param: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	private, err := h.service.GetOne(uid, id)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get one private: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"private": private}})
}

func (h *privateHandler) Update(ctx echo.Context) error {
	// POSTから新規登録するイベントの情報を取得
	e := new(model.EventUpdate)
	if err := ctx.Bind(e); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get Update Event Request: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}

	idString := ctx.Param("id")
	uid := ctx.Get("uid").(string)
	if idString == "" {
		ctx.Logger().Error("[FATAL] faild to get ID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: errors.New("faild to get ID")})
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] ID is bad param: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	revision, err := h.service.Update(*e, uid, id)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to update event: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"revision": revision}})
}

func (h *privateHandler) Delete(ctx echo.Context) error {
	idString := ctx.Param("id")
	uid := ctx.Get("uid").(string)
	if idString == "" {
		ctx.Logger().Error("[FATAL] faild to get ID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: errors.New("faild to get ID")})
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] ID is bad param: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	revision, err := h.service.Delete(uid, id)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to delete event: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"revision": revision}})
}

func (h *privateHandler) GetRevision(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}
	revision, err := h.service.GetRevision(uid)
	if err != nil || revision == -1 {
		ctx.Logger().Errorf("[FATAL] failed to get private revision: %+v", err)
		return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return util.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: map[string]interface{}{"revision": revision}})
}
