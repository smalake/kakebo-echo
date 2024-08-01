package private

import (
	"context"
	"kakebo-echo/internal/model"
	"kakebo-echo/internal/repository/private"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/pkg/errors"
	"kakebo-echo/pkg/util"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type privateService struct {
	repo        private.PrivateRepository
	transaction transaction.TransactionRepository
}

func New(repo private.PrivateRepository, transRepo transaction.TransactionRepository) PrivateService {
	return &privateService{repo: repo, transaction: transRepo}
}

func (s privateService) Create(p model.EventCreate, uid string) ([]int, int, error) {
	// 日付をstring型からdate型へと変換
	formattedDate, err := time.Parse("2006-01-02T15:04:05.000Z", p.Date)
	if err != nil {
		return nil, -1, err
	}
	// private1は必ず存在するためそのまま処理
	private1 := model.Event{
		Amount:    p.Amount1,
		Category:  p.Category1,
		Memo:      p.Memo1,
		StoreName: p.StoreName,
		Date:      formattedDate,
	}

	// バリデーション
	checkPrivate1, source := util.EventValidation(private1)
	if !checkPrivate1 {
		log.Printf("[ERROR] private1 %s is bad value", source)
		return nil, -1, errors.BadRequest
	}

	// トランザクション内で使用するため、明示的にprivate2を初期化。privateが1つだけの場合使用されない
	private2 := model.Event{}

	// イベントが2件の場合は分割してprivate2として処理する
	if p.Amount2 > 0 {
		private2 = model.Event{
			Amount:    p.Amount2,
			Category:  p.Category2,
			Memo:      p.Memo2,
			StoreName: p.StoreName,
			Date:      formattedDate,
		}
		// バリデーション
		checkPrivate2, source := util.EventValidation(private2)
		if !checkPrivate2 {
			log.Printf("[ERROR] private2 %s is bad value", source)
			return nil, -1, errors.BadRequest
		}
	}

	// 追加したカラムのIDを結果として返す
	var ids []int
	var revision int

	// トランザクション処理
	if err := s.transaction.Transaction(context.TODO(), func(tx *sqlx.Tx) error {
		revision1, err := s.repo.UpdateRevision(tx, uid)
		if err != nil {
			return err
		}
		log.Printf("[revision1] ")
		// privateのバリデーションはトランザクション開始前に完了させておく
		id1, err := s.repo.Create(tx, private1, revision1, uid)
		if err != nil {
			return err
		}
		ids = append(ids, id1)
		revision = revision1

		if p.Amount2 > 0 {
			revision2, err := s.repo.UpdateRevision(tx, uid)
			if err != nil {
				return err
			}
			id2, err := s.repo.Create(tx, private2, revision2, uid)
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

func (s privateService) GetAll(uid string) ([]model.EventResponse, error) {
	privates, err := s.repo.GetAll(uid)
	if err != nil {
		return nil, err
	}

	// Dateを文字列かつ"%Y-%m-%d"形式にフォーマットするための処理
	response := make([]model.EventResponse, len(privates))
	for i, private := range privates {
		response[i].ID = private.ID
		response[i].Amount = private.Amount
		response[i].Category = private.Category
		response[i].StoreName = private.StoreName
		response[i].Date = private.Date.Format("2006-01-02")
	}
	return response, nil
}

func (s privateService) GetOne(uid string, id int) (model.PrivateOne, error) {
	private, err := s.repo.GetOne(uid, id)
	if err != nil {
		return model.PrivateOne{}, err
	}
	private.CreatedAt = private.CreatedAtDate.Format("2006-01-02 15:04:05")
	private.UpdatedAt = private.UpdatedAtDate.Format("2006-01-02 15:04:05")
	return private, nil
}

func (s privateService) Update(e model.EventUpdate, uid string, id int) (int, error) {
	var revision int
	var err error
	// トランザクション処理
	if err := s.transaction.Transaction(context.TODO(), func(tx *sqlx.Tx) error {
		revision, err = s.repo.UpdateRevision(tx, uid)
		if err != nil {
			return err
		}
		err = s.repo.Update(tx, e, uid, id, revision)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return revision, nil
}

func (s privateService) Delete(uid string, id int) (int, error) {
	var revision int
	var err error
	// トランザクション処理
	if err := s.transaction.Transaction(context.TODO(), func(tx *sqlx.Tx) error {
		revision, err = s.repo.UpdateRevision(tx, uid)
		if err != nil {
			return err
		}
		err = s.repo.Delete(tx, uid, id)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return -1, err
	}
	return revision, nil
}

func (s privateService) GetRevision(uid string) (int, error) {
	return s.repo.GetRevision(uid)
}
