package event

import (
	"kakebo-echo/internal/appmodel"
	"kakebo-echo/internal/model"
	eventRepo "kakebo-echo/internal/repository/event"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/internal/service/event"
	"kakebo-echo/pkg/errors"
	"kakebo-echo/pkg/structs"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type eventHandler struct {
	service event.EventService
}

func New(am appmodel.AppModel) EventHandler {
	repo := eventRepo.New(am)
	transRepo := transaction.New(am.PsgrCli.DB)
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
		return errors.InternalServerError
	}

	if err := h.service.Create(*e, uid); err != nil {
		ctx.Logger().Errorf("[FATAL] failed to create event: %+v", err)
		return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusInternalServerError, Error: err})
	}
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *eventHandler) GetAll(ctx echo.Context) error {
	uid := ctx.Get("uid").(string)
	if uid == "" {
		ctx.Logger().Error("[FATAL] faild to get UID")
		return errors.InternalServerError
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
	if idString == "" {
		ctx.Logger().Error("[FATAL] faild to get ID")
		return errors.BadRequest
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.Logger().Error("[FATAL] ID is bad param")
		return errors.BadRequest
	}
	event, err := h.service.GetOne(id)
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK, Data: event})
}

func (h *eventHandler) Update(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *eventHandler) Delete(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}
