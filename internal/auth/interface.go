package auth

import "github.com/labstack/echo"

type JWT struct {
	Token string `json:"token"`
}

type AuthMail struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type AuthGoogle struct {
	Mail string `json:"mail"`
}

type Register struct {
	Mail     string `json:"mail"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}

type ParentName struct {
	Name string `json:"name"`
}

type AuthCode struct {
	Code  int    `json:"code"`
	Email string `json:"email"`
}

type AuthInterface interface {
	// メールアドレスでログイン
	// (POST /login-mail)
	LoginMail(ctx echo.Context) error

	// Googleアカウントでログイン
	// (POST /login-google)
	LoginGoogle(ctx echo.Context) error

	// ログアウト処理
	// (GET /logout)
	Logout(ctx echo.Context) error

	// ユーザの新規登録
	// (POST /register)
	Register(ctx echo.Context) error

	// 共有グループに参加
	// (POST /join)
	Join(ctx echo.Context) error

	// 参加するグループの親のユーザ名を取得
	// (GET /get_name/${group})
	GetParentName(ctx echo.Context, group string) error

	// ログインしているかチェック
	// (GET /login-check)
	LoginCheck(ctx echo.Context) error

	// (POST /auth-code)
	AuthCode(ctx echo.Context) error

	// (GET /resend-code)
	ResendCode(ctx echo.Context) error
}
