package get_user_settings

import (
	"github.com/angelbachev/go-money-api/domain/user"
)

type GetUserSettingsQueryHandler struct {
	userSettingsRepository user.UserSettingsRepositoryInterface
}

func NewGetUserSettingsQueryHandler(
	userSettingsRepository user.UserSettingsRepositoryInterface,
) *GetUserSettingsQueryHandler {
	return &GetUserSettingsQueryHandler{
		userSettingsRepository: userSettingsRepository,
	}
}

func (h GetUserSettingsQueryHandler) Handle(query GetUserSettingsQuery) (*GetUserSettingsResponse, error) {
	settings, err := h.userSettingsRepository.GetUserSettings(query.UserID)
	if err != nil {
		return nil, err
	}
	return NewGetUserSettingsResponse(settings), nil
}
