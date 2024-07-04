package event

import "kakebo-echo/internal/repository/event"

type EventService interface {
	Create() error
	GetAll() error
}

type eventService struct {
	repo event.EventRepository
}

func New(repo event.EventRepository) EventService {
	return &eventService{repo: repo}
}
