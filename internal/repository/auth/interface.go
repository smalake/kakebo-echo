package auth

import (
	"kakebo-echo/internal/model"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindUser(string) (int, error)
	CreateGroup(*sqlx.Tx) (int, error)
	CreateUser(*sqlx.Tx, *model.RegisterRequest, int) (int, error)
	CreateRevision(*sqlx.Tx, int) error
	LoginCheck(string) (int, error)
	Logout() error
}
