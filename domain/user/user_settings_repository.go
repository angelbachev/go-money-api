package user

type UserSettingsRepository interface {
	CreateUserSettings(settings *UserSettings) error
	UpdateUserSettings(settings *UserSettings) error
	GetUserSettings(userID int64) (*UserSettings, error)
}
