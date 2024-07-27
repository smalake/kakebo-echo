package auth

import (
	"kakebo-echo/internal/model"
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/database/postgresql/auth"
	"kakebo-echo/pkg/errors"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
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
	var isAdmin int
	db := r.client.GetDB()
	err := db.Get(&isAdmin, query, uid)
	if err != nil {
		return -1, err
	}
	return isAdmin, nil
}

func (r *authRepository) CreateGroup(tx *sqlx.Tx) (int, error) {
	query := auth.CreateGroup
	var groupId int
	if err := tx.QueryRowx(query, time.Now(), time.Now()).Scan(&groupId); err != nil {
		log.Println("[FATAL] failed to CREATE group")
		return -1, err
	}
	return groupId, nil
}

func (r *authRepository) CreateUser(tx *sqlx.Tx, register *model.RegisterRequest, groupId int) (int, error) {
	query := auth.CreateUser
	var userId int
	if err := tx.QueryRowx(query, register.Uid, register.Name, groupId, register.Type, time.Now(), time.Now()).Scan(&userId); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Println("[ERROR] ユーザがすでに登録されています")
			return -1, errors.ErrUserAlreadyExist
		} else {
			return -1, err
		}
	}
	return userId, nil
}

func (r *authRepository) CreateRevision(tx *sqlx.Tx, userId int) error {
	query := auth.CreateRevision
	if _, err := tx.Exec(query, userId, time.Now(), time.Now()); err != nil {
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
