package service

import "github.com/labstack/echo/v4"

type ServerInterface interface {

	// (POST /auth-code)
	AuthCode(ctx echo.Context) error

	// (GET /display-name)
	GetDisplayName(ctx echo.Context) error

	// (POST /display-name)
	UpdateDisplayName(ctx echo.Context) error

	// (GET /event)
	GetAllEvent(ctx echo.Context) error

	// (POST /event)
	CreateEvent(ctx echo.Context) error

	// (DELETE /event/${id})
	DeleteEvent(ctx echo.Context, id int) error

	// (GET /event/${id})
	GetOneEvent(ctx echo.Context, id int) error

	// (PUT /event/${id})
	UpdateEvent(ctx echo.Context, id int) error
	// 参加するグループの親のユーザ名を取得
	// (GET /get_name/${group})
	GetParentName(ctx echo.Context, group string) error

	// (GET /health-check)
	HealthCheck(ctx echo.Context) error

	// (GET /invite)
	Invite(ctx echo.Context) error

	// (GET /is-parent)
	IsParent(ctx echo.Context) error
	// 共有グループに参加
	// (POST /join)
	Join(ctx echo.Context) error
	// ログインしているかチェック
	// (GET /login-check)
	LoginCheck(ctx echo.Context) error
	// Googleアカウントでログイン
	// (POST /login-google)
	LoginGoogle(ctx echo.Context) error
	// メールアドレスでログイン
	// (POST /login-mail)
	LoginMail(ctx echo.Context) error
	// ログアウト処理
	// (GET /logout)
	Logout(ctx echo.Context) error

	// (GET /pattern)
	GetPattern(ctx echo.Context) error

	// (POST /pattern)
	RegisterPattern(ctx echo.Context) error

	// (DELETE /pattern/${id})
	DeletePattern(ctx echo.Context, id int) error

	// (PUT /pattern/${id})
	UpdatePattern(ctx echo.Context, id int) error

	// (GET /private)
	GetAllPrivate(ctx echo.Context) error

	// (DELETE /private/${id})
	DeletePrivate(ctx echo.Context, id int) error

	// (GET /private/${id})
	GetOnePrivate(ctx echo.Context, id int) error

	// (PUT /private/${id})
	UpdatePrivate(ctx echo.Context, id int) error
	// ユーザの新規登録
	// (POST /register)
	RegisterUser(ctx echo.Context) error

	// (GET /resend-code)
	ResendCode(ctx echo.Context) error

	// (GET /revision)
	Revision(ctx echo.Context) error

	// (POST /send-mail)
	SendMail(ctx echo.Context) error
}
