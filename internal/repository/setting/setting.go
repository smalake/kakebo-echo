package setting

import (
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/database/postgresql/setting"
	"time"
)

type settingRepository struct {
	client postgresql.ClientInterface
}

func New(cl postgresql.ClientInterface) SettingRepository {
	return &settingRepository{client: cl}
}

func (r settingRepository) GetAdminByUID(uid string) (int, error) {
	query := setting.GetAdminByUID
	var adminCheck int
	db := r.client.GetDB()
	if err := db.Get(&adminCheck, query, uid); err != nil {
		return -1, err
	}
	return adminCheck, nil
}

func (r settingRepository) GetName(uid string) (string, error) {
	query := setting.GetName
	var name string
	db := r.client.GetDB()
	if err := db.Get(&name, query, uid); err != nil {
		return "", err
	}
	return name, nil
}

func (r settingRepository) UpdateName(uid, name string) error {
	query := setting.UpdateName
	db := r.client.GetDB()
	if _, err := db.Exec(query, name, time.Now(), uid); err != nil {
		return err
	}
	return nil
}
