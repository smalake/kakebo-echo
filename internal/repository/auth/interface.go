package auth

import (
	"kakebo-echo/internal/model"
)

type AuthRepository interface {
	FindUser(string) (int, error)
	Register(*model.RegisterRequest) error
	LoginCheck(string) (int, error)
	Logout() error
}
