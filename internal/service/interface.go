package service

import "github.com/labstack/echo/v4"

type ServerInterface interface {

	// (POST /auth-code)
	AuthCodeHandler(ctx echo.Context) error

	// (GET /display-name)
	GetDisplayNameHandler(ctx echo.Context) error

	// (POST /display-name)
	UpdateDisplayNameHandler(ctx echo.Context) error

	// (GET /event)
	GetAllEventHandler(ctx echo.Context) error

	// (POST /event)
	CreateEventHandler(ctx echo.Context) error

	// (DELETE /event/${id})
	DeleteEventHandler(ctx echo.Context, id int) error

	// (GET /event/${id})
	GetOneEventHandler(ctx echo.Context, id int) error

	// (PUT /event/${id})
	UpdateEventHandler(ctx echo.Context, id int) error
	// 参加するグループの親のユーザ名を取得
	// (GET /get_name/${group})
	GetParentNameHandler(ctx echo.Context, group string) error

	// (GET /health-check)
	HealthCheckHandler(ctx echo.Context) error

	// (GET /invite)
	InviteHandler(ctx echo.Context) error

	// (GET /is-parent)
	IsParentHandler(ctx echo.Context) error
	// 共有グループに参加
	// (POST /join)
	JoinHandler(ctx echo.Context) error
	// ログインしているかチェック
	// (GET /login-check)
	LoginCheckHandler(ctx echo.Context) error
	// Googleアカウントでログイン
	// (POST /login-google)
	LoginGoogleHandler(ctx echo.Context) error
	// メールアドレスでログイン
	// (POST /login-mail)
	LoginMailHandler(ctx echo.Context) error
	// ログアウト処理
	// (GET /logout)
	LogoutHandler(ctx echo.Context) error

	// (GET /pattern)
	GetPatternHandler(ctx echo.Context) error

	// (POST /pattern)
	RegisterPatternHandler(ctx echo.Context) error

	// (DELETE /pattern/${id})
	DeletePatternHandler(ctx echo.Context, id int) error

	// (PUT /pattern/${id})
	UpdatePatternHandler(ctx echo.Context, id int) error

	// (GET /private)
	GetAllPrivateHandler(ctx echo.Context) error

	// (DELETE /private/${id})
	DeletePrivateHandler(ctx echo.Context, id int) error

	// (GET /private/${id})
	GetOnePrivateHandler(ctx echo.Context, id int) error

	// (PUT /private/${id})
	UpdatePrivateHandler(ctx echo.Context, id int) error
	// ユーザの新規登録
	// (POST /register)
	RegisterUserHandler(ctx echo.Context) error

	// (GET /resend-code)
	ResendCodeHandler(ctx echo.Context) error

	// (GET /revision)
	RevisionHandler(ctx echo.Context) error

	// (POST /send-mail)
	SendMailHandler(ctx echo.Context) error
}
