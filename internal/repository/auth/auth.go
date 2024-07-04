package auth

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/errors"
	"log"
	"strings"
	"time"
)

func (r *authRepository) FindUser(uid string) (int, error) {
	// UIDがDBに登録されているかチェック
	query := postgresql.CheckUserByUid
	var uidCount int
	err := r.appModel.PsgrCli.DB.Get(&uidCount, query, uid)
	if err != nil {
		return -1, err
	}
	return uidCount, nil
}

func (r *authRepository) Register(register *model.RegisterRequest) error {
	// トランザクション開始
	tx, err := r.appModel.PsgrCli.DB.Beginx()
	if err != nil {
		log.Println("[FATAL] failed to transaction")
		return err
	}

	// グループ作成
	var groupQuery = postgresql.CreateGroup
	var groupId int
	err = tx.QueryRowx(groupQuery, time.Now(), time.Now()).Scan(&groupId)
	if err != nil {
		log.Println("[FATAL] failed to CREATE group")
		_ = tx.Rollback()
		return err
	}

	//ユーザ作成
	query := postgresql.RegisterUser
	_, err = tx.Exec(query, register.Uid, register.Name, groupId, register.Type, time.Now(), time.Now())
	if err != nil {
		_ = tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Println("[ERROR] ユーザがすでに登録されています")
			return errors.ErrUserAlreadyExist
		} else {
			log.Println("[FATAL] ユーザ登録に失敗しました")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("[FATAL] failed to commit")
		_ = tx.Rollback()
		return err
	}
	return nil
}

func (r *authRepository) LoginCheck(uid string) (int, error) {
	// UIDがDBに登録されているかチェック
	query := postgresql.CheckUserByUid
	var loginCheck model.LoginCheck
	err := r.appModel.PsgrCli.DB.Get(&loginCheck, query, uid)
	if err != nil {
		return -1, err
	}
	if !loginCheck.GroupAdmin.Valid {
		return -1, errors.ErrUserNotFound
	}
	return int(loginCheck.GroupAdmin.Int32), nil
}

func (r *authRepository) Logout() error {
	return nil
}
