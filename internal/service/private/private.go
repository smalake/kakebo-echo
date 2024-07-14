package private

import (
	"context"
	"kakebo-echo/internal/model"
	"kakebo-echo/internal/repository/private"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/pkg/errors"
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

func (s privateService) Create(p model.PrivateCreate, uid string) error {
	// 日付をstring型からdate型へと変換
	formattedDate, err := time.Parse("2006-01-02T15:04:05.000Z", p.Date)
	if err != nil {
		return err
	}
	// private1は必ず存在するためそのまま処理
	private1 := model.Private{
		Amount:    p.Amount1,
		Category:  p.Category1,
		Memo:      p.Memo1,
		StoreName: p.StoreName,
		Date:      formattedDate,
	}
	// トランザクション内で使用するため、明示的にprivate2を初期化。privateが1つだけの場合使用されない
	private2 := model.Private{}

	// バリデーション
	checkPrivate1, source := privateValidation(private1)
	if !checkPrivate1 {
		log.Printf("[ERROR] private1 %s is bad value", source)
		return errors.BadRequest
	}

	// イベントが2件の場合は分割してprivate2として処理する
	if p.Amount2 > 0 {
		private2 = model.Private{
			Amount:    p.Amount2,
			Category:  p.Category2,
			Memo:      p.Memo2,
			StoreName: p.StoreName,
			Date:      formattedDate,
		}
		// バリデーション
		checkPrivate2, source := privateValidation(private2)
		if !checkPrivate2 {
			log.Printf("[ERROR] private2 %s is bad value", source)
			return errors.BadRequest
		}
	}

	// トランザクション処理
	if err := s.transaction.Transaction(context.TODO(), func(tx *sqlx.Tx) error {
		// privateのバリデーションはトランザクション開始前に完了させておく
		if err := s.repo.Create(private1, uid); err != nil {
			return err
		}
		if p.Amount2 > 0 {
			if err := s.repo.Create(private2, uid); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s privateService) GetAll(uid string) ([]model.PrivateGet, error) {
	return s.repo.GetAll(uid)
}

func (s privateService) GetOne(uid string, id int) (model.PrivateGet, error) {
	return s.repo.GetOne(uid, id)
}

// イベントの内容についてバリデーション
func privateValidation(private model.Private) (bool, string) {
	if private.Amount <= 0 {
		return false, "amount"
	}
	if private.Category < 0 || private.Category > 9 {
		return false, "category"
	}
	return true, ""
}
