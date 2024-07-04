package auth

import (
	"database/sql"
	"kakebo-echo/internal/model"
	"kakebo-echo/internal/repository/auth"
)

type LoginCheck struct {
	GroupAdmin sql.NullInt32 `db:"group_admin"`
}

type AuthService interface {
	Login(string) error
	Register(*model.RegisterRequest) error
	LoginCheck(string) (int, error)
	Logout() error
}

type authService struct {
	repo auth.AuthRepository
}

func New(repo auth.AuthRepository) AuthService {
	return &authService{repo: repo}
}
