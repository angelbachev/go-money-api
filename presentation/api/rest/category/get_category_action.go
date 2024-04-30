package category

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/query/category/get_category"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type GetCategoryAction struct {
	*rest.BaseAction
	handler *get_category.GetCategoryQueryHandler
}

func NewGetCategoryAction(handler *get_category.GetCategoryQueryHandler) *GetCategoryAction {
	return &GetCategoryAction{
		BaseAction: rest.NewBaseAction("Get", "/accounts/{accountID}/categories/{categoryID}", false),
		handler:    handler,
	}
}

func (a GetCategoryAction) Handle(w http.ResponseWriter, r *http.Request) {
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

	// TODO: validate userID
	query := get_category.GetCategoryQuery{
		ID:        categoryID,
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
