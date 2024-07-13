package event

import (
	"kakebo-echo/internal/appmodel"
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql/event"
	"log"
	"time"
)

type eventRepository struct {
	appModel appmodel.AppModel
}

func New(am appmodel.AppModel) EventRepository {
	return &eventRepository{appModel: am}
}

func (r eventRepository) Create(e model.Event, uid string) error {
	// トランザクション開始
	tx, err := r.appModel.PsgrCli.DB.Beginx()
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
	if err := r.appModel.PsgrCli.DB.Get(&events, query, uid); err != nil {
		return nil, err
	}
	return events, nil
}

func (r eventRepository) GetOne(id int) (model.Event, error) {
	return model.Event{}, nil
}
