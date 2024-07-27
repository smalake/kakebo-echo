package setting

type SettingRepository interface {
	GetAdminByUID(string) (int, error)
}
