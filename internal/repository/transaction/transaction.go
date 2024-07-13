package transaction

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type transactionRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) Transaction(ctx context.Context, f func(tx *sqlx.Tx) error) error {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	err = f(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
