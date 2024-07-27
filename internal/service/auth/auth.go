package auth

import (
	"context"
	"kakebo-echo/internal/model"
	"kakebo-echo/internal/repository/auth"
	"kakebo-echo/internal/repository/transaction"
	"kakebo-echo/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type authService struct {
	repo        auth.AuthRepository
	transaction transaction.TransactionRepository
}

func New(repo auth.AuthRepository, transRepo transaction.TransactionRepository) AuthService {
	return &authService{repo: repo, transaction: transRepo}
}

// ログイン処理（FirebaseのUIDがusersテーブルに登録されているかチェック）
func (s *authService) Login(uid string) error {
	count, err := s.repo.FindUser(uid)
	if err != nil || count == -1 {
		return errors.InternalServerError
	}
	return nil
}

// ユーザ登録
func (s *authService) Register(req *model.RegisterRequest) error {
	// トランザクション処理
	if err := s.transaction.Transaction(context.TODO(), func(tx *sqlx.Tx) error {
		groupId, err := s.repo.CreateGroup(tx)
		if err != nil {
			return err
		}
		userId, err := s.repo.CreateUser(tx, req, groupId)
		if err != nil {
			return err
		}
		err = s.repo.CreateRevision(tx, userId)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// ログアウト
func (s *authService) Logout() error {
	return s.repo.Logout()
}

// ログイン確認
func (s *authService) LoginCheck(uid string) (int, error) {
	admin, err := s.repo.LoginCheck(uid)
	if err != nil {
		return admin, err
	}
	return admin, nil
}
