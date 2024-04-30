package category

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/query/category/list_categories"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type ListCategoriesAction struct {
	*rest.BaseAction
	handler *list_categories.ListCategoriesQueryHandler
}

func NewListCategoriesAction(handler *list_categories.ListCategoriesQueryHandler) *ListCategoriesAction {
	return &ListCategoriesAction{
		BaseAction: rest.NewBaseAction("Get", "/accounts/{accountID}/categories", false),
		handler:    handler,
	}
}

func (a ListCategoriesAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	// TODO: validate userID
	query := list_categories.ListCategoriesQuery{
		UserID:    a.GetAuthUserID(r),
		AccountID: accountID,
	}

	response, err := a.handler.Handle(query)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusOK, response.CategoryTree)
}
