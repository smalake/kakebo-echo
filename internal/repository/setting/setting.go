package setting

import (
	"kakebo-echo/pkg/database/postgresql"
	"kakebo-echo/pkg/database/postgresql/setting"
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
