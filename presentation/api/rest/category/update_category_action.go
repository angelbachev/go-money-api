package category

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/command/category/update_category"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type UpdateCategoryAction struct {
	*rest.BaseAction
	handler *update_category.UpdateCategoryCommandHandler
}

func NewUpdateCategoryAction(handler *update_category.UpdateCategoryCommandHandler) *UpdateCategoryAction {
	return &UpdateCategoryAction{
		BaseAction: rest.NewBaseAction("Put", "/accounts/{accountID}/categories/{categoryID}", false),
		handler:    handler,
	}
}

func (a UpdateCategoryAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	categoryID, err := strconv.ParseInt(chi.URLParam(r, "categoryID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	var command update_category.UpdateCategoryCommand
	if err := a.Body(r, &command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	command.ID = categoryID
	command.UserID = a.GetAuthUserID(r)
	command.AccountID = accountID

	response, err := a.handler.Handle(command)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusCreated, response.Category)
}
