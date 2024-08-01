package private

import "kakebo-echo/internal/model"

type PrivateService interface {
	Create(model.EventCreate, string) ([]int, int, error)
	GetAll(string) ([]model.EventResponse, error)
	GetOne(string, int) (model.PrivateOne, error)
	Update(model.EventUpdate, string, int) (int, error)
	Delete(string, int) (int, error)
	GetRevision(string) (int, error)
}
