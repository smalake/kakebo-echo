package event

import (
	"kakebo-echo/internal/model"

	"github.com/jmoiron/sqlx"
)

type EventRepository interface {
	GetIDs(string) (int, int, error)
	Create(*sqlx.Tx, model.Event, int, int, int) (int, error)
	GetAll(string) ([]model.EventGet, error)
	GetOne(string, int) (model.EventOne, error)
	Update(*sqlx.Tx, model.EventUpdate, string, int, int, int) error
	Delete(*sqlx.Tx, int, int) error
	GetRevision(int) (int, error)
	UpdateRevision(*sqlx.Tx, int) (int, error)
}
