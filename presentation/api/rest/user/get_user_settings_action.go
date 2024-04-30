package user

import (
	"net/http"

	"github.com/angelbachev/go-money-api/application/query/user/get_user_settings"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
)

type GetUserSettingsAction struct {
	*rest.BaseAction
	handler *get_user_settings.GetUserSettingsQueryHandler
}

func NewGetUserSettingsAction(handler *get_user_settings.GetUserSettingsQueryHandler) *GetUserSettingsAction {
	return &GetUserSettingsAction{
		BaseAction: rest.NewBaseAction("Get", "/user/settings", false),
		handler:    handler,
	}
}

func (a GetUserSettingsAction) Handle(w http.ResponseWriter, r *http.Request) {
	query := get_user_settings.GetUserSettingsQuery{UserID: a.GetAuthUserID(r)}

	response, err := a.handler.Handle(query)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusOK, response.UserSettings)
}
