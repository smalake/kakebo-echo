package event

import "kakebo-echo/internal/appmodel"

type EventRepository interface {
	Create() error
	GetAll() error
}

type eventRepository struct {
	appModel appmodel.AppModel
}

func New(am appmodel.AppModel) EventRepository {
	return &eventRepository{appModel: am}
}
