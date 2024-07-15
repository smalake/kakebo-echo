package event

import (
	"kakebo-echo/internal/model"

	"github.com/jmoiron/sqlx"
)

type EventRepository interface {
	GetIDs(string) (int, int, error)
	Create(*sqlx.Tx, model.Event, int, int, int) error
	GetAll(string) ([]model.EventGet, error)
	GetOne(string, int) (model.EventGet, error)
	GetRevision(int) (int, error)
	UpdateRevision(*sqlx.Tx, int) (int, error)
}
