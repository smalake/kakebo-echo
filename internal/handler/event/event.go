package event

import (
	"errors"
	"kakebo-echo/internal/model"
	eventRepo "kakebo-echo/internal/repository/event"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/internal/service/event"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type eventHandler struct {
	service event.EventService
}

func New(cl postgresql.ClientInterface) EventHandler {
	repo := eventRepo.New(cl)
	db := cl.GetDB()
	transRepo := transaction.New(db)
	service := event.New(repo, transRepo)
	return &eventHandler{service: service}
}

func (h *eventHandler) Create(ctx echo.Context) error {
	// POSTから新規登録するイベントの情報を取得
	e := new(model.EventCreate)
	if err := ctx.Bind(e); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get Create Event Request: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}

	ids, err := h.service.Create(*e, uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to create event: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: ids})
}

func (h *eventHandler) GetAll(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}

	events, err := h.service.GetAll(uid)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get events: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: events})
}

func (h *eventHandler) GetOne(ctx echo.Context) error {
	idString := ctx.Param("id")
	uid := ctx.Get("uid").(string)
	if idString == "" {
		ctx.Logger().Error("[FATAL] faild to get ID")
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: errors.New("faild to get ID")})
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] ID is bad param: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusBadRequest, Error: err})
	}
	event, err := h.service.GetOne(uid, id)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to get one event: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: event})
}

func (h *eventHandler) Update(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *eventHandler) Delete(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *eventHandler) GetRevision(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: errors.New("faild to get UID")})
	}
	revision, err := h.service.GetRevision(uid)
	if err != nil || revision == -1 {
		ctx.Logger().Errorf("[FATAL] failed to get revision: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: revision})
}
