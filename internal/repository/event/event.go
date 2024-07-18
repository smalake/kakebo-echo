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

func (r eventRepository) GetIDs(uid string) (int, int, error) {
	query := event.GetID
	ids := model.GetIDs{}
	db := r.client.GetDB()
	if err := db.Get(&ids, query, uid); err != nil {
		return 0, 0, err
	}
	return ids.ID, ids.GroupID, nil
}

func (r eventRepository) Create(tx *sqlx.Tx, e model.Event, revision, userId, groupId int) (int, error) {
	// イベントの追加
	query := event.EventCreate
	var id int
	err := tx.QueryRow(query, e.Amount, e.Category, e.StoreName, e.Memo, e.Date, groupId, revision, userId, userId, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		log.Println("[FATAL] イベントの新規作成に失敗しました")
		return -1, err
	}
	// 登録したeventのidを返す
	return id, nil
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

func (r eventRepository) UpdateRevision(tx *sqlx.Tx, gid int) (int, error) {
	query := event.UpdateRevision
	var revision int
	if err := tx.Get(&revision, query, gid); err != nil {
		return -1, err
	}
	return revision, nil
}
