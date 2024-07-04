package handler

import (
	"kakebo-echo/internal/appmodel"
	eventRepo "kakebo-echo/internal/repository/event"
	"kakebo-echo/internal/service/event"
	"kakebo-echo/pkg/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EventHandler interface {
	Create(echo.Context) error
	GetAll(echo.Context) error
	GetOne(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type eventHandler struct {
	service event.EventService
}

func New(am appmodel.AppModel) EventHandler {
	repo := eventRepo.New(am)
	service := event.New(repo)
	return &eventHandler{service: service}
}

func (h *eventHandler) Create(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *eventHandler) GetAll(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *eventHandler) GetOne(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *eventHandler) Update(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}

func (h *eventHandler) Delete(ctx echo.Context) error {
	return structs.ResponseHandler(ctx, structs.HttpResponse{Code: http.StatusOK})
}
