package account

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/command/account/delete_account"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type DeleteAccountAction struct {
	*rest.BaseAction
	handler *delete_account.DeleteAccountCommandHandler
}

func NewDeleteAccountAction(handler *delete_account.DeleteAccountCommandHandler) *DeleteAccountAction {
	return &DeleteAccountAction{
		BaseAction: rest.NewBaseAction("Delete", "/accounts/{accountID}", false),
		handler:    handler,
	}
}

func (a DeleteAccountAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	q := r.URL.Query()
	force, _ := strconv.ParseInt(q.Get("force"), 10, 0)

	command := delete_account.DeleteAccountCommand{
		ID:     accountID,
		UserID: a.GetAuthUserID(r),
		Force:  force != 0,
	}

	if err := a.handler.Handle(command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusNoContent, nil)
}
