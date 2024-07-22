package event

import "kakebo-echo/internal/model"

type EventService interface {
	Create(model.EventCreate, string) ([]int, error)
	GetAll(string) ([]model.EventResponse, error)
	GetOne(string, int) (model.EventGet, error)
	GetRevision(string) (int, error)
}
