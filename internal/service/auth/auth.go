package auth

import (
	"errors"
	"kakebo-echo/internal/appmodels"
	"kakebo-echo/pkg/postgresql"
	"kakebo-echo/pkg/structs"
	"kakebo-echo/pkg/user"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Service struct {
	appModel appmodels.AppModel
}

func New(appModel *appmodels.AppModel) *Service {
	return &Service{appModel: *appModel}
}

// ログイン処理（FirebaseのUIDがusersテーブルに登録されているかチェック）
func (s *Service) Login(ctx echo.Context) structs.HttpResponse {
	// POSTからログイン情報を取得
	u := new(user.LoginRequest)
	if err := ctx.Bind(u); err != nil {
		return structs.HttpResponse{Code: 400, Error: err}
	}

	// UIDがDBに登録されているかチェック
	query := postgresql.CheckUserByUid
	var uidCount int
	err := s.appModel.PsgrCli.DB.Get(&uidCount, query, u.Uid)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}
	if uidCount == 0 {
		return structs.HttpResponse{Code: 401, Message: "Unregistered User"}
	}
	return structs.HttpResponse{Code: 200}
}

// ユーザ登録
func (s *Service) RegisterUser(ctx echo.Context) structs.HttpResponse {
	// POSTからユーザ情報を取得
	u := new(user.RegisterUserRequest)
	if err := ctx.Bind(u); err != nil {
		return structs.HttpResponse{Code: 400, Error: err}
	}

	// トランザクション開始
	tx, err := s.appModel.PsgrCli.DB.Beginx()
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to transaction: %+v", err)
		return structs.HttpResponse{Code: 500, Error: err}
	}

	// グループ作成
	var groupQuery = postgresql.CreateGroup
	var groupId int
	err = tx.QueryRowx(groupQuery, time.Now(), time.Now()).Scan(&groupId)
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to CREATE group: %+v", err)
		_ = tx.Rollback()
		return structs.HttpResponse{Code: 500, Error: err}
	}

	//ユーザ作成
	query := postgresql.RegisterUser
	_, err = tx.Exec(query, u.Uid, u.Name, groupId, u.Type, time.Now(), time.Now())
	if err != nil {
		_ = tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			ctx.Logger().Errorf("[ERROR] ユーザがすでに登録されています: %+v", err)
			return structs.HttpResponse{Code: 409, Error: errors.New("ユーザがすでに登録されています")}
		} else {
			ctx.Logger().Errorf("[FATAL] ユーザ登録に失敗しました: %+v", err)
			return structs.HttpResponse{Code: 500, Error: err}
		}
	}

	err = tx.Commit()
	if err != nil {
		ctx.Logger().Errorf("[FATAL] failed to commit: %+v", err)
		_ = tx.Rollback()
		return structs.HttpResponse{Code: 500, Error: err}
	}
	return structs.HttpResponse{Code: 200}
}

// ログアウト
func (s *Service) Logout(ctx echo.Context) structs.HttpResponse {
	return structs.HttpResponse{Code: 200}
}

// ログイン確認
func (s *Service) LoginCheck(ctx echo.Context) structs.HttpResponse {
	uid := ctx.Get("uid")
	// JWTから取得したUIDがDBに登録されているかチェック
	query := postgresql.CheckUserByUid
	var loginCheck LoginCheck
	err := s.appModel.PsgrCli.DB.Get(&loginCheck, query, uid)
	if err != nil {
		return structs.HttpResponse{Code: 500, Error: err}
	}
	if !loginCheck.GroupAdmin.Valid {
		return structs.HttpResponse{Code: 401, Message: "Unregistered User"}
	}
	return structs.HttpResponse{Code: 200, Data: map[string]interface{}{"parent": loginCheck.GroupAdmin.Int32}}
}
