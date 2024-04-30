package account

import (
	"net/http"

	"github.com/angelbachev/go-money-api/application/command/account/update_account"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
)

type UpdateAccountAction struct {
	*rest.BaseAction
	handler *update_account.UpdateAccountCommandHandler
}

func NewUpdateAccountAction(handler *update_account.UpdateAccountCommandHandler) *UpdateAccountAction {
	return &UpdateAccountAction{
		BaseAction: rest.NewBaseAction("Put", "/accounts/{accountID}", false),
		handler:    handler,
	}
}

func (a UpdateAccountAction) Handle(w http.ResponseWriter, r *http.Request) {
	var command update_account.UpdateAccountCommand
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

	a.JSON(w, http.StatusCreated, response.Account)
}
