package get_user_settings

import (
	"github.com/angelbachev/go-money-api/domain/user"
)

type GetUserSettingsResponse struct {
	UserSettings *user.UserSettings
}

func NewGetUserSettingsResponse(userSettings *user.UserSettings) *GetUserSettingsResponse {
	return &GetUserSettingsResponse{
		UserSettings: userSettings,
	}
}
