package private

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/database/postgresql/private"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type privateRepository struct {
	client postgresql.ClientInterface
}

func New(cl postgresql.ClientInterface) PrivateRepository {
	return &privateRepository{client: cl}
}

func (r privateRepository) Create(tx *sqlx.Tx, e model.Event, revision int, uid string) (int, error) {
	query := private.PrivateCreate
	var id int
	err := tx.QueryRow(query, e.Amount, e.Category, e.StoreName, e.Memo, e.Date, uid, revision, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println("[FATAL] プライベートイベントの新規作成に失敗しました")
		return -1, err
	}
	return id, nil
}

func (r privateRepository) GetAll(uid string) ([]model.EventGet, error) {
	query := private.PrivateGetAll
	privates := []model.EventGet{}
	db := r.client.GetDB()
	if err := db.Select(&privates, query, uid); err != nil {
		return nil, err
	}
	return privates, nil
}

func (r privateRepository) GetOne(uid string, id int) (model.PrivateOne, error) {
	query := private.PrivateGetOne
	private := model.PrivateOne{}
	db := r.client.GetDB()
	if err := db.Get(&private, query, id, uid); err != nil {
		return private, err
	}
	return private, nil
}

func (r privateRepository) Update(tx *sqlx.Tx, e model.EventUpdate, uid string, id, revision int) error {
	query := private.PrivateUpdate
	if _, err := tx.Exec(query, e.Amount, e.Category, e.Memo, e.StoreName, e.Date, time.Now(), revision, id, uid); err != nil {
		return err
	}
	return nil
}

func (r privateRepository) Delete(tx *sqlx.Tx, uid string, id int) error {
	query := private.PrivateDelete
	if _, err := tx.Exec(query, id, uid); err != nil {
		return err
	}
	return nil
}

func (r privateRepository) GetRevision(uid string) (int, error) {
	query := private.GetRevision
	var revision int
	db := r.client.GetDB()
	if err := db.Get(&revision, query, uid); err != nil {
		return -1, err
	}
	return revision, nil
}

func (r privateRepository) UpdateRevision(tx *sqlx.Tx, uid string) (int, error) {
	query := private.UpdateRevision
	var revision int
	if err := tx.Get(&revision, query, time.Now(), uid); err != nil {
		return -1, err
	}
	return revision, nil
}
