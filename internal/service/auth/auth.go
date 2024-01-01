package auth

import (
	"kakebo-echo/internal/appmodels"
	"kakebo-echo/pkg/mysql"
	"kakebo-echo/pkg/structs"
	"kakebo-echo/pkg/user"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	appModel appmodels.AppModel
}

func New(appModel *appmodels.AppModel) *Service {
	return &Service{appModel: *appModel}
}

// メールアドレスでログイン
func (s *Service) LoginMail(ctx echo.Context) structs.HttpResponse {
	// POSTからログイン情報を取得
	u := new(user.LoginMailRequest)
	if err := ctx.Bind(u); err != nil {
		return structs.HttpResponse{Code: 400, Error: err}
	}

	query := mysql.LoginMail
	var uid user.LoginMailInfo
	err := s.appModel.MysqlCli.DB.Get(&uid, query, u.Email)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	if err := compareHashAndPassword(uid.Password, u.Password); err != nil {
		// 認証エラー
		return structs.HttpResponse{Code: 401, Error: err}
	}

	// トークンを発行
	token, err := issueToken(uid.ID)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	return structs.HttpResponse{Code: 200, Data: map[string]string{"accessToken": token}}
}

// Googleアカウントでログイン
func (s *Service) LoginGoogle(ctx echo.Context) structs.HttpResponse {
	// POSTからログイン情報を取得
	u := new(user.LoginGoogleRequest)
	if err := ctx.Bind(u); err != nil {
		return structs.HttpResponse{Code: 400, Error: err}
	}

	query := mysql.LoginGoogle
	var uid user.UserID
	err := s.appModel.MysqlCli.DB.Get(&uid, query, u.Email)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	// トークンを発行
	token, err := issueToken(uid.ID)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	return structs.HttpResponse{Code: 200, Data: map[string]string{"accessToken": token}}
}

// ユーザ登録
func (s *Service) RegisterUser(ctx echo.Context) structs.HttpResponse {
	// POSTからユーザ情報を取得
	u := new(user.RegisterUserRequest)
	if err := ctx.Bind(u); err != nil {
		return structs.HttpResponse{Code: 400, Error: err}
	}

	// パスワードをハッシュ化
	password, err := passwordEncrypt(u.Password)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}

	query := mysql.RegisterUser
	_, err = s.appModel.MysqlCli.DB.Exec(query, u.Email, password, u.Name)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}
	return structs.HttpResponse{Code: 200}
}

// ログアウト
func (s *Service) Logout(ctx echo.Context) structs.HttpResponse {
	return structs.HttpResponse{Code: 200}
}

// トークン発行
func issueToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = time.Now().Unix()                     // 発行時間
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 有効期限
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// パスワードのハッシュ化
func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// ハッシュ化されたパスワードとの比較(returnがnilならログイン成功)
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
