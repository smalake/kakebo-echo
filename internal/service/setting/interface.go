package setting

type SettingService interface {
	AdminCheck(string) (int, error)
	GetName(string) (string, error)
	UpdateName(string, string) error
}
