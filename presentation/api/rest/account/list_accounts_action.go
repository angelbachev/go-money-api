package account

import (
	"net/http"

	"github.com/angelbachev/go-money-api/application/query/account/list_accounts"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
)

type ListAccountsAction struct {
	*rest.BaseAction
	handler *list_accounts.ListAccountsQueryHandler
}

func NewListAccountsAction(handler *list_accounts.ListAccountsQueryHandler) *ListAccountsAction {
	return &ListAccountsAction{
		BaseAction: rest.NewBaseAction("Get", "/accounts", false),
		handler:    handler,
	}
}

func (a ListAccountsAction) Handle(w http.ResponseWriter, r *http.Request) {
	// TODO: validate userID
	query := list_accounts.ListAccountsQuery{UserID: a.GetAuthUserID(r)}

	response, err := a.handler.Handle(query)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusOK, response.Accounts)
}
