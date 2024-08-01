package event

import (
	"context"
	"kakebo-echo/internal/model"
	"kakebo-echo/internal/repository/event"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/pkg/errors"
	"kakebo-echo/pkg/util"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type eventService struct {
	repo        event.EventRepository
	transaction transaction.TransactionRepository
}

func New(repo event.EventRepository, transRepo transaction.TransactionRepository) EventService {
	return &eventService{repo: repo, transaction: transRepo}
}

func (s eventService) Create(e model.EventCreate, uid string) ([]int, int, error) {
	// 日付をstring型からdate型へと変換
	formattedDate, err := time.Parse("2006-01-02T15:04:05.000Z", e.Date)
	if err != nil {
		return nil, -1, err
	}

	// event1は必ず存在するためそのまま処理
	event1 := model.Event{
		Amount:    e.Amount1,
		Category:  e.Category1,
		Memo:      e.Memo1,
		StoreName: e.StoreName,
		Date:      formattedDate,
	}
	// トランザクション内で使用するため、明示的にevent2を初期化。eventが1つだけの場合使用されない
	event2 := model.Event{}

	// バリデーション
	checkEvent1, source := util.EventValidation(event1)
	if !checkEvent1 {
		log.Printf("[ERROR] event1 %s is bad value", source)
		return nil, -1, errors.BadRequest
	}

	// イベントが2件の場合は分割してevent2として処理する
	if e.Amount2 > 0 {
		event2 = model.Event{
			Amount:    e.Amount2,
			Category:  e.Category2,
			Memo:      e.Memo2,
			StoreName: e.StoreName,
			Date:      formattedDate,
		}
		// バリデーション
		checkEvent2, source := util.EventValidation(event2)
		if !checkEvent2 {
			log.Printf("[ERROR] event2 %s is bad value", source)
			return nil, -1, errors.BadRequest
		}
	}

	// uidからgroup_idを取得
	usid, gid, err := s.repo.GetIDs(uid)
	if err != nil {
		return nil, -1, err
	}

	// 追加したカラムのIDを結果として返す
	var ids []int
	var revision int

	// トランザクション処理
	if err := s.transaction.Transaction(context.TODO(), func(tx *sqlx.Tx) error {
		revision1, err := s.repo.UpdateRevision(tx, gid)
		if err != nil {
			return err
		}
		// eventのバリデーションはトランザクション開始前に完了させておく
		id1, err := s.repo.Create(tx, event1, revision1, usid, gid)
		if err != nil {
			return err
		}
		ids = append(ids, id1)
		revision = revision1
		if e.Amount2 > 0 {
			revision2, err := s.repo.UpdateRevision(tx, gid)
			if err != nil {
				return err
			}
			id2, err := s.repo.Create(tx, event2, revision2, usid, gid)
			if err != nil {
				return err
			}
			ids = append(ids, id2)
			revision = revision2
		}
		return nil
	}); err != nil {
		return nil, -1, err
	}
	return ids, revision, nil
}

func (s eventService) GetAll(uid string) ([]model.EventResponse, error) {
	events, err := s.repo.GetAll(uid)
	if err != nil {
		return nil, err
	}

	// Dateを文字列かつ"%Y-%m-%d"形式にフォーマットするための処理
	response := make([]model.EventResponse, len(events))
	for i, event := range events {
		response[i].ID = event.ID
		response[i].Amount = event.Amount
		response[i].Category = event.Category
		response[i].StoreName = event.StoreName
		response[i].Date = event.Date.Format("2006-01-02")
	}
	return response, nil
}

func (s eventService) GetOne(uid string, id int) (model.EventOne, error) {
	event, err := s.repo.GetOne(uid, id)
	if err != nil {
		return model.EventOne{}, err
	}
	event.CreatedAt = event.CreatedAtDate.Format("2006-01-02 15:04:05")
	event.UpdatedAt = event.UpdatedAtDate.Format("2006-01-02 15:04:05")
	return event, nil
}

func (s eventService) Update(e model.EventUpdate, uid string, id int) (int, error) {
	// uidからgroup_idを取得
	usid, gid, err := s.repo.GetIDs(uid)
	if err != nil {
		return -1, err
	}
	var revision int
	// トランザクション処理
	if err := s.transaction.Transaction(context.TODO(), func(tx *sqlx.Tx) error {
		revision, err = s.repo.UpdateRevision(tx, gid)
		if err != nil {
			return err
		}
		err = s.repo.Update(tx, e, uid, id, usid, revision)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return revision, nil
}

func (s eventService) Delete(uid string, id int) (int, error) {
	// uidからgroup_idを取得
	_, gid, err := s.repo.GetIDs(uid)
	if err != nil {
		return -1, err
	}
	var revision int
	// トランザクション処理
	if err := s.transaction.Transaction(context.TODO(), func(tx *sqlx.Tx) error {
		revision, err = s.repo.UpdateRevision(tx, gid)
		if err != nil {
			return err
		}
		err = s.repo.Delete(tx, gid, id)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return revision, nil
}

func (s eventService) GetRevision(uid string) (int, error) {
	// uidからgroup_idを取得
	_, gid, err := s.repo.GetIDs(uid)
	if err != nil {
		return -1, err
	}
	return s.repo.GetRevision(gid)
}
