package event

import (
	"kakebo-echo/internal/model"

	"github.com/jmoiron/sqlx"
)

type EventRepository interface {
	GetGroupID(*sqlx.Tx, string) (int, error)
	Create(*sqlx.Tx, model.Event, int) error
	GetAll(string) ([]model.EventGet, error)
	GetOne(string, int) (model.EventGet, error)
}
