package category

import (
	"net/http"

	"github.com/angelbachev/go-money-api/application/query/category/list_category_icons"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
)

type ListCategoryIconsAction struct {
	*rest.BaseAction
	handler *list_category_icons.ListCategoryIconsQueryHandler
}

func NewListCategoryIconsAction(handler *list_category_icons.ListCategoryIconsQueryHandler) *ListCategoryIconsAction {
	return &ListCategoryIconsAction{
		BaseAction: rest.NewBaseAction("Get", "/category-icons", true),
		handler:    handler,
	}
}

func (a ListCategoryIconsAction) Handle(w http.ResponseWriter, r *http.Request) {
	// TODO: validate userID
	query := list_category_icons.ListCategoryIconsQuery{}

	response, err := a.handler.Handle(query)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusOK, response.Icons)
}
