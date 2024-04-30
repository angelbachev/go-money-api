package update_user_settings

import "github.com/angelbachev/go-money-api/domain/user"

type UpdateUserSettingsResponse struct {
	UserSettings *user.UserSettings
}

func NewUpdateUserSettingsResponse(userSettings *user.UserSettings) *UpdateUserSettingsResponse {
	return &UpdateUserSettingsResponse{
		UserSettings: userSettings,
	}
}
