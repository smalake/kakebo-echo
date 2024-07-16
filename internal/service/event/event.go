package event

import (
	"context"
	"kakebo-echo/internal/model"
	"kakebo-echo/internal/repository/event"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/pkg/errors"
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

func (s eventService) Create(e model.EventCreate, uid string) ([]int, error) {
	// 日付をstring型からdate型へと変換
	formattedDate, err := time.Parse("2006-01-02T15:04:05.000Z", e.Date)
	if err != nil {
		return nil, err
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
	checkEvent1, source := eventValidation(event1)
	if !checkEvent1 {
		log.Printf("[ERROR] event1 %s is bad value", source)
		return nil, errors.BadRequest
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
		checkEvent2, source := eventValidation(event2)
		if !checkEvent2 {
			log.Printf("[ERROR] event2 %s is bad value", source)
			return nil, errors.BadRequest
		}
	}

	// uidからgroup_idを取得
	usid, gid, err := s.repo.GetIDs(uid)
	if err != nil {
		return nil, err
	}

	// 追加したカラムのIDを結果として返す
	var result []int

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
		result = append(result, id1)
		if e.Amount2 > 0 {
			revision2, err := s.repo.UpdateRevision(tx, gid)
			if err != nil {
				return err
			}
			id2, err := s.repo.Create(tx, event2, revision2, usid, gid)
			if err != nil {
				return err
			}
			result = append(result, id2)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func (s eventService) GetAll(uid string) ([]model.EventGet, error) {
	return s.repo.GetAll(uid)
}

func (s eventService) GetOne(uid string, id int) (model.EventGet, error) {
	return s.repo.GetOne(uid, id)
}

func (s eventService) GetRevision(uid string) (int, error) {
	// uidからgroup_idを取得
	_, gid, err := s.repo.GetIDs(uid)
	if err != nil {
		return -1, err
	}
	return s.repo.GetRevision(gid)
}

// イベントの内容についてバリデーション
func eventValidation(event model.Event) (bool, string) {
	if event.Amount <= 0 {
		return false, "amount"
	}
	if event.Category < 0 || event.Category > 9 {
		return false, "category"
	}
	return true, ""
}
