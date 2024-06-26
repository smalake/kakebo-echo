package auth

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/internal/repository"
)

type AuthService interface {
	Login(string) (int, error)
	Register(model.RegisterRequest) (int, error)
	LoginCheck(string) (int, error)
	Logout() error
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

// ログイン処理（FirebaseのUIDがusersテーブルに登録されているかチェック）
func (s *authService) Login(uid string) (int, error) {
	// POSTからログイン情報を取得
	// TODO：メモ用（転機したら削除）
	// u := new(user.LoginRequest)
	// if err := ctx.Bind(u); err != nil {
	// 	return structs.HttpResponse{Code: 400, Error: err}
	// }
	return s.repo.Login(uid)
}

// ユーザ登録
func (s *authService) Register(req model.RegisterRequest) (int, error) {
	// POSTからユーザ情報を取得
	// u := new(user.RegisterUserRequest)
	// if err := ctx.Bind(u); err != nil {
	// 	return structs.HttpResponse{Code: 400, Error: err}
	// }

	return s.repo.Register(req)
}

// ログアウト
func (s *authService) Logout() error {
	return s.repo.Logout()
}

// ログイン確認
func (s *authService) LoginCheck(uid string) (int, error) {
	return s.repo.LoginCheck(uid)
}
