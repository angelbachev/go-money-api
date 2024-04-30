package category

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/command/category/delete_category"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type DeleteCategoryAction struct {
	*rest.BaseAction
	handler *delete_category.DeleteCategoryCommandHandler
}

func NewDeleteCategoryAction(handler *delete_category.DeleteCategoryCommandHandler) *DeleteCategoryAction {
	return &DeleteCategoryAction{
		BaseAction: rest.NewBaseAction("Delete", "/accounts/{accountID}/categories/{categoryID}", false),
		handler:    handler,
	}
}

func (a DeleteCategoryAction) Handle(w http.ResponseWriter, r *http.Request) {
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

	q := r.URL.Query()
	force, _ := strconv.ParseInt(q.Get("force"), 10, 0)

	command := delete_category.DeleteCategoryCommand{
		ID:        categoryID,
		AccountID: accountID,
		UserID:    a.GetAuthUserID(r),
		Force:     force != 0,
	}

	if err := a.handler.Handle(command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusNoContent, nil)
}
