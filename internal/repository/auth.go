package repository

import (
	"kakebo-echo/internal/appmodel"
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql"
	"log"
	"net/http"
	"strings"
	"time"
)

type AuthRepository interface {
	Login(string) (int, error)
	Register(model.RegisterRequest) (int, error)
	LoginCheck(string) (int, error)
	Logout() error
}

type repository struct {
	appModel appmodel.AppModel
}

func NewAuthRepository(am *appmodel.AppModel) *repository {
	return &repository{appModel: *am}
}

func (r *repository) Login(uid string) (int, error) {
	// UIDがDBに登録されているかチェック
	query := postgresql.CheckUserByUid
	var uidCount int
	err := r.appModel.PsgrCli.DB.Get(&uidCount, query, uid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if uidCount == 0 {
		return http.StatusUnauthorized, err
	}
	return http.StatusOK, nil
}

func (r *repository) Register(register model.RegisterRequest) (int, error) {
	// トランザクション開始
	tx, err := r.appModel.PsgrCli.DB.Beginx()
	if err != nil {
		log.Printf("[FATAL] failed to transaction: %+v", err)
		return http.StatusInternalServerError, err
	}

	// グループ作成
	var groupQuery = postgresql.CreateGroup
	var groupId int
	err = tx.QueryRowx(groupQuery, time.Now(), time.Now()).Scan(&groupId)
	if err != nil {
		log.Printf("[FATAL] failed to CREATE group: %+v", err)
		_ = tx.Rollback()
		return http.StatusInternalServerError, err
	}

	//ユーザ作成
	query := postgresql.RegisterUser
	_, err = tx.Exec(query, register.Uid, register.Name, groupId, register.Type, time.Now(), time.Now())
	if err != nil {
		_ = tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Printf("[ERROR] ユーザがすでに登録されています: %+v", err)
			return http.StatusConflict, err
		} else {
			log.Printf("[FATAL] ユーザ登録に失敗しました: %+v", err)
			return http.StatusInternalServerError, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("[FATAL] failed to commit: %+v", err)
		_ = tx.Rollback()
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (r *repository) LoginCheck(uid string) (int, error) {
	// UIDがDBに登録されているかチェック
	query := postgresql.CheckUserByUid
	var loginCheck model.LoginCheck
	err := r.appModel.PsgrCli.DB.Get(&loginCheck, query, uid)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if !loginCheck.GroupAdmin.Valid {
		return http.StatusUnauthorized, err
	}
	return http.StatusOK, nil
}

func (r *repository) Logout() error {
	return nil
}
