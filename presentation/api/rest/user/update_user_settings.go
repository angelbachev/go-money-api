package user

import (
	"net/http"

	"github.com/angelbachev/go-money-api/application/command/user/update_user_settings"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
)

type UpdateUserSettingsAction struct {
	*rest.BaseAction
	handler *update_user_settings.UpdateUserSettingsCommandHandler
}

func NewUpdateUserSettingsAction(handler *update_user_settings.UpdateUserSettingsCommandHandler) *UpdateUserSettingsAction {
	return &UpdateUserSettingsAction{
		BaseAction: rest.NewBaseAction("Put", "/user/settings", false),
		handler:    handler,
	}
}

func (a UpdateUserSettingsAction) Handle(w http.ResponseWriter, r *http.Request) {
	var command update_user_settings.UpdateUserSettingsCommand
	if err := a.Body(r, &command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	command.UserID = a.GetAuthUserID(r)

	response, err := a.handler.Handle(command)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusCreated, response.UserSettings)
}
