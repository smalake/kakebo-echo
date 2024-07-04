package auth

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/errors"
)

// ログイン処理（FirebaseのUIDがusersテーブルに登録されているかチェック）
func (s *authService) Login(uid string) error {
	count, err := s.repo.FindUser(uid)
	if err != nil || count == -1 {
		return errors.InternalServerError
	}

	if count == 0 {
		return errors.ErrUserNotFound
	}
	return nil
}

// ユーザ登録
func (s *authService) Register(req *model.RegisterRequest) error {
	return s.repo.Register(req)
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
