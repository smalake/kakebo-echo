package event

import (
	"kakebo-echo/internal/model"
)

type EventRepository interface {
	Create(model.Event, string) error
	GetAll(string) ([]model.EventGet, error)
	GetOne(int) (model.Event, error)
}
