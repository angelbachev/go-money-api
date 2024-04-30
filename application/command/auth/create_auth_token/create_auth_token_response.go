package create_auth_token

import "github.com/angelbachev/go-money-api/domain/user"

type CreateAuthTokenResponse struct {
	AccessToken string             `json:"accessToken"`
	Settings    *user.UserSettings `json:"settings"`
}

func NewCreateAuthTokenResponse(accessToken string, settings *user.UserSettings) *CreateAuthTokenResponse {
	return &CreateAuthTokenResponse{
		AccessToken: accessToken,
		Settings:    settings,
	}
}
