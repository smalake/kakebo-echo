package event

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/database/postgresql/event"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type eventRepository struct {
	client postgresql.ClientInterface
}

func New(cl postgresql.ClientInterface) EventRepository {
	return &eventRepository{client: cl}
}

func (r eventRepository) GetGroupID(uid string) (int, error) {
	query := event.GetGroupID
	var gid int
	db := r.client.GetDB()
	if err := db.Get(&gid, query, uid); err != nil {
		return 0, err
	}
	return gid, nil
}

func (r eventRepository) Create(tx *sqlx.Tx, e model.Event, groupId int) error {
	// イベントの追加
	query := event.EventCreate
	_, err := tx.Exec(query, e.Amount, e.Category, e.StoreName, e.Memo, e.Date, groupId, time.Now(), time.Now())
	if err != nil {
		log.Println("[FATAL] イベントの新規作成に失敗しました")
		return err
	}
	return nil
}

func (r eventRepository) GetAll(uid string) ([]model.EventGet, error) {
	query := event.EventGetAll
	events := []model.EventGet{}
	db := r.client.GetDB()
	if err := db.Select(&events, query, uid); err != nil {
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

func (r eventRepository) GetRevision(gid int) (int, error) {
	query := event.GetRevision
	var revision int
	db := r.client.GetDB()
	if err := db.Get(&revision, query, gid); err != nil {
		return -1, err
	}
	return revision, nil
}
