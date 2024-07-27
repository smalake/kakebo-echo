package setting

type SettingService interface {
	AdminCheck(string) (int, error)
}
