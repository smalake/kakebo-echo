package event

import "kakebo-echo/internal/model"

type EventService interface {
	Create(model.EventCreate, string) ([]int, error)
	GetAll(string) ([]model.EventResponse, error)
	GetOne(string, int) (model.EventOne, error)
	Update(model.EventUpdate, string, int) (int, error)
	GetRevision(string) (int, error)
}
