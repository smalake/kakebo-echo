package event

import "kakebo-echo/internal/model"

type EventService interface {
	Create(model.EventCreate, string) error
	GetAll(string) ([]model.EventGet, error)
	GetOne(string, int) (model.EventGet, error)
	GetRevision(string) (int, error)
}
