package event

import (
	"kakebo-echo/internal/model"

	"github.com/jmoiron/sqlx"
)

type EventRepository interface {
	GetGroupID(string) (int, error)
	Create(*sqlx.Tx, model.Event, int) error
	GetAll(string) ([]model.EventGet, error)
	GetOne(string, int) (model.EventGet, error)
	GetRevision(int) (int, error)
}
