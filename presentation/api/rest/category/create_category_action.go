package category

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/command/category/create_category"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type CreateCategoryAction struct {
	*rest.BaseAction
	handler *create_category.CreateCategoryCommandHandler
}

func NewCreateCategoryAction(handler *create_category.CreateCategoryCommandHandler) *CreateCategoryAction {
	return &CreateCategoryAction{
		BaseAction: rest.NewBaseAction("Post", "/{accountId}/categories", false),
		handler:    handler,
	}
}

func (a CreateCategoryAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	var command create_category.CreateCategoryCommand
	if err := a.Body(r, &command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	command.UserID = a.GetAuthUserID(r)
	command.AccountID = accountID

	response, err := a.handler.Handle(command)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusCreated, response.Category)
}
