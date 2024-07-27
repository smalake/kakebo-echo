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
