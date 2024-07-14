package private

import (
	"kakebo-echo/internal/model"
)

type PrivateRepository interface {
	Create(model.Private, string) error
	GetAll(string) ([]model.PrivateGet, error)
	GetOne(string, int) (model.PrivateGet, error)
}
