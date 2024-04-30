package account

import (
	"net/http"

	"github.com/angelbachev/go-money-api/application/command/account/create_account"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
)

type CreateAccountAction struct {
	*rest.BaseAction
	handler *create_account.CreateAccountCommandHandler
}

func NewCreateAccountAction(handler *create_account.CreateAccountCommandHandler) *CreateAccountAction {
	return &CreateAccountAction{
		BaseAction: rest.NewBaseAction("Post", "/accounts", false),
		handler:    handler,
	}
}

func (a CreateAccountAction) Handle(w http.ResponseWriter, r *http.Request) {
	var command create_account.CreateAccountCommand
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
