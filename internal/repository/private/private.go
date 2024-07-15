package private

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/database/postgresql/private"
	"log"
	"time"
)

type privateRepository struct {
	client postgresql.ClientInterface
}

func New(cl postgresql.ClientInterface) PrivateRepository {
	return &privateRepository{client: cl}
}

func (r privateRepository) Create(e model.Private, uid string) error {
	// トランザクション開始
	db := r.client.GetDB()
	tx, err := db.Beginx()
	if err != nil {
		log.Println("[FATAL] failed to transaction")
		return err
	}
	// イベントの追加
	query := private.PrivateCreate
	_, err = tx.Exec(query, e.Amount, e.Category, e.StoreName, e.Memo, e.Date, time.Now(), time.Now())
	if err != nil {
		_ = tx.Rollback()
		log.Println("[FATAL] イベントの新規作成に失敗しました")
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("[FATAL] failed to commit")
		_ = tx.Rollback()
		return err
	}
	return nil
}

func (r privateRepository) GetAll(uid string) ([]model.PrivateGet, error) {
	query := private.PrivateGetAll
	privates := []model.PrivateGet{}
	db := r.client.GetDB()
	if err := db.Select(&privates, query, uid); err != nil {
		return nil, err
	}
	return privates, nil
}

func (r privateRepository) GetOne(uid string, id int) (model.PrivateGet, error) {
	query := private.PrivateGetOne
	private := model.PrivateGet{}
	db := r.client.GetDB()
	if err := db.Get(&private, query, uid, id); err != nil {
		return private, err
	}
	return private, nil
}
