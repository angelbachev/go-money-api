package auth

import (
	"net/http"

	"github.com/angelbachev/go-money-api/application/command/auth/create_auth_token"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
)

type CreateAuthTokenAction struct {
	*rest.BaseAction
	handler *create_auth_token.CreateAuthTokenCommandHandler
}

func NewCreateAuthTokenAction(handler *create_auth_token.CreateAuthTokenCommandHandler) *CreateAuthTokenAction {
	return &CreateAuthTokenAction{
		BaseAction: rest.NewBaseAction("Post", "/auth/tokens", true),
		handler:    handler,
	}
}

func (a CreateAuthTokenAction) Handle(w http.ResponseWriter, r *http.Request) {
	var command create_auth_token.CreateAuthTokenCommand
	if err := a.Body(r, &command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	response, err := a.handler.Handle(command)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusCreated, response)
}
