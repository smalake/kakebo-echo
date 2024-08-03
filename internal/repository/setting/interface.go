package setting

type SettingRepository interface {
	GetAdminByUID(string) (int, error)
	GetName(string) (string, error)
	UpdateName(string, string) error
}
