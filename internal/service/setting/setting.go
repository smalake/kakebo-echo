package setting

import (
	"kakebo-echo/internal/repository/setting"
	"kakebo-echo/internal/repository/transaction"
)

type settingService struct {
	repo        setting.SettingRepository
	transaction transaction.TransactionRepository
}

func New(repo setting.SettingRepository, transRepo transaction.TransactionRepository) SettingService {
	return &settingService{repo: repo, transaction: transRepo}
}

func (s settingService) AdminCheck(uid string) (int, error) {
	return s.repo.GetAdminByUID(uid)
}

func (s settingService) GetName(uid string) (string, error) {
	return s.repo.GetName(uid)
}

func (s settingService) UpdateName(uid, name string) error {
	return s.repo.UpdateName(uid, name)
}
