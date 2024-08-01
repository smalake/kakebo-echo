package private

import (
	"kakebo-echo/internal/model"

	"github.com/jmoiron/sqlx"
)

type PrivateRepository interface {
	Create(*sqlx.Tx, model.Event, int, string) (int, error)
	GetAll(string) ([]model.EventGet, error)
	GetOne(string, int) (model.PrivateOne, error)
	Update(*sqlx.Tx, model.EventUpdate, string, int, int) error
	Delete(*sqlx.Tx, string, int) error
	GetRevision(string) (int, error)
	UpdateRevision(*sqlx.Tx, string) (int, error)
}
