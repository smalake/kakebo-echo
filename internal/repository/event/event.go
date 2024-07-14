package event

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/database/postgresql/event"
	"log"
	"time"
)

type eventRepository struct {
	client postgresql.ClientInterface
}

func New(cl postgresql.ClientInterface) EventRepository {
	return &eventRepository{client: cl}
}

func (r eventRepository) Create(e model.Event, uid string) error {
	// トランザクション開始
	db := r.client.GetDB()
	tx, err := db.Beginx()
	if err != nil {
		log.Println("[FATAL] failed to transaction")
		return err
	}
	// イベントの追加
	query := event.EventCreate
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

func (r eventRepository) GetAll(uid string) ([]model.EventGet, error) {
	query := event.EventGetAll
	events := []model.EventGet{}
	db := r.client.GetDB()
	if err := db.Get(&events, query, uid); err != nil {
		return nil, err
	}
	return events, nil
}

func (r eventRepository) GetOne(uid string, id int) (model.EventGet, error) {
	query := event.EventGetOne
	event := model.EventGet{}
	db := r.client.GetDB()
	if err := db.Get(&event, query, uid, id); err != nil {
		return event, err
	}
	return event, nil
}
