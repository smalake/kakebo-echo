package private

import "kakebo-echo/internal/model"

type PrivateService interface {
	Create(model.PrivateCreate, string) error
	GetAll(string) ([]model.PrivateGet, error)
	GetOne(string, int) (model.PrivateGet, error)
}
