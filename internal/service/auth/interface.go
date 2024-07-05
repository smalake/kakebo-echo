package auth

import (
	"kakebo-echo/internal/model"
)

type AuthService interface {
	Login(string) error
	Register(*model.RegisterRequest) error
	LoginCheck(string) (int, error)
	Logout() error
}
