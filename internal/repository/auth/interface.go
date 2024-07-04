package auth

import (
	"kakebo-echo/internal/appmodel"
	"kakebo-echo/internal/model"
)

type AuthRepository interface {
	FindUser(string) (int, error)
	Register(*model.RegisterRequest) error
	LoginCheck(string) (int, error)
	Logout() error
}

type authRepository struct {
	appModel appmodel.AppModel
}

func New(am appmodel.AppModel) AuthRepository {
	return &authRepository{appModel: am}
}
