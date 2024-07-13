package transaction

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Transaction(ctx context.Context, f func(tx *sqlx.Tx) error) error
}
