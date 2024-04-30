package user

import (
	"net/http"

	"github.com/angelbachev/go-money-api/application/command/user/create_user"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
)

type CreateUserAction struct {
	*rest.BaseAction
	handler *create_user.CreateUserCommandHandler
}

func NewCreateUserAction(handler *create_user.CreateUserCommandHandler) *CreateUserAction {
	return &CreateUserAction{
		BaseAction: rest.NewBaseAction("Post", "/users", true),
		handler:    handler,
	}
}

func (a CreateUserAction) Handle(w http.ResponseWriter, r *http.Request) {
	var command create_user.CreateUserCommand
	if err := a.Body(r, &command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	response, err := a.handler.Handle(command)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusCreated, response.User)
}
