package auth

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/database/postgresql/auth"
	"kakebo-echo/pkg/errors"
	"log"
	"strings"
	"time"
)

type authRepository struct {
	client postgresql.ClientInterface
}

func New(cl postgresql.ClientInterface) AuthRepository {
	return &authRepository{client: cl}
}

func (r *authRepository) FindUser(uid string) (int, error) {
	// UIDがDBに登録されているかチェック
	query := auth.CheckUserByUid
	var uidCount int
	db := r.client.GetDB()
	err := db.Get(&uidCount, query, uid)
	if err != nil {
		return -1, err
	}
	return uidCount, nil
}

func (r *authRepository) Register(register *model.RegisterRequest) error {
	// トランザクション開始
	db := r.client.GetDB()
	tx, err := db.Beginx()
	if err != nil {
		log.Println("[FATAL] failed to transaction")
		return err
	}

	// グループ作成
	var groupQuery = auth.CreateGroup
	var groupId int
	err = tx.QueryRowx(groupQuery, time.Now(), time.Now()).Scan(&groupId)
	if err != nil {
		log.Println("[FATAL] failed to CREATE group")
		_ = tx.Rollback()
		return err
	}

	//ユーザ作成
	query := auth.RegisterUser
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
	query := auth.CheckUserByUid
	loginCheck := model.LoginCheck{}
	db := r.client.GetDB()
	err := db.Get(&loginCheck, query, uid)
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
