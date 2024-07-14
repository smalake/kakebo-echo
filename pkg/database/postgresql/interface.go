package postgresql

import "github.com/jmoiron/sqlx"

type ClientInterface interface {
	GetDB() *sqlx.DB
}
