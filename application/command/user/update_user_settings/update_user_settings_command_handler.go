package update_user_settings

import (
	"github.com/angelbachev/go-money-api/domain/user"
)

type UpdateUserSettingsCommandHandler struct {
	userSettingsRepository user.UserSettingsRepositoryInterface
}

func NewUpdateUserSettingsCommandHandler(
	userSettingsRepository user.UserSettingsRepositoryInterface,
) *UpdateUserSettingsCommandHandler {
	return &UpdateUserSettingsCommandHandler{
		userSettingsRepository: userSettingsRepository,
	}
}

func (h UpdateUserSettingsCommandHandler) Handle(command UpdateUserSettingsCommand) (*UpdateUserSettingsResponse, error) {
	settings, err := h.userSettingsRepository.GetUserSettings(command.UserID)
	if err != nil {
		return nil, err
	}

	settings.Update(command.DefaultAccountID, command.Theme)
	h.userSettingsRepository.UpdateUserSettings(settings)

	return NewUpdateUserSettingsResponse(settings), nil
}
