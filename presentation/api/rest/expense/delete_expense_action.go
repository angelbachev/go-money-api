package expense

import (
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/application/command/expense/delete_expense"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type DeleteExpenseAction struct {
	*rest.BaseAction
	handler *delete_expense.DeleteExpenseCommandHandler
}

func NewDeleteExpenseAction(handler *delete_expense.DeleteExpenseCommandHandler) *DeleteExpenseAction {
	return &DeleteExpenseAction{
		BaseAction: rest.NewBaseAction("Delete", "/accounts/{accountID}/expenses/{expenseID}", false),
		handler:    handler,
	}
}

func (a DeleteExpenseAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	expenseID, err := strconv.ParseInt(chi.URLParam(r, "expenseID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	command := delete_expense.DeleteExpenseCommand{
		ID:        expenseID,
		AccountID: accountID,
		UserID:    a.GetAuthUserID(r),
	}

	if err := a.handler.Handle(command); err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
